FROM alpine AS builder
RUN apk add --no-cache ca-certificates

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
ADD out/ufp.recipes /recipes

ENTRYPOINT ["/recipes"]
