package main

import (
	"fmt"
	"sync"
	"time"
)

func processImage(imageURL string){
	// simulate image processing 
	fmt.Printf("Processing image: %s\n",imageURL)
	time.Sleep(3 * time.Second) // introduce 3 second delay using time.Sleep()
}

func ProcessImageNGC(){
	// declare waitGroup for processImage process 
	var wg sync.WaitGroup

	// simulate multiple image uploads 
	image_URLs := [4]string{"https://example.com/image1.jpg", "https://example.com/image2.jpg", "https://example.com/image3.jpg", "https://example.com/image4.jpg"}

	// print out image processing started
	fmt.Println("Image processing started, main application continues...")

	// add a waitGroup for processImage
	wg.Add(1)	

	go func(){
		defer wg.Done()

		for i := 0; i < len(image_URLs); i++ {
			processImage(image_URLs[i])
		}
	}()
		
	// wait for the waitGroup to finish
	wg.Wait()

	// output image processing completed should be asynchronous in nature
	for i := 0; i < len(image_URLs); i++ {
		go fmt.Printf("Image processing completed: %s\n", image_URLs[i])
	}

	// invoke time.Sleep (hard-coded) for four URLs above
	time.Sleep(4 * time.Duration(len(image_URLs)))

	// all image processing service have finished 
	fmt.Println("All image processing completed.")
}