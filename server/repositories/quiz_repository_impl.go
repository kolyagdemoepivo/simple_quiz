package repositories

import (
	"server/data"
	"server/models"
)

type QuizRepositoryImpl struct {
}

func (q *QuizRepositoryImpl) GetQuiz(ID int32) *models.Quiz {
	for _, quiz := range data.QuizData {
		if quiz.ID == ID {
			return &quiz
		}
	}
	return nil
}

func (q *QuizRepositoryImpl) GetQuestionCorrectAnswer(ID int32, questionID int32) *models.Answer {

	quiz := q.GetQuiz(ID)
	for _, quizItem := range quiz.QuizItems {
		if quizItem.Question.ID == questionID {
			for _, answer := range quizItem.Answers {
				if answer.IsCorrect {
					return &answer
				}
			}
		}
	}
	return nil
}

func NewQuizRepositoryImpl() QuizRepository {
	return &QuizRepositoryImpl{}
}
