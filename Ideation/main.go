package main

import (
	"fmt"
	"sync"
	"time"
)

func scheduleTask(task func()) {
	go task()
	time.Sleep(1 * time.Second) // simulate 1s time.Sleep() for next task
}

func task1(){
	fmt.Println("Task 1 is being executed.")
}

func task2(){
	fmt.Println("Task 2 is being executed.")
}

func task3(){
	fmt.Println("Task 3 is being executed.")
}

func main() {
	// call scheduleTask on each task
	scheduleTask(task1)
	scheduleTask(task2)
	scheduleTask(task3)
	
	fmt.Println("Main application continues...")
	
	// Homework: know how a WaitGroup work
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		// wait for the scheduleTask goroutine process to finish
		wg.Wait()
	}()

	// defer wg.Done until Goroutine process finishes 
	wg.Done()
	wg.Done()
	wg.Done()

	// outputs all tasks completed
	fmt.Println("All tasks completed.") 
}