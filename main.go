package main

import(
    "flag"
    "github.com/tricertc/quizzer/models"
    "log"
)

var filename string
var maxerror int
var quiz models.Quiz

func main() {
    flag.StringVar(&filename, "fn", "", "data file path")
    flag.IntVar(&maxerror, "max", 0, "maximum errors before exiting")
    flag.Parse()

    _, err := quiz.Load(filename)
    if err != nil {
        log.Fatal(err)
    }

    quiz.Play(maxerror)
}
