package models
import (
    "os"
    "bufio"
    "strings"
    "math/rand"
    "fmt"
    "github.com/tricertc/quizzer/utils"
    "time"
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
    rand.Seed(time.Now().UnixNano())
    N := len(q.Questions)
    for i := 0; i < N; i++ {
        j := i + rand.Intn(N - i)
        q.Questions[i], q.Questions[j] = q.Questions[j], q.Questions[i]
    }
}

func (q *Quiz) Play(maxerror int, shuffle bool) {
    if shuffle {
        q.shuffle()
    }

    for i, qq := range(q.Questions) {
        utils.Clear()

        fmt.Printf("Question %d:\n\n", i + 1)
        fmt.Printf("  %s\n\n", qq.Text)

        for _, opt := range(qq.Options) {
            fmt.Printf("    %s) %s\n\n", opt.Label, opt.Text)
        }

        N := len(qq.Answers)
        answers := make([]string, N)

        for i, _ := range(answers) {
            var ans string

            if N == 1 {
                fmt.Printf("  Answer: ")
            } else {
                fmt.Printf(" Answer %d of %d: ", i+1, N)
            }

            fmt.Scanf("%s\n", &ans)
            answers[i] = ans
        }

        result := qq.Validate(answers)
        fmt.Printf("\n")
        if result {
            q.Score++
            fmt.Printf("Correct!")
        } else {
            fmt.Printf("Wrong.  The correct answer was %s", strings.Join(qq.Answers, " and "))
        }
        fmt.Printf("\n\n")
        fmt.Printf("Explanation:\n\n  %s\n\n", qq.Explanation)

        fmt.Printf("<enter> to continue...")
        fmt.Scanf("\n")

        if maxerror > 0 && (i + 1) - q.Score >= maxerror {
            fmt.Printf("\n\n*** FAILED ***\n\n")
            fmt.Printf("<enter> to continue...")
            fmt.Scanf("\n")
            break
        }
    }

    utils.Clear()

    fmt.Printf("*** TEST COMPLETE ***\n\n")
    fmt.Printf("  You answered %d correctly out of %d\n\n\n", q.Score, len(q.Questions))

    pct := (float64(q.Score) / float64(len(q.Questions))) * 100.00
    fmt.Printf("Final Score: %.2f%%\n\n\n\n", pct)
}