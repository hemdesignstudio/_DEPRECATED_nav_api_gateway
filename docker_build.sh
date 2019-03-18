#!/bin/bash

docker build --tag=nav-api .

docker tag nav-api gcr.io/staging-217409/nav-api

docker push gcr.io/staging-217409/nav-api
