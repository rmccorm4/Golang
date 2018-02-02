package main
import (
	"fmt"
	"math/rand"
)

func main()	{
	// create int variable for user inputted guess
	var guess int
	fmt.Println("Guess a number between 1 and 1729:")
	_, err := fmt.Scan(&guess)

	var answer int = rand.Intn(1)
	fmt.Println("The answer was:", answer)

	if guess == answer {
		fmt.Println("You got it! Great job!")
	} else {
		fmt.Println("Better luck next time!")
	}
}
