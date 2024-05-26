package models

import "fmt"

type Response struct {
	Score      int    `json:"score"`
	Respondent string `json:"respondent"`
	QuizId     string `json:"quizId"`
}

func (r Response) Key() string {
	return fmt.Sprintf("%s:%s", r.Respondent, r.QuizId)
}
