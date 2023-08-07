#!/bin/bash

#delete the database
redis-cli flushdb

#load pubs
cat pubs.json | jq -c '.[]' |\
    while read json_object; do \
        pubid=$(jq -r '.id' <<< $json_object); \
        #echo $pubid  \
        rediscmd="redis-cli JSON.set pubs:$pubid . '$json_object'"; \
        echo $rediscmd; \
        eval $rediscmd; \
    done 

#load reading list
cat readinglist.json | jq -c '.[]' |\
    while read json_object; do \
        rlid=$(jq -r '.id' <<< $json_object); \
        #echo $pubid  \
        rediscmd="redis-cli JSON.set publist:$rlid . '$json_object'"; \
        echo $rediscmd; \
        eval $rediscmd; \
    done 

