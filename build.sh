#! /bin/sh

game=$1
go run github.com/tmedwards/tweego games/$game -l -o dist/$game.html --head=analytics.html
