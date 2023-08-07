#!/bin/bash
VAR=${1:-localhost}    

#delete previous db if it exists and install jq
apt-get -y install jq

#delete the database
redis-cli -h $1 flushdb

#load pubs
cat /data/pubs.json | jq -c '.[]' |\
    while read json_object; do \
        pubid=$(jq -r '.id' <<< $json_object); \
        #echo $pubid  \
        rediscmd="redis-cli -h $1 JSON.set pubs:$pubid . '$json_object'"; \
        echo $rediscmd; \
        eval $rediscmd; \
    done 

#load reading list
cat /data/readinglist.json | jq -c '.[]' |\
    while read json_object; do \
        rlid=$(jq -r '.id' <<< $json_object); \
        #echo $pubid  \
        rediscmd="redis-cli -h $1 JSON.set publist:$rlid . '$json_object'"; \
        echo $rediscmd; \
        eval $rediscmd; \
    done 

