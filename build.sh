#! /bin/sh

game=$1
go run github.com/tmedwards/tweego \
    --log-files -l \
    --head=analytics.html \
    shared/ \
    games/$game \
    -o dist/$game.html
