package soalprakerja

import (
	"fmt"
	"math"
)


type mystruct struct {
	fuel float64
	Type string
}

func (m *mystruct) cost() (float64, string) {
	fuelEfficiency := 1.5 // L/mill
	estimatedDistance := m.fuel * fuelEfficiency
	return estimatedDistance, m.Type
}

func EstimatedDistance (fuel float64, Type string) (float64, string) {
	m := mystruct{fuel :fuel, Type: Type}
	return m.cost()
}

type Student struct {
	Name  string
	Score float64
}

type Students []Student

func (s Students) AverageScore() float64 {
	totalScore := 0.0
	for _, student := range s {
		totalScore += student.Score
	}
	return totalScore / float64(len(s))
}

func (s Students) MinMaxScores() (float64, float64) {
	minScore := math.MaxFloat64
	maxScore := -math.MaxFloat64

	for _, student := range s {
		if student.Score < minScore {
			minScore = student.Score
		}
		if student.Score > maxScore {
			maxScore = student.Score
		}
	}

	return minScore, maxScore
}

func (s Students) PrintStatistics() {
	avgScore := s.AverageScore()
	minScore, maxScore := s.MinMaxScores()

	fmt.Printf("Average Score: %.2f\n", avgScore)
	fmt.Printf("Minimum Score: %.2f\n", minScore)
	fmt.Printf("Maximum Score: %.2f\n", maxScore)
}


