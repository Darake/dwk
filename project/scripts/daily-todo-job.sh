#!/usr/bin/env bash
set -e

RANDOM_WIKI_ARTICLE=`curl -Ls -w %{url_effective} -o /dev/null https://en.wikipedia.org/wiki/Special:Random`
NEW_TODO="Read $RANDOM_WIKI_ARTICLE"
curl -d "$NEW_TODO" -X POST $POST_URL