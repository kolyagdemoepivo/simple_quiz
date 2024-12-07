package repositories

import "server/models"

type QuizRepository interface {
	GetQuiz(ID int32) *models.Quiz
	GetQuestionCorrectAnswer(ID int32, questionID int32) *models.Answer
}
