// Go routine example
package main

import (
	"fmt"
	"time"
)

func routine(){
	for i := 0; i < 5; i++ {
		fmt.Println("routine: ", i)
		time.Sleep(1 * time.Second)
	}
}

func main(){
	// start a goroutine - go [function_name]
	// func(){
	// 	for i := 0; i < 5; i++ {
	// 		fmt.Println("goroutine: ", i) 
	// 		time.Sleep(1 * time.Second) 	
	// 	}
	// }()

	// start a goroutine
	// go routine()	
	
	// Main Goroutine
	// for i := 0; i < 2; i++ {
	// 	fmt.Println("main: ", i)
	// 	// time.Sleep(1 * time.Second)
	// }

	// to output all of the output of routine
	// time.Sleep(6 * time.Second)

	sendMail()	
}