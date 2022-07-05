package main

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

//questions to simulate database
var questions = []Question{

	{
		ID:            1,
		Question:      "Vilken färg har svenska flaggan?",
		Answers:       []string{"Gul + blå", "Röd + vit", "Svart + rosa", "Grön + brun"},
		CorrectAnswer: 1,
	},

	{
		ID:            2,
		Question:      "Hur många landskap finns det i Sverige?",
		Answers:       []string{"21", "24", "25", "28"},
		CorrectAnswer: 3,
	},

	{
		ID:            3,
		Question:      "Vilket år mördades Olof Palme?",
		Answers:       []string{"1989", "1986", "1982", "1984"},
		CorrectAnswer: 2,
	},

	{
		ID:            4,
		Question:      "Vad hette den ryska hunden som blev skickad ut i rymden år 1957?",
		Answers:       []string{"Balalajka", "Likea", "Lajka", "Taiga"},
		CorrectAnswer: 3,
	},
	{
		ID:            5,
		Question:      "Vilket år skedde den första månlandningen?",
		Answers:       []string{"1969", "1972", "1965", "1979"},
		CorrectAnswer: 1,
	},
	{
		ID:            6,
		Question:      "Vad heter Fantomens varg?",
		Answers:       []string{"Fido", "Brian", "Jesus", "Devil"},
		CorrectAnswer: 4,
	},
}

//answers which will expand when players are posting their answers
var answers = []Answer{
	{
		Username: "Jörgen (seed)",
		Score:    3,
	},
	{
		Username: "Sara (seed)",
		Score:    2,
	},
	{
		Username: "Lasse (seed)",
		Score:    4,
	},
}
