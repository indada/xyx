package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	mu.Lock()
	mu.Unlock()
	var rwm sync.RWMutex
	rwm.RLock()
	rwm.RUnlock()
	rwm.Lock()
	rwm.Unlock()

	var mg sync.WaitGroup
	mg.Add(1)
	mg.Done()
	mg.Wait()

	var once sync.Once
	once.Do(func() {
		fmt.Println(1)
	})
	g, _ := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return oo()
	})
	g.Go(func() error {
		return qq()
	})
	err := g.Wait()
	fmt.Println("err:", err)
	time.Sleep(10 * time.Second)
	fmt.Println("end")
}
func oo() error {
	time.Sleep(5 * time.Second)
	fmt.Println(5)
	return nil
}
func qq() error {
	time.Sleep(2 * time.Second)
	fmt.Println(2)
	return errors.New("2qq")
}
