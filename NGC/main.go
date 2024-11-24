package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
)

func processImage(imageURL string){
	// simulate image processing 
	fmt.Printf("Processing image: %s\n",imageURL)
	time.Sleep(3 * time.Second) // introduce 3 second delay using time.Sleep()
}

func main(){
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
	
	// Shuffle the order of the array to simulate "randomness" of array processing
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	rand.Shuffle(len(image_URLs), func(i, j int) {
		image_URLs[i], image_URLs[j] = image_URLs[j], image_URLs[i]
	})

	// add a waitGroup for image processing
	wg.Add(1)

	// output image processing completed for the above four URLs
	go func(){
		defer wg.Done()

		for i := 0; i < len(image_URLs); i++ {
			fmt.Printf("Image processing completed: %s\n", image_URLs[i])
			time.Sleep(1 * time.Second) // ensure 1 second delay with the next output
		}
	}()

	// wait for the final waitGroup to finish
	wg.Wait()
	
	// all image processing service have finished 
	fmt.Println("All image processing completed.")
}