package main

import (
	"math/rand"
)

// ScoreFunc is the signature of all score functions
type ScoreFunc func(User, Resource) float64

// UserType defines the type of user
type UserType string

// TypeA is one type of user
const TypeA = UserType("A")

//TypeB is another type of user
const TypeB = UserType("B")

// User is a user in your system :)
type User struct {
	Money float64
	Type  UserType
}

// Resource is whatever we are trying to allocate
type Resource struct {
	Value        float64
	Consumptions int
	Impressions  int
	CTR          float64
	Name         string
}

// Use consumes the resource
func (r *Resource) Use() {

	r.Impressions = r.Impressions + 1
	if rand.Float64() < r.CTR {
		r.Consumptions = r.Consumptions + 1
	}
}

func randomUser() User {

	userType := TypeA
	if rand.Float32() < 0.4 {
		userType = TypeB
	}
	user := User{
		Money: rand.Float64() * 4,
		Type:  userType,
	}
	return user
}

func randomResource() Resource {

	res := Resource{
		Value: rand.Float64() * float64(rand.Int31n(9)),
		CTR:   rand.Float64(),
		Name:  "random",
	}
	return res
}

func initialResources() []Resource {

	cheapResource := Resource{
		Value: 0.05,
		CTR:   0.13,
		Name:  "cheap",
	}

	expensiveResource := Resource{
		Value: 4.10,
		CTR:   0.06,
		Name:  "expensive",
	}

	dealResource := Resource{
		Value: 2.20,
		CTR:   0.192,
		Name:  "deal",
	}

	resources := []Resource{cheapResource, expensiveResource, dealResource}

	return resources
}
