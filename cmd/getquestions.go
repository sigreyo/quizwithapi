/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// getquestionsCmd represents the getquestions command
var getquestionsCmd = &cobra.Command{
	Use:     "start",
	Short:   "Start the quiz",
	Example: "go run . start",
	Run: func(cmd *cobra.Command, args []string) {
		runQuiz()
	},
}

type Question struct {
	ID            int      `json:"id"`
	Question      string   `json:"question"`
	Answers       []string `json:"answers"`
	CorrectAnswer int      `json:"correctanswer"`
}

type Answer struct {
	Username string `json:"username"`
	Score    int    `json:"score"`
}

//GET questions+answer from API and run the quiz
func getQuestions() []Question {
	resp, err := http.Get("http://localhost:8080/api/questions")
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Could not read response")
	}

	var response []Question

	if err := json.Unmarshal(body, &response); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Could not unmarshal JSON")
	}

	return response

}

//the actual quiz
func runQuiz() {

	//GET quizset
	quiz := getQuestions()

	correctAnswers := 0
	var user string
	fmt.Println("Quiz time! Skriv in ditt namn: ")
	fmt.Scanln(&user)
	user = cases.Title(language.Und).String(user)

	for _, v := range quiz {

		//prints the questions with alternatives
		fmt.Printf("(%d.) %v\n", v.ID, v.Question)
		for i, v := range v.Answers {
			fmt.Printf("[%v] %s ", i+1, v) //increment i+1 to match expected input from user
		}
		fmt.Print("\n")

		//check if answer is valid
		var answer string
		validAnswer := []string{"1", "2", "3", "4"}
		for {
			fmt.Scanln(&answer)
			if slices.Contains(validAnswer, answer) {
				break
			} else {
				fmt.Println("Svaret måste vara en siffra mellan 1-4.")
			}
		}
		entry, _ := strconv.Atoi(answer)

		if entry == v.CorrectAnswer {
			fmt.Printf("Rätt svar!\n")
			correctAnswers++
		} else {
			fmt.Printf("FEL! Rätt svar är [%v] %s \n\n", v.CorrectAnswer, v.Answers[v.CorrectAnswer-1])
		}
	}

	fmt.Printf("Snyggt %s! Du fick %v/%v rätt. \n", user, correctAnswers, len(quiz))
	fmt.Printf("Ditt resultat är bättre än %v%% av andra tävlande.", calculatePerformance(correctAnswers))

	//what gets posted to the API
	score := Answer{
		Username: user,
		Score:    correctAnswers,
	}

	//POST to API
	postToAPI(score)

}
func postToAPI(score Answer) {
	reqBody, err := json.Marshal(score)
	if err != nil {
		fmt.Println("Could not marshal json")
	}
	resp, err := http.Post("http://localhost:8080/api/answers", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()
}

func init() {
	rootCmd.AddCommand(getquestionsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getquestionsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getquestionsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// getquestionsCmd.Flags().StringP("id", "i", "", "ID of the question to be displayed")
}
