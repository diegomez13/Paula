package main

import "sync"

type Respuesta struct {
	Status bool
}

func main() {

	var wg sync.WaitGroup
	ch := make(chan Respuesta)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go Test(&wg, ch, i)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

}

func Test(wg *sync.WaitGroup, ch chan<- Respuesta, i int) {
	defer wg.Done()
	ch <- Respuesta{Status: true}
}
