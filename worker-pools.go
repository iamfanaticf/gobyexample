package main

import(
  "fmt"
  "time"
)

func worker(id int, jobs <- chan int, results chan <- int) {
  for j := range jobs {
    fmt.Println("worker", id, "started job", j)
    time.Sleep(time.Second)
    fmt.Println("worker", id, "finished job", j)
    results <- j * j
  }
}

func main() {
  numOfWorkers := 3
  numOfJobs := 7

  jobs := make(chan int, numOfJobs)
  results := make(chan int, numOfJobs)

  for i := range numOfWorkers {
    go worker(i, jobs, results)
  }

  for i := range numOfJobs {
    jobs <- i
  }
  close(jobs)
  
  for _ = range numOfJobs {
    <- results
  }
}
