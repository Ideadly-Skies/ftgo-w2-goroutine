// go routine to send mail
package main

import (
	"fmt"
	"time"
	"sync"	
)

func sendEmail(){
	for i := 0; i < 2; i++ {
		// SMTP code to send email 
		fmt.Println("Sending email: ", i)	
		time.Sleep(1 * time.Second)	
	}
}

func sendSMS(){
	for i := 0; i < 2; i++ {
		// SMS code to send SMS
		fmt.Println("Sending SMS: ", i)
		time.Sleep(1 * time.Second)
	}
}

func storeData(){
	for i := 0; i < 5; i++ {
		// Database code to store data
		fmt.Println("Storing data: ", i)
		time.Sleep(1 * time.Second)
	}
}

func sendMail(){
	// init waitgroup 
	var wg sync.WaitGroup	
	
	// run the sendEmail process first
	wg.Add(1)
	go func() {
		defer wg.Done()
		sendEmail()
	}()	
	wg.Wait()
	
	// run the sendSMS process next
	wg.Add(1)
	go func(){
		defer wg.Done()
		sendSMS()
	}()
	wg.Wait()
	
	// store data run synchronously 
	storeData()

	time.Sleep(6 * time.Second)
}