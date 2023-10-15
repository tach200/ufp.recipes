package db

import (
	"fmt"
	"strings"

	sqlmapper "github.com/jackskj/carta"
	"github.com/nullism/bqb"
)

type Product struct {
	ID          string  `db:"id" json:"id"`
	Supermarket string  `db:"supermarket" json:"supermarket"`
	ItemID      string  `db:"item_id" json:"item_id"`
	ProductName string  `db:"product_name" json:"product_name"`
	ProductDesc string  `db:"product_desc" json:"product_desc"`
	SalesUnit   string  `db:"sales_unit" json:"sales_unit"`
	PricePerUOM string  `db:"price_per_uom" json:"price_per_uom"`
	Price       string  `db:"price" json:"price"`
	Rank        float64 `db:"rank" json:"rank"`
}

const selectProductBuilder = `
	SELECT *,
		ts_rank(to_tsvector('english', product_name), plainto_tsquery('english', ?)) AS rank
	FROM products
	WHERE (
		?
	)
	ORDER BY rank DESC;
`

// SelectProduct
func (db *DB) SelectProduct(ingrediants string) ([]*Product, error) {
	// build the sql query, this query is a bit hacky but it works for now
	ingrediantsSlice := strings.Fields(ingrediants)

	subWhere := bqb.Optional("")

	i := 0
	for i < len(ingrediantsSlice) {
		subWhere.Or("to_tsvector('english', product_name) @@ plainto_tsquery('english', ?)", ingrediantsSlice[i])
		i++
	}

	selectProductBuilder := bqb.New(selectProductBuilder, ingrediants, subWhere)

	selectProductSQL, args, err := selectProductBuilder.ToPgsql()
	if err != nil {
		return nil, err
	}

	rows, err := db.Client.Query(selectProductSQL, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []*Product{}
	err = sqlmapper.Map(rows, &products)
	if err != nil {
		return nil, err
	}

	if len(products) == 0 {
		fmt.Printf("WARN: no products found for %s\n", ingrediants)
	}

	return products, nil
}
