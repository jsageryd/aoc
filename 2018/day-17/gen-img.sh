#!/bin/sh
mkdir img 2>/dev/null
< example-input go run main.go -ascii > img/example-output.txt
< example-input go run main.go -gif -scale=15 > img/example-output.gif
< example-input go run main.go -agif -scale=15 > img/example-output-animated.gif
< input go run main.go -ascii > img/output.txt
< input go run main.go -gif -scale=3 > img/output.gif
< input go run main.go -agif -scale=2 > img/output-animated.gif
