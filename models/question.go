package models

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