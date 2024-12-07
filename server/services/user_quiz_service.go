package services

type UserQuizService interface {
	SetUser(name string) int32
	SetUserQuizScore(name string, quizID int32, score float32)
	UserQuizScore(name string, quizID int32) float32
	UserQuizRank(name string, quizID int32) float32
}
