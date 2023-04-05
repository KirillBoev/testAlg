package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type User struct {
	FullName string
	Age      int
	Income   int
}

// User создай структуру User.

// generateAge сгенерируй возраст от 18 до 70 лет.
func generateAge() int {
	return randInt(18, 70)
}

func randInt(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}

// generateIncome сгенерируй доход от 0 до 500000.
func generateIncome() int {
	return rand.Intn(500000)
}

// generateFullName сгенерируй полное имя. например "John Doe".
func generateFullName() string {
	firstName := []string{"Max", "Bella", "Vici", "Nade", "Lion", "John"}
	lastName := []string{"Messi", "Owen", "Benald", "Lex", "Gates", "Mask"}
	fullName := firstName[rand.Intn(len(firstName)-1)] + " " + lastName[rand.Intn(len(lastName)-1)]
	// создай слайс с именами и слайс с фамилиями.
	// используй rand.Intn() для выбора случайного имени и фамилии.

	return fullName
}

func main() {

	// Сгенерируй 1000 пользователей и заполни ими слайс users.
	var users []User
	for i := 0; i <= 1000; i++ {
		temp := User{
			FullName: generateFullName(),
			Age:      generateAge(),
			Income:   generateIncome(),
		}
		users = append(users, temp)
	}

	// Выведи средний возраст пользователей.
	sumAge := 0
	for _, i := range users {
		sumAge += i.Age
	}
	midleAge := sumAge / 1000
	fmt.Println(midleAge)

	// Выведи средний доход пользователей.
	sumIncome := 0
	for _, i := range users {
		sumIncome += i.Income
	}
	midleIncome := sumIncome / 1000
	fmt.Println(midleIncome)

	// Выведи количество пользователей, чей доход превышает средний доход.
	countRichUser := 0
	for _, i := range users {
		if i.Income > midleIncome {
			countRichUser++
		}
	}
	fmt.Println(countRichUser)
	//отсортировать пользователей по возрасту, используя функцию sort.Slice
	sort.SliceStable(users, func(i, j int) bool {
		return users[i].Age < users[j].Age
	})
	fmt.Println(users)
}
