package main

import (
	"fmt"
)

func main() {
	jobsFib := make(chan int, 100)
	resultsFib := make(chan int, 100)
	jobs := make(chan int64, 100)
	results := make(chan int64, 100)

	go workerFactorial(jobs, results)

	go worker(jobsFib, resultsFib)
	go worker(jobsFib, resultsFib)
	go worker(jobsFib, resultsFib)
	go worker(jobsFib, resultsFib)
	go worker(jobsFib, resultsFib)
	go worker(jobsFib, resultsFib)
	go worker(jobsFib, resultsFib)
	go worker(jobsFib, resultsFib)
	go worker(jobsFib, resultsFib)
	go worker(jobsFib, resultsFib)
	go worker(jobsFib, resultsFib)
	go worker(jobsFib, resultsFib)

	counter := 50
	for i := 0; i < counter; i++ {
		jobs <- int64(i)
		jobsFib <- i
	}
	close(jobs)

	for j := 0; j < counter; j++ {
		// select {
		// case <-results:
		// 	fmt.Println("and FAC result of", j, "is:", <-results)

		// case <-resultsFib:
		// 	fmt.Println("and FIB result of", j, "is:", <-resultsFib)
		// }

		//fmt.Println("and FAC result of", j, "is:", <-results)
		fmt.Println("and FIB result of", j, "is:", <-resultsFib)
	}
}

func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

func workerFactorial(jobs <-chan int64, results chan<- int64) {
	for n := range jobs {
		results <- factorial(int64(n))
	}
}
