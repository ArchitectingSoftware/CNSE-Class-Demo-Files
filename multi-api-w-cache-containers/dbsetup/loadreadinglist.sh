#!/bin/bash
cat readinglist.json | jq -c '.[]' |\
    while read json_object; do \
        rlid=$(jq -r '.id' <<< $json_object); \
        #echo $pubid  \
        rediscmd="redis-cli JSON.set publist:$rlid . '$json_object'"; \
        echo $rediscmd; \
        eval $rediscmd; \
    done 