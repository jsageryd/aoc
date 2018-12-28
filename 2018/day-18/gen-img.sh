#!/bin/sh
mkdir img 2>/dev/null
< example-input go run main.go -ascii -steps=10 > img/example-output.txt
< example-input go run main.go -agif -delay=40 -scale=20 -steps=20 > img/example-output.gif
< input go run main.go -ascii -steps=10 > img/output.txt
< input go run main.go -agif -delay=8 -scale=10 -steps=800 > img/output.gif
