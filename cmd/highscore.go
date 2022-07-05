/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"sort"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// highscoreCmd represents the highscore command
var highscoreCmd = &cobra.Command{
	Aliases: []string{"hs"},
	Use:     "highscore",
	Short:   "Show highscore list",
	Long: `Display a highscore list of all users that have taken the quiz.
	Only displays in-memory records since the API was started.`,
	Run: func(cmd *cobra.Command, args []string) {
		displayHighscore()
	},
}

//GET highscore list from API
func getHighscore() []Answer {
	resp, err := http.Get("http://localhost:8080/api/answers")
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Could not read response")
	}
	var Result []Answer

	if err := json.Unmarshal(body, &Result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Could not unmarshal JSON")
	}
	return Result
}

//sorting
type ByScore []Answer

func (a ByScore) Len() int           { return len(a) }
func (a ByScore) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByScore) Less(i, j int) bool { return a[i].Score < a[j].Score }

//prints a sorted highscore list
func displayHighscore() {

	//get list from API
	score := getHighscore()

	//sort by score
	sort.Slice(score, func(i, j int) bool {
		return score[i].Score > score[j].Score
	})

	//tablewriter to present a pretty highscore list
	var data [][]string
	for _, v := range score {
		var scores []string
		scores = append(scores, v.Username, strconv.Itoa(v.Score))
		data = append(data, scores)
	}
	highscore := tablewriter.NewWriter(os.Stdout)
	highscore.SetHeader([]string{"Namn", "Score"})
	highscore.AppendBulk(data)
	highscore.Render()
}

//calculates how many percentage of earlier players you were better than
func calculatePerformance(score int) int {
	// 	P = (n/N) × 100
	// Where,
	// P is percentile
	// n – Number of values below ‘x’
	// N – Total count of population
	hs := getHighscore()

	var lowerScore int

	//check numbers below actual score
	for _, v := range hs {
		if v.Score > score {
			lowerScore++
		}
	}

	//returns which percentile you are part of
	percentile := int(math.Round((float64(lowerScore) / float64(len(hs))) * 100))

	//subtracts from 100% to see how many percent of total players you were better than
	percentage := 100 - percentile
	return percentage

}
func init() {
	rootCmd.AddCommand(highscoreCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// highscoreCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// highscoreCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
