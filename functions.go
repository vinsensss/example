package main

func fib(n int) int {
	if n <= 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

func factorial(n int64) int64 {
	if n <= 1 {
		return 1
	} else {
		return n * factorial(n-1)
	}

}
