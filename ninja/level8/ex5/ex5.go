package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByAge []user

func (users ByAge) Len() int {
	return len(users)
}

func (users ByAge) Less(i int, j int) bool {
	return users[i].Age < users[j].Age
}

func (users ByAge) Swap(i int, j int) {
	users[i], users[j] = users[j], users[i]
}

type ByLast []user

func (users ByLast) Len() int {
	return len(users)
}

func (users ByLast) Less(i, j int) bool {
	return users[i].Last < users[j].Last
}

func (users ByLast) Swap(i int, j int) {
	users[i], users[j] = users[j], users[i]
}
func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	usersByAge := ByAge{u1, u2, u3}
	usersByLast := ByLast{u1, u2, u3}

	fmt.Println("\nUNSORTED:")
	printUsers([]user{u1, u2, u3})

	fmt.Println("\nSORTED by age: ")

	sort.Sort(usersByAge)
	printUsers(usersByAge)

	fmt.Println("\nSORTED by last name: ")

	sort.Sort(usersByLast)
	printUsers(usersByLast)
}

func printUsers(usersByAge []user) {
	for _, v := range usersByAge {
		fmt.Println("I'm", v.First, v.Last, "and I'm", v.Age, "years old")
		for _, v := range v.Sayings {
			fmt.Println("\t", v)
		}
	}
}
