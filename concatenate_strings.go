package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func inefficient_conc() string {
	var s string
	for i := 1; i < len(os.Args); i++ {
		s = s + os.Args[i] + " "
	}
	return s
}

func efficient_join() string {
	return strings.Join(os.Args[1:], " ")
}

func benchmark(fn func() string) int64 {
	start := time.Now()
	for i := 1; i < 5000000; i++ {
		fn()
	}
	return time.Since(start).Milliseconds()
}

func benchmark_string_concat() {
	fmt.Println(benchmark(inefficient_conc))
	fmt.Println(benchmark(efficient_join))
}
