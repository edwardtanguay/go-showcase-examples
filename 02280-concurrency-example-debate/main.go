package main

import (
	"math/rand"
	"fmt"
	"time"
)

func speaker(name string, debate chan int) {
	for {
		microphone := <-debate // wait for microphone

		fmt.Printf("Turn %d: %s says \"%s\"\n", microphone, name, randomAnswer())
		time.Sleep(400 * time.Millisecond)

		microphone++
		debate <- microphone // give microphone back
	}
}

func randomAnswer() string {
	answers := []string{"Great.", "Ok, thanks.", "That's right.", "Super.", "Sounds good to me.", "We'll have to see."}
	return answers[rand.Intn(len(answers))]
}

func main() {
	debate := make(chan int)

	go speaker("Erik", debate)
	go speaker("Angela", debate)
	go speaker("Georg", debate)

	microphone := 1

	debate <-microphone // start debate

	time.Sleep(2 * time.Second)

	<-debate // end debate

}
