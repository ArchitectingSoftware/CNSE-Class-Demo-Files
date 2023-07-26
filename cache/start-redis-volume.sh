#!/bin/bash
docker run -d --rm --name cnse-redis -v ./cache-data/:/data -e REDIS_ARGS="--appendonly yes"  -p 6379:6379 -p 8001:8001 redis/redis-stack:latest 