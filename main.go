package main

import(
    "flag"
)

var filename string

func main() {
    flag.StringVar(&filename, "fn", "", "data file path")
    flag.Parse()
}