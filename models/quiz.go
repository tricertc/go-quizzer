package models
import "os"

type Quiz struct {
    Questions []Question
    Score int
}

func (q *Quiz) Load(filename string) (bool, error) {
    q.Score = 0
    q.Questions = make([]Question, 0)

    f, err := os.Open(filename)
    if err != nil {
        return false, err
    }
    defer f.Close()

    return true, nil
}