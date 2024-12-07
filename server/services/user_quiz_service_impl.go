package services

import "server/repositories"

type UserQuizServiceImpl struct {
	userQuizRepository repositories.UserQuizRepository
}

func NewUserQuizServiceImpl(userQuizRepository repositories.UserQuizRepository) UserQuizService {
	return &UserQuizServiceImpl{userQuizRepository: userQuizRepository}
}

func (u *UserQuizServiceImpl) SetUser(name string) int32 {
	return u.userQuizRepository.SetUser(name)
}

func (u *UserQuizServiceImpl) UserQuizRank(name string, quizID int32) float32 {
	return u.userQuizRepository.UserQuizRank(name, quizID)
}

func (u *UserQuizServiceImpl) SetUserQuizScore(name string, quizID int32, score float32) {
	u.userQuizRepository.SetUserQuizScore(name, quizID, score)
}

func (u *UserQuizServiceImpl) UserQuizScore(name string, quizID int32) float32 {
	return u.userQuizRepository.UserQuizScore(name, quizID)
}
