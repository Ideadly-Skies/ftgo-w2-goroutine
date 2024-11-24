package main

import (
	"fmt"
	"sync"
	"time"
) 

func downloadFile(url string, i int) {
	fmt.Printf("Downloaded file from %s to video%d.mp4\n",url, i)
	time.Sleep(1 * time.Second)
}

func DownloadManager() {
	var download_urls = []string{
		"https://sample-videos.com/video123/mp4/720/big_buck_bunny_720p_1mb.mp4",
		"https://sample-videos.com/video123/mp4/720/big_buck_bunny_720p_2mb.mp4",
		"https://sample-videos.com/video123/mp4/720/big_buck_bunny_720p_3mb.mp4",
	}
	
	// declare WaitGroup for download_urls 
	var wg sync.WaitGroup

	// simulate download files as Goroutine process
	wg.Add(1)
	go func(){
		defer wg.Done()

		// download all files from download urls
		for i := 0; i < len(download_urls); i++{
			downloadFile(download_urls[i], i+1)
		}
	}()

	// wait for download goroutine to finish
	wg.Wait()
	
	// output all files downloaded
	fmt.Println("All files downloaded.")
}