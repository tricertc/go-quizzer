package models
import (
    "os"
    "bufio"
    "strings"
    "math/rand"
    "fmt"
    "github.com/tricertc/quizzer/utils"
)

type Quiz struct {
    Questions []Question
    Score int
}

func (q *Quiz) AddQuestion(question Question) {
    q.Questions = append(q.Questions, question)
}

func (q *Quiz) Load(filename string) (bool, error) {
    q.Score = 0
    q.Questions = make([]Question, 0)

    f, err := os.Open(filename)
    if err != nil {
        return false, err
    }
    defer f.Close()

    reader := bufio.NewReader(f)
    scanner := bufio.NewScanner(reader)

    for scanner.Scan() {
        s := scanner.Text()
        cols := strings.Split(s, "\t")

        if len(cols) < 4 {
            continue
        }

        question := Question{ Text: cols[0], Explanation: cols[len(cols) - 1]}

        options := cols[1:len(cols) - 2]
        for _, o := range(options) {
            split := strings.Split(o, "|")
            question.AddOption(Option{split[0], strings.Join(split[1:len(split)], "|")})
        }

        answers := strings.Split(cols[len(cols) - 2], "|")
        for _, a := range(answers) {
            question.AddAnswer(a)
        }

        q.AddQuestion(question)
    }

    return true, nil
}

func (q *Quiz) shuffle() {
    N := len(q.Questions)
    for i := 0; i < N; i++ {
        j := i + rand.Intn(N - i)
        q.Questions[i], q.Questions[j] = q.Questions[j], q.Questions[i]
    }
}

func (q *Quiz) Play() {
    q.shuffle()
    for i, qq := range(q.Questions) {
        utils.Clear()

        fmt.Printf("Question %d:\n\n", i + 1)
        fmt.Printf("  %s\n\n", qq.Text)
    }
}