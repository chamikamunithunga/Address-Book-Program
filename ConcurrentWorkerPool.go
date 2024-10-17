package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Task structure
type Task struct {
	ID int
}

// Worker function that processes tasks
func worker(id int, tasks <-chan Task, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		// Simulate task processing
		fmt.Printf("Worker %d is processing task %d\n", id, task.ID)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond) // Simulate varying work time
		fmt.Printf("Worker %d completed task %d\n", id, task.ID)
	}
}

// Function to create and assign tasks to the worker pool
func startWorkerPool(numWorkers int, numTasks int) {
	tasks := make(chan Task, numTasks)
	var wg sync.WaitGroup

	// Create workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, tasks, &wg)
	}

	// Send tasks to the task channel
	for i := 1; i <= numTasks; i++ {
		tasks <- Task{ID: i}
	}

	// Close task channel after all tasks are sent
	close(tasks)

	// Wait for all workers to finish
	wg.Wait()
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed random number generator

	// Number of workers and tasks
	numWorkers := 5
	numTasks := 10

	fmt.Printf("Starting worker pool with %d workers to process %d tasks...\n", numWorkers, numTasks)
	startWorkerPool(numWorkers, numTasks)
	fmt.Println("All tasks completed.")
}
