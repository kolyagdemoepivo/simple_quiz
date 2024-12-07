package mappers

import (
	pb "quiz_schema"
	"server/models"
)

func QuizModelToProto(quiz *models.Quiz) []*pb.QuizEntry {
	var quizEntries []*pb.QuizEntry
	for _, quizItem := range quiz.QuizItems {
		var answers []*pb.AnswerOption
		for _, quizAnswer := range quizItem.Answers {
			answers = append(answers, &pb.AnswerOption{
				Id:   quizAnswer.ID,
				Text: quizAnswer.Text,
			})
		}
		quizEntries = append(quizEntries, &pb.QuizEntry{
			Question: &pb.QuizQuestion{
				Id:   quizItem.Question.ID,
				Text: quizItem.Question.Text,
			},
			Answers: answers,
		},
		)
	}
	return quizEntries
}
