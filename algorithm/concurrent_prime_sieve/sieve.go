package concurrent_prime_sieve

import (
	"fmt"
	"sync"
)

func Start(primeRange int) {
	ch := make(chan int, primeRange)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go filter(ch, primeRange, wg)
	for i := 2; i <= primeRange; i++ {
		ch <- i
	}
	close(ch)
	wg.Wait()
}

func filter(ch <-chan int, primeRange int, wg *sync.WaitGroup) {
	defer wg.Done()
	prime, ok := <-ch
	if !ok {
		return
	}
	newChan := make(chan int, primeRange)
	wg.Add(1)
	go filter(newChan, primeRange, wg)
	fmt.Println("prime:", prime)
	for num := range ch {
		if num%prime != 0 {
			newChan <- num
		}
	}
	close(newChan)
}
