package models

import (
	"fmt"
	"sort"
	"strconv"
	"time"
)

// QuizLeaderboard Quiz leaderboard element
type QuizLeaderboard struct {
	User      string    `json:"user"`
	Score     int       `json:"score"`
	CreatedAt time.Time `json:"createdAt"`
}

// QuizUserSubmission Quiz user submission
type QuizUserSubmission struct {
	User    string         `json:"user"`
	Answers map[string]int `json:"answers"`
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
			answerIndex: 1,
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
			answerIndex: 1,
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
			answerIndex: 2,
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
			answerIndex: 3,
		},
	},
}

// Sorting
type byScore []QuizLeaderboard

func (a byScore) Len() int           { return len(a) }
func (a byScore) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byScore) Less(i, j int) bool { return a[i].Score > a[j].Score }

// GetQuiz Get quiz data
func GetQuiz() Quiz {
	return quiz
}

// SaveUserAnswer Save user answers
func SaveUserAnswer(sub QuizUserSubmission) {
	// Calculate score
	var score, ptPerAnswer = 0, (100 / len(quiz.Questions))

	for _, element := range quiz.Questions {
		if sub.Answers[element.ID] == element.answerIndex {
			score += ptPerAnswer
		}
	}

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
	if len(quizLeaderboard) > 10 {
		return quizLeaderboard[:10]
	}

	return quizLeaderboard
}

// GetRelativeQuizLeaderboard Get quiz user relative stat
func GetRelativeQuizLeaderboard(user string) []string {
	var messages = []string{}

	if len(messages) > 10 {
		messages = append(messages, "There has been "+strconv.Itoa(len(messages))+" quizes completed!")
	}

	// TODO: Clean this up!
	var userTopScore, scores = 0, []QuizLeaderboard{}
	for k, v := range quizAnswers {
		// Find user's TOP score
		if k == user && v.score > userTopScore {
			userTopScore = v.score
		}

		// Convert into array QuizLeaderboard format
		scores = append(scores, QuizLeaderboard{
			Score: v.score,
			User:  v.User,
		})
	}

	// Sort by score
	sort.Sort(byScore(scores))

	// Find the index at which the Score is less than top of user
	var indexOfLessPts = 1
	for index, element := range scores {
		if element.Score < userTopScore {
			// This is the index that the User is better than!
			indexOfLessPts = index + 1
			break
		}
	}

	fmt.Println(strconv.Itoa(len(scores)))
	fmt.Println(strconv.Itoa(indexOfLessPts))
	fmt.Println(strconv.Itoa((indexOfLessPts / len(scores)) * 100))

	var scoreCount = len(scores)
	messages = append(messages, "The user is "+strconv.Itoa((indexOfLessPts/scoreCount)*100)+"% better than all users!")

	return messages
}
