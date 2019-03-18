
openssl genrsa -out localhost.key 2048

openssl req -new -sha256 \
    -key localhost.key \
    -subj "/C=SE/ST=Stockholm/O=Hem Design Studio, Inc./CN=localhost" \
    -reqexts SAN \
    -config <(cat /etc/ssl/openssl.cnf \
        <(printf "\n[SAN]\nsubjectAltName=DNS:localhost,DNS:127.0.0.1")) \
    -out localhost.csr

openssl req -in localhost.csr -noout -text

openssl x509 -req -in localhost.csr -CA rootCA.crt -CAkey rootCA.key -CAcreateserial -out localhost.crt -days 500 -sha256

openssl x509 -in localhost.crt -text -noout

#------------ nginx -------------------#

openssl genrsa -out nginx.key 2048

openssl req -new -sha256 \
    -key nginx.key \
    -subj "/C=SE/ST=Stockholm/O=Hem Design Studio, Inc./CN=nav-api.endpoints.staging-217409.cloud.goog" \
    -reqexts SAN \
    -config <(cat /etc/ssl/openssl.cnf \
        <(printf "\n[SAN]\nsubjectAltName=DNS:nav-api.endpoints.staging-217409.cloud.goog,DNS:www.nav-api.endpoints.staging-217409.cloud.goog")) \
    -out nginx.csr

openssl req -in nginx.csr -noout -text

openssl x509 -req -in nginx.csr -CA rootCA.crt -CAkey rootCA.key -CAcreateserial -out nginx.crt -days 500 -sha256

openssl x509 -in nginx.crt -text -noout