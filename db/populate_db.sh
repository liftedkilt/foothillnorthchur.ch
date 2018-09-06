#!/bin/bash


if [[ $1 == "songs" ]]; then
    curl -XPOST localhost:8888/songs -d '{"Title": "a", "YoutubeTitle": "aa"}'
    curl -XPOST localhost:8888/songs -d '{"Title": "b", "YoutubeTitle": "bb"}'
    curl -XPOST localhost:8888/songs -d '{"Title": "c", "YoutubeTitle": "cc"}'
    curl -XPOST localhost:8888/songs -d '{"Title": "d", "YoutubeTitle": "dd"}'
    curl -XPOST localhost:8888/songs -d '{"Title": "e", "YoutubeTitle": "ee"}'
    curl -XPOST localhost:8888/songs -d '{"Title": "f", "YoutubeTitle": "ff"}'
fi

if [[ $1 == "playlist" ]]; then
    curl -XPOST localhost:8888/playlist -d '{
        "Video1": 1,
        "Video2": 3,
        "Video3": 5,
        "Video4": 6,
        "Video5": 2,
        "Video6": 4
    }'

fi