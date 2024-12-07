package models

type UserQuiz struct {
	UserName  string
	QuizScore map[int32]float32
}
