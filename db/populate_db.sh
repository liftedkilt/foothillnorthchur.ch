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
        "Video1": "Come all ye Faithful",
        "Video2": "Amazing Grace",
        "Video3": "A Mighty Fortress",
        "Video4": "Be Thou my Vision",
        "Video5": "I will Build My Church",
        "Video6": "Oh For a Thousand Tongues"
    }'

fi