package repository

type ScoreRepository interface {
	SaveScore(respondent string, score int, quizId string) error
	// Return the score of 'respondent' for 'quizId'
	GetScore(respondent string, quizId string) (int, error)
	// Returns the score of everyone except 'respondent' for 'quizId'
	GetOthersScore(respondent string, quizId string) ([]int, error)
}
