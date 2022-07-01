/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

// getquestionsCmd represents the getquestions command
var getquestionsCmd = &cobra.Command{
	Use:   "getquestions",
	Short: "Get all questions",
	Long:  "Getting all questions from seed",
	Run: func(cmd *cobra.Command, args []string) {
		getAll()
	},
}

type QuestionResponse struct {
	ID       int    `json:"ID"`
	Question string `json:"Question"`
}

func getAll() {
	resp, err := http.Get("http://localhost:8080/api/questions")
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Could not read response")
	}

	var result []QuestionResponse

	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	for _, v := range result {
		fmt.Println(v.ID, " ", v.Question)
	}
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
}
