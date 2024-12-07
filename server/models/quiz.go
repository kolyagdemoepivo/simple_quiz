package models

type QuizItem struct {
	ID       int32
	Question Question
	Answers  []Answer
}

type Quiz struct {
	ID        int32
	Title     string
	QuizItems []QuizItem
}
