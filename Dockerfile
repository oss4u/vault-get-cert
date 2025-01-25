FROM gcr.io/distroless/static-debian12
RUN mkdir /app
COPY vault-get-cert /app/vault-get-cert
ENTRYPOINT ["/app/vault-get-cert"]
