package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	pb "quiz_schema"
	"server/controllers"
	"server/models"
	"server/repositories"
	"server/services"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 6690, "The server port")
)

func main() {
	inMemoryDb := make(map[string]models.UserQuiz)

	quizRepository := repositories.NewQuizRepositoryImpl()
	userQuizRepository := repositories.NewUserQuizRepositoryImpl(inMemoryDb)

	quizService := services.NewQuizServiceImpl(quizRepository)
	userQuizService := services.NewUserQuizServiceImpl(userQuizRepository)

	quizController := controllers.NewQuizController(quizService, userQuizService)
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterQuizGrpcServer(s, quizController)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
