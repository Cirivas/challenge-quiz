// Code generated by MockGen. DO NOT EDIT.
// Source: answer_quiz.go

// Package mock_answer_quiz is a generated GoMock package.
package mock_answer_quiz

import (
	reflect "reflect"

	entities "github.com/cirivas/challenge-quiz/core/entities"
	gomock "github.com/golang/mock/gomock"
)

// MockAnswerQuizUseCase is a mock of AnswerQuizUseCase interface.
type MockAnswerQuizUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockAnswerQuizUseCaseMockRecorder
}

// MockAnswerQuizUseCaseMockRecorder is the mock recorder for MockAnswerQuizUseCase.
type MockAnswerQuizUseCaseMockRecorder struct {
	mock *MockAnswerQuizUseCase
}

// NewMockAnswerQuizUseCase creates a new mock instance.
func NewMockAnswerQuizUseCase(ctrl *gomock.Controller) *MockAnswerQuizUseCase {
	mock := &MockAnswerQuizUseCase{ctrl: ctrl}
	mock.recorder = &MockAnswerQuizUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAnswerQuizUseCase) EXPECT() *MockAnswerQuizUseCaseMockRecorder {
	return m.recorder
}

// AnswerQuiz mocks base method.
func (m *MockAnswerQuizUseCase) AnswerQuiz(quizId, respondent string, quiz *entities.Quiz, answers []entities.AnswerKey) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AnswerQuiz", quizId, respondent, quiz, answers)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AnswerQuiz indicates an expected call of AnswerQuiz.
func (mr *MockAnswerQuizUseCaseMockRecorder) AnswerQuiz(quizId, respondent, quiz, answers interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AnswerQuiz", reflect.TypeOf((*MockAnswerQuizUseCase)(nil).AnswerQuiz), quizId, respondent, quiz, answers)
}
