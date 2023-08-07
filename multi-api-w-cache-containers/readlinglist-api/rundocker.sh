#!/bin/bash
docker run --name cnse-publist-api --rm -e RLAPI_PUB_API_URL=http://host.docker.internal:2080 -p 3080:3080 architectingsoftware/cnse-publist-api:v1
