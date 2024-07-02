package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()	
	launchDateTime := time.Date(2009, time.November, 10, 23, 0,0,0,time.UTC)

	fmt.Printf("Time is %v\n", now)
	fmt.Printf("Go launched at %v\n", launchDateTime)
	fmt.Printf("Go launched at %v\n", launchDateTime.Format(time.ANSIC))
	fmt.Printf("Go launched at %v\n", launchDateTime.Format(time.RFC1123))
}

