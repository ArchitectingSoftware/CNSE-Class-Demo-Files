#!/bin/bash
cat pubs.json | jq -c '.[]' |\
    while read json_object; do \
        pubid=$(jq -r '.id' <<< $json_object); \
        #echo $pubid  \
        rediscmd="redis-cli JSON.set pubs:$pubid . '$json_object'"; \
        echo $rediscmd; \
        eval $rediscmd; \
    done 