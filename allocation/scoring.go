package main

import (
	"math"
)

var scoreFuncs = []ScoreFunc{reactions, money, value}

func scoreCalculation(resource Resource, user User) float64 {

	score := 1.0
	for _, scoreFunc := range scoreFuncs {
		score = score * scoreFunc(user, resource)
	}
	return score

}

func money(user User, resource Resource) float64 {
	return 1 - (math.Abs(resource.Value-user.Money) / 10)
}

func userTypeA(user User, resource Resource) float64 {
	if user.Type == TypeA && resource.Name == "random" {
		return 0
	}
	return 1
}

func reactions(user User, resource Resource) float64 {

	if resource.Impressions == 0 {
		return 1
	}
	return float64(resource.Consumptions) / float64(resource.Impressions)
}

func value(user User, resource Resource) float64 {

	return resource.Value
}
