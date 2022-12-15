#!/bin/sh

mkdir img 2>/dev/null

# Part 1
< example-input go run main.go -ascii > img/part-1-example-output.txt
< example-input go run main.go -gif -scale=15 > img/part-1-example-output.gif
< example-input go run main.go -agif -scale=15 > img/part-1-example-output-animated.gif
< input go run main.go -ascii > img/part-1-output.txt
< input go run main.go -gif -scale=10 > img/part-1-output.gif
< input go run main.go -agif -scale=10 > img/part-1-output-animated.gif

# Part 2
< example-input go run main.go -floor -ascii > img/part-2-example-output.txt
< example-input go run main.go -floor -gif -scale=15 > img/part-2-example-output.gif
< example-input go run main.go -floor -agif -scale=15 > img/part-2-example-output-animated.gif
< input go run main.go -floor -ascii > img/part-2-output.txt
< input go run main.go -floor -gif -scale=10 > img/part-2-output.gif
# < input go run main.go -floor -agif -scale=10 > img/part-2-output-animated.gif
