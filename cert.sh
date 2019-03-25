#!/bin/bash

#---- Root CA----------#
openssl genrsa -des3 -out rootCA.key 2048

openssl req -x509 -new -nodes -key rootCA.key -sha256 -days 1825 -out rootCA.pem


#---------- localhost commands ---------- #

openssl genrsa -out localhost.key 2048

openssl req -new -key localhost.key -out localhost.csr

openssl req -in localhost.csr -noout -text

openssl x509 -req -in localhost.csr -CA rootCA.pem -CAkey rootCA.key -CAcreateserial -out localhost.crt -days 1825 -sha256 -extfile localhost.ext

openssl x509 -in localhost.crt -text -noout



#------------ nginx test-------------------#

openssl genrsa -out test.key 2048

openssl req -new -key test.key -out test.csr

openssl req -in test.csr -noout -text

openssl x509 -req -in test.csr -CA rootCA.pem -CAkey rootCA.key -CAcreateserial -out test.crt -days 1825 -sha256 -extfile test.ext

openssl x509 -in test.crt -text -noout

#------------ nginx prod-------------------#

openssl genrsa -out nginx.key 2048

openssl req -new -key nginx.key -out nginx.csr

openssl req -in nginx.csr -noout -text

openssl x509 -req -in nginx.csr -CA rootCA.pem -CAkey rootCA.key -CAcreateserial -out nginx.crt -days 1825 -sha256 -extfile nginx.ext

openssl x509 -in nginx.crt -text -noout

#--------- cloud secret creation -----------#
kubectl create secret generic nginx-ssl --from-file=./nginx.crt --from-file=./nginx.key
