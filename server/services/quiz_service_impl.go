package services

import (
	"server/models"
	"server/repositories"
)

type QuizServiceImpl struct {
	quizRepo repositories.QuizRepository
}

func (q *QuizServiceImpl) GetQuiz(ID int32) *models.Quiz {
	return q.quizRepo.GetQuiz(ID)
}

func (q *QuizServiceImpl) GetQuestionCorrectAnswer(ID int32, questionID int32) *models.Answer {
	return q.quizRepo.GetQuestionCorrectAnswer(ID, questionID)
}

func NewQuizServiceImpl(quizRepo repositories.QuizRepository) QuizService {
	return &QuizServiceImpl{quizRepo: quizRepo}
}
