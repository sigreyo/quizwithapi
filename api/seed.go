package main

type Question struct {
	ID            int       `json:"id"`
	Question      string    `json:"question"`
	Answers       [4]string `json:"answers"`
	CorrectAnswer int       `json:"correctanswer"`
}

var questions = map[int][4]Question{
	1: {
		Question{
			ID:            1,
			Question:      "Vilken färg har svenska flaggan?",
			Answers:       [4]string{"Gul + blå", "Röd + vit", "Svart + rosa", "Grön + brun"},
			CorrectAnswer: 1,
		},
	},
	2: {
		Question{
			ID:            1,
			Question:      "",
			Answers:       [4]string{""},
			CorrectAnswer: 1,
		},
	},
	3: {
		Question{
			ID:            1,
			Question:      "Vilken färg har svenska flaggan?",
			Answers:       [4]string{""},
			CorrectAnswer: 1,
		},
	},
	4: {
		Question{
			ID:            1,
			Question:      "Vilken färg har svenska flaggan?",
			Answers:       [4]string{""},
			CorrectAnswer: 1,
		},
	},
}
