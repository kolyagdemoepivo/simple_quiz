package controllers

import (
	"context"
	pb "quiz_schema"
	"server/mappers"
	"server/services"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type QuizController struct {
	quizService     services.QuizService
	userQuizService services.UserQuizService
	pb.UnimplementedQuizGrpcServer
}

func (q *QuizController) GetQuizEntries(context.Context, *emptypb.Empty) (*pb.QuizResponse, error) {
	quiz := q.quizService.GetQuiz(1)
	quizEntries := mappers.QuizModelToProto(quiz)
	return &pb.QuizResponse{
		QuizEntries: quizEntries,
		QuizId:      quiz.ID,
	}, nil
}

func (q *QuizController) CreateUser(_ context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	status := q.userQuizService.SetUser(in.Name)

	return &pb.CreateUserResponse{
		Status: status,
	}, nil
}

func (q *QuizController) SubmitQuiz(_ context.Context, in *pb.QuizSubmitRequest) (*pb.QuizSubmitResponse, error) {
	quiz := q.quizService.GetQuiz(in.QuizId)
	totalQuestionsLen := max(float32(len(quiz.QuizItems)), 1)
	var correectAnswerCounter float32
	for _, submission := range in.Submissions {
		questionID := submission.QuestionId
		answerID := submission.AnswerId
		correctAnswer := q.quizService.GetQuestionCorrectAnswer(in.QuizId, questionID)
		if correctAnswer != nil && answerID == correctAnswer.ID {
			correectAnswerCounter += 1
		}
	}
	q.userQuizService.SetUserQuizScore(in.Name, in.QuizId, correectAnswerCounter/totalQuestionsLen)
	return &pb.QuizSubmitResponse{
		CorrectAnswers: int32(correectAnswerCounter),
		Rank:           q.userQuizService.UserQuizRank(in.Name, in.QuizId),
	}, nil
}

func NewQuizController(quizService services.QuizService, userQuizService services.UserQuizService) *QuizController {
	return &QuizController{quizService: quizService, userQuizService: userQuizService}
}
