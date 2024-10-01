package main

import (
	"context"
	"fmt"
	"math"
	"runtime"
	"sync"
	"time"
)

// User used for storing id and name
type User struct {
	ID   int
	Name string
}

func main() {
	numUsers := 100000000
	fmt.Println("numUsers", numUsers)
	users := []User{}
	for i := 0; i < numUsers; i++ {
		u := User{
			ID:   i,
			Name: fmt.Sprintf("user_%d", i),
		}
		users = append(users, u)
	}
	userID := numUsers - 1

	// concurrent search
	begin := time.Now()
	// ctx, done := context.WithCancel(context.Background())
	numCPUs := runtime.NumCPU()
	fmt.Println("numCPU", numCPUs)
	resultChan := make(chan User)
	batchSize := int(math.Ceil(float64(len(users)) / float64(numCPUs)))
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer wg.Done()
		defer cancel()
		select {
		case res := <-resultChan:
			fmt.Printf("concurrency found user:%+v in:%s\n", res, time.Since(begin))
		case <-ctx.Done():
			fmt.Println("concurrency timeout, user not found within time")
		}
	}()
	wg.Add(numCPUs)
	for i := 0; i < numCPUs; i++ {
		startIndex := i * batchSize
		endIndex := startIndex + batchSize

		var slice []User
		if i < numCPUs-1 {
			slice = users[startIndex:endIndex]
		} else {
			slice = users[startIndex:]
		}
		go searchUser(&wg, slice, userID, resultChan)
	}
	wg.Wait()
}

func searchUser(wg *sync.WaitGroup, slice []User, id int, result chan<- User) {
	defer wg.Done()
	for _, user := range slice {
		if user.ID == id {
			result <- user
			break
		}
	}
}
