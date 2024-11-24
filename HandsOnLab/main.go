package main

import (
	"fmt"
	"sync"
	"time"
)

// function to send email to five emails (hard-coded)
func sendEmail(){
	var emails = [5]string{"obie.kal22@gmail.com", "swaggykho@gmail.com", "nd.jyakarta@gmail.com", "maddison12@gmail.com", "randalf_mckinsey@gmail.com"}
	for i := 0; i < len(emails); i++ {
		fmt.Printf("Process %d sending email to: %s\n", i+1, emails[i])
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// init WaitGroup
	var wg sync.WaitGroup		
	
	// add one waitgroup process for sendEmail function
	wg.Add(1)

	// make sendEmail function asynchronous
	go func(){
		defer wg.Done() // defer wg.done until end of function execution 
		sendEmail() 	// send email to five intended recipients
	}()	
		
	wg.Wait()		// wait for the waitgroup to finish
		
	// output success message for wg
	fmt.Println("email have been successfully sent out to intended recipient")
	time.Sleep(1 * time.Second)
}