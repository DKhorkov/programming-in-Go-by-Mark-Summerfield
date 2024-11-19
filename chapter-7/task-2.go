package chapter_7

import (
	"fmt"
	"image"
	"os"
	"path/filepath"
	"runtime"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

func RunTask2(files []string) {
	workers := runtime.GOMAXPROCS(runtime.NumCPU()) // Use all the machine's cores
	jobs := make(chan string, workers*4)
	done := make(chan struct{}, workers)
	results := make(chan string, len(files))
	go prepareJobs(jobs, files)
	for range workers {
		go process(done, results, jobs)
	}

	waitResults(workers, done, results)
	printResults(results)
}

func prepareJobs(jobs chan<- string, files []string) {
	for _, filename := range files {
		jobs <- filename
	}
	close(jobs)
}

func process(done chan<- struct{}, results chan<- string, jobs <-chan string) {
	for filename := range jobs {
		if info, err := os.Stat(filename); err != nil ||
			(info.Mode()&os.ModeType != 0) {
			fmt.Println(err)
			return // Ignore errors and nonregular files
		}

		file, err := os.Open(filename)
		if err != nil {
			return // Ignore errors
		}
		defer file.Close()

		config, _, err := image.DecodeConfig(file)
		if err != nil {
			return // Ignore errors
		}

		results <- fmt.Sprintf(
			`<img src="%s" width="%d" height="%d" />`,
			filepath.Base(filename),
			config.Width,
			config.Height,
		)
	}

	done <- struct{}{}
}

func waitResults(workers int, done <-chan struct{}, results chan string) {
	for i := workers; i > 0; i-- {
		<-done
		workers--
	}

	close(results)
}

func printResults(results <-chan string) {
	for result := range results {
		fmt.Println(result)
	}
}
