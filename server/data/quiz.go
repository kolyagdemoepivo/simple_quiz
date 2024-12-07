package data

import "server/models"

var QuizData = []models.Quiz{{
	ID:    1,
	Title: "Awesome quiz",
	QuizItems: []models.QuizItem{{
		ID: 1,
		Question: models.Question{
			ID:   1,
			Text: "First question?",
		},
		Answers: []models.Answer{
			{
				ID:        1,
				Text:      "Wrong answer",
				IsCorrect: false,
			},
			{
				ID:        2,
				Text:      "Wrong answer",
				IsCorrect: false,
			}, {
				ID:        3,
				Text:      "Right answer",
				IsCorrect: true,
			}},
	},
		{
			ID: 2,
			Question: models.Question{
				ID:   2,
				Text: "Second question?",
			},
			Answers: []models.Answer{
				{
					ID:        4,
					Text:      "Right answer",
					IsCorrect: true,
				},
				{
					ID:        5,
					Text:      "Wrong answer",
					IsCorrect: false,
				}, {
					ID:        6,
					Text:      "Wrong answer",
					IsCorrect: false,
				}},
		},
		{
			ID: 3,
			Question: models.Question{
				ID:   3,
				Text: "Third question?",
			},
			Answers: []models.Answer{
				{
					ID:        7,
					Text:      "Wrong answer",
					IsCorrect: false,
				},
				{
					ID:        8,
					Text:      "Wrong answer",
					IsCorrect: false,
				}, {
					ID:        9,
					Text:      "Right answer",
					IsCorrect: true,
				}},
		},
		{
			ID: 4,
			Question: models.Question{
				ID:   4,
				Text: "Fourth question?",
			},
			Answers: []models.Answer{
				{
					ID:        10,
					Text:      "Wrong answer",
					IsCorrect: false,
				},
				{
					ID:        11,
					Text:      "Right answer",
					IsCorrect: true,
				}, {
					ID:        12,
					Text:      "Wrong answer",
					IsCorrect: false,
				}},
		},
		{
			ID: 4,
			Question: models.Question{
				ID:   5,
				Text: "Fifth question?",
			},
			Answers: []models.Answer{
				{
					ID:        13,
					Text:      "Wrong answer",
					IsCorrect: false,
				},
				{
					ID:        14,
					Text:      "Right answer",
					IsCorrect: true,
				}, {
					ID:        15,
					Text:      "Wrong answer",
					IsCorrect: false,
				}},
		}},
}}
