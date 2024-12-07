package api

import (
	"context"
	pb "quiz_schema"

	"google.golang.org/protobuf/types/known/emptypb"
)

type QuizApi struct {
	quizGrpcClient pb.QuizGrpcClient
	ctx            context.Context
}

func NewQuizApi(quizGrpcClient pb.QuizGrpcClient,
	ctx context.Context) *QuizApi {
	return &QuizApi{
		quizGrpcClient: quizGrpcClient,
		ctx:            ctx,
	}
}

func (q *QuizApi) CreateUser(name string) int32 {
	res, _ := q.quizGrpcClient.CreateUser(q.ctx, &pb.CreateUserRequest{
		Name: name,
	})
	return res.GetStatus()
}

func (q *QuizApi) GetQuizEntries() *pb.QuizResponse {
	quizEntries, _ := q.quizGrpcClient.GetQuizEntries(q.ctx, &emptypb.Empty{})
	return quizEntries
}

func (q *QuizApi) SubmitQuizResult(name string, quizSubmitEntry []*pb.QuizSubmitEntry, quizID int32) *pb.QuizSubmitResponse {
	resp, _ := q.quizGrpcClient.SubmitQuiz(q.ctx, &pb.QuizSubmitRequest{
		Name:        name,
		Submissions: quizSubmitEntry,
		QuizId:      quizID,
	})
	return resp
}
