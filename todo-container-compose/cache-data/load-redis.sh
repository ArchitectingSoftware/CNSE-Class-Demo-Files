#!/bin/bash
VAR=${1:-localhost}    
cat /data/redis-load.redis | redis-cli -h $1