package repositories

import (
	"server/models"
	"server/utils"
)

type UserQuizRepositoryImpl struct {
	inmemoryDb map[string]models.UserQuiz
}

func NewUserQuizRepositoryImpl(inmemoryDb map[string]models.UserQuiz) UserQuizRepository {
	return &UserQuizRepositoryImpl{inmemoryDb: inmemoryDb}
}

func (u *UserQuizRepositoryImpl) SetUser(name string) int32 {
	if _, ok := u.inmemoryDb[name]; !ok {
		quizScore := make(map[int32]float32)
		u.inmemoryDb[name] = models.UserQuiz{
			UserName:  name,
			QuizScore: quizScore,
		}
		return utils.STATUS_USER_CREATED
	}
	return utils.STATUS_USER_EXISTS
}

func (u *UserQuizRepositoryImpl) UserQuizRank(name string, quizID int32) float32 {
	userData := u.inmemoryDb[name]
	userQuizScore := userData.QuizScore
	var userHasBetterResult float32
	var totalUsersForQuiz float32
	if len(u.inmemoryDb) == 1 {
		return 1
	}
	for _, k := range u.inmemoryDb {
		if k.UserName != name {
			if score, ok := k.QuizScore[quizID]; ok {
				if score <= userQuizScore[quizID] {
					userHasBetterResult += 1
				}
				totalUsersForQuiz += 1
			}
		}
	}
	return userHasBetterResult / max(totalUsersForQuiz, 1.0)
}

func (u *UserQuizRepositoryImpl) SetUserQuizScore(name string, quizID int32, score float32) {
	u.inmemoryDb[name].QuizScore[quizID] = score
}

func (u *UserQuizRepositoryImpl) UserQuizScore(name string, quizID int32) float32 {
	return u.inmemoryDb[name].QuizScore[quizID]
}
