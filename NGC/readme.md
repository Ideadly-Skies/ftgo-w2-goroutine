Scenario: Image Processing Service

The objective of this assignment is to implement a basic program in Golang that demonstrates asynchronous processing using Goroutines. You will simulate
a scenario where multiple tasks need to be processed concurrently in the background, without blocking the main application's execution. 

Objectives:

The objective of this quiz is to implement a basic program in Golang that demonstrates asynchronous processing using Goroutines. You will simulate a scenario
where multiple tasks need to be processed concurrently in the background, without blocking the main application's execution

Requirements:

1. Implement a function processImage(imageURL string) that takes the URL of an image as input and simulates image processing. For simplicity, the function
can print a message indiciating that the image is being processed and introduce a 3-second delay using time.sleep to simulate processing time. 

2. Create a Goroutine for each uploaded image to process it asynchronously in the background. The main application should not wait for the image processing to complete

3. In the main function, simulate multiple images uploads by calling the processImage function with different image URLs concurrently.

4. After starting the image processing Goroutines, the main application should continue executing and print a message like "Image processing started, main application continues..."

5. Simulate other operations or tasks that the main application might be performing using time.sleep

6. After the Goroutines complete processing the image (simulated by the 3-second delay in processImage), print a message like "All image processing completed"