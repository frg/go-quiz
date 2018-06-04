package models

import (
	"sort"
	"time"
)

// TODO: external and internal contracts should be split

// QuizLeaderboard Quiz leaderboard element
type QuizLeaderboard struct {
	User      string    `json:"user"`
	Score     int       `json:"score"`
	CreatedAt time.Time `json:"createdAt"`
}

// QuizAnswer Quiz submitted answer
type QuizAnswer struct {
	QuestionID  int       `json:"questionId"`
	AnswerIndex int       `json:"answerIndex"`
	AnsweredAt  time.Time `json:"answeredAt"`
}

// QuizUserSubmission Quiz user submission
type QuizUserSubmission struct {
	User    string       `json:"user"`
	Answers []QuizAnswer `json:"answers"`
	score   int
}

// QuizQuestion A quiz question w/ answers
type QuizQuestion struct {
	ID          string   `json:"id"`
	Question    string   `json:"question"`
	Answers     []string `json:"answers"`
	answerIndex int
}

// Quiz A list of quiz questions
type Quiz struct {
	Questions []QuizQuestion `json:"questions"`
}

var quizLeaderboard = []QuizLeaderboard{}
var quizAnswers = make(map[string]QuizUserSubmission)
var quiz = Quiz{
	Questions: []QuizQuestion{
		QuizQuestion{
			ID:       "bc899244-cb49-44dd-9a91-c6af237bd451",
			Question: "In which decade was the American Institute of Electrical Engineers (AIEE) founded?",
			Answers: []string{
				"1850s",
				"1880s", // Answer
				"1930s",
				"1950s",
			},
		},
		QuizQuestion{
			ID:       "1a8ebfa7-b482-4876-955d-a72f76d11ce5",
			Question: "What is part of a database that holds only one type of information?",
			Answers: []string{
				"Report",
				"Field", // Answer
				"Record",
				"File",
			},
		},
		QuizQuestion{
			ID:       "bdfec3d2-6432-4a4e-8044-d20207992684",
			Question: "'OS' computer abbreviation usually means ?",
			Answers: []string{
				"Order of Significance",
				"Open Software",
				"Operating System", // Answer
				"Optical Sensor",
			},
		},
		QuizQuestion{
			ID:       "cb065339-ba58-461c-9736-abf777ebb43f",
			Question: "In which decade with the first transatlantic radio broadcast occur?",
			Answers: []string{
				"1850s",
				"1860s",
				"1870s",
				"1900s", // Answer
			},
		},
	},
}

// Sorting
type byScore []QuizLeaderboard

func (a byScore) Len() int           { return len(a) }
func (a byScore) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byScore) Less(i, j int) bool { return a[i].Score < a[j].Score }

// GetQuiz Get quiz data
func GetQuiz() Quiz {
	return quiz
}

// SaveUserAnswer Save user answers
func SaveUserAnswer(sub QuizUserSubmission) {
	// TODO: calculate score
	var score = 0

	quizAnswers[sub.User] = QuizUserSubmission{
		User:    sub.User,
		Answers: sub.Answers,
		score:   score,
	}

	quizLeaderboard = append(quizLeaderboard, QuizLeaderboard{
		User:      sub.User,
		Score:     score,
		CreatedAt: time.Now().UTC(),
	})

	sort.Sort(byScore(quizLeaderboard))
}

// GetQuizLeaderboard Get quiz leader board
func GetQuizLeaderboard() []QuizLeaderboard {
	return quizLeaderboard
}

// GetRelativeQuizLeaderboard Get quiz user relative stat
func GetRelativeQuizLeaderboard(user string) string {
	return "A relative message"
}
