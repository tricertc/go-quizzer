package models
import "strings"

type Question struct {
    Text string
    Options []Option
    Answers []string
    Explanation string
}

func (q *Question) AddOption(option Option) {
    q.Options = append(q.Options, option)
}

func (q *Question) AddAnswer(answer string) {
    q.Answers = append(q.Answers, answer)
}

func (q *Question) Validate(answers []string) bool {
    N := len(q.Answers)
    count := 0

    for i := 0; i < N; i++ {
        for j := 0; j < len(answers); j++ {
            if strings.ToLower(q.Answers[i]) == strings.ToLower(answers[j]) {
                count++
            }
        }
    }

    return count == N
}