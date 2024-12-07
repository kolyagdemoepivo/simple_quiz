package cmd

import (
	"bufio"
	"client/api"
	"fmt"
	"os"
	pb "quiz_schema"
	"strings"

	"github.com/spf13/cobra"
)

func NewQuizCommand(quizApi *api.QuizApi, name string) *cobra.Command {
	return &cobra.Command{
		Use:   "quiz",
		Short: "This command starts the simple test",
		Run: func(cmd *cobra.Command, args []string) {
			quizShell(quizApi, name)
		},
	}
}

func handleCommand(input string, ids []int32) int32 {

	switch input {
	case "\\q":
		{
			fmt.Println("Exiting quiz shell.")
			return -1
		}
	case "a":
		{
			return ids[0]
		}
	case "b":
		{
			return ids[1]
		}
	case "c":
		{
			return ids[2]
		}
	default:
		{
			fmt.Println("Select a, b or c")
			return -2
		}

	}
}

func quizShell(quizApi *api.QuizApi, name string) {
	reader := bufio.NewReader(os.Stdin)
	quizEntries := quizApi.GetQuizEntries()
	var quizAnswers []*pb.QuizSubmitEntry
	fmt.Println("\n\nType `a`, `b` or `c` to select the correct answer.")
	for i, entry := range quizEntries.QuizEntries {
		var answerIDs []int32
		fmt.Printf("Question %d: %s\n", i+1, entry.Question.Text)
		fmt.Print("\033[36m\n")
		for j, answer := range entry.Answers {
			fmt.Printf("%s) %s\n", []string{"a", "b", "c"}[j], answer.Text)
			answerIDs = append(answerIDs, answer.Id)
		}
		for {
			fmt.Print("\033[0m")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			answerID := handleCommand(input, answerIDs)
			if answerID == -2 {
				continue
			} else if answerID == -1 {
				return
			}
			quizAnswers = append(quizAnswers, &pb.QuizSubmitEntry{
				AnswerId:   answerID,
				QuestionId: entry.Question.Id,
			})
			break
		}

	}
	resp := quizApi.SubmitQuizResult(name, quizAnswers, quizEntries.QuizId)
	fmt.Println("Congratulations, you completed the test!")
	fmt.Println("Your number of correct answers:", resp.CorrectAnswers)
	fmt.Printf("You were better than %d%% of all quizzers\n", int(resp.Rank*100))
}
