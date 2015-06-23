package main

import(
    "flag"
    "github.com/tricertc/quizzer/models"
    "log"
    "fmt"
)

var filename string
var quiz models.Quiz

func main() {
    flag.StringVar(&filename, "fn", "", "data file path")
    flag.Parse()

    _, err := quiz.Load(filename)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(quiz)
}
