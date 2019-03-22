#!/bin/bash

docker build --tag=nav-api .

#test
docker tag nav-api gcr.io/staging-217409/nav-api

docker push gcr.io/staging-217409/nav-api


#prod
docker tag nav-api gcr.io/production-217408/nav-api

docker push gcr.io/production-217408/nav-api
