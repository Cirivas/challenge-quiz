package quiz_controller

import "github.com/cirivas/challenge-quiz/core/entities"

type QuizDTO struct {
	Id        string        `json:"id"`
	Questions []QuestionDTO `json:"questions"`
}

type QuestionDTO struct {
	Text         string                        `json:"text"`
	Alternatives map[entities.AnswerKey]string `json:"alternatives"`
}

func QuestionEntityToDTO(entity *entities.Question) *QuestionDTO {
	if entity == nil {
		return nil
	}

	return &QuestionDTO{
		Text:         entity.Text,
		Alternatives: entity.Alternatives,
	}
}

func QuizEntityToDTO(quizId string, entity *entities.Quiz) *QuizDTO {
	if entity == nil {
		return nil
	}

	questions := make([]QuestionDTO, len(entity.Questions))

	for i, q := range entity.Questions {
		questions[i] = *QuestionEntityToDTO(&q)
	}

	return &QuizDTO{
		Id:        quizId,
		Questions: questions,
	}
}
