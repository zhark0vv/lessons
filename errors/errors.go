package main

import (
	"errors"
	"fmt"
	"log"
)

// Custom error declarations
var ErrSimple = errors.New("a simple error occurred")
var ErrPanic = errors.New("something caused a panic")

func main() {
	// Example of handling an error
	err := wrappedError(errors.New("something went wrong"))
	if err != nil {
		// Checking if the error matches the specific error using errors.Is
		if errors.Is(err, ErrSimple) {
			fmt.Println("Handled simple error:", err)
		} else {
			fmt.Println("An unknown error occurred:", err)
		}
	}

	// Example of handling a panic
	fmt.Println("Starting panic example:")
	handlePanic()

	// Example of handling a panic caused by log.Panic
	fmt.Println("Starting log.Panic example:")
	handlePanicWithLogPanic()

	// Example of log.Fatal
	fmt.Println("Starting log.Fatal example:")
	logFatalExample()
}

// Function that returns a wrapped error
func wrappedError(err error) error {
	// Wrapping the custom error with additional context
	return fmt.Errorf("additional context: %w: %w", ErrSimple, err)
}

// Function that causes a panic
func causePanic() {
	// Panicking with a custom error
	panic(ErrPanic)
}

// Function that handles a panic using recover
func handlePanic() {
	defer func() {
		// Recover from panic and check if it's the expected error
		if r := recover(); r != nil {
			if errors.Is(r.(error), ErrPanic) {
				fmt.Println("Caught a panic:", r)
			} else {
				fmt.Println("Unknown panic:", r)
			}
		}
	}()

	causePanic()
	fmt.Println("This code will not be executed if a panic occurs.")
}

// Example of using log.Panic with recover
func handlePanicWithLogPanic() {
	defer func() {
		// Recover from panic triggered by log.Panic
		if r := recover(); r != nil {
			// Check if recovered value is an error and matches ErrPanic
			if err, ok := r.(error); ok && errors.Is(err, ErrPanic) {
				fmt.Println("Caught a panic from log.Panic:", err)
			} else {
				// log.Panic is unknown because it doesn't return an error, it raised with a string
				fmt.Println("Unknown panic from log.Panic:", r)
			}
		}
	}()

	log.Panic(ErrPanic)
	fmt.Println("This code will not be executed after log.Panic.")
}

// Example of using log.Fatal
func logFatalExample() {
	// log.Fatal prints the message and exits the program
	log.Fatal("This is a fatal error, the program will exit immediately")
	fmt.Println("This code will not be executed after log.Fatal.")
}
