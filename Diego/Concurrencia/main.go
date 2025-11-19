package main

import (
	"fmt"
	"sync"
	"time"
)

type Respuesta struct {
	Status bool
	Valor  int
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

	for x := range ch {
		fmt.Println(x)
	}

}
func Silence(r Respuesta) {

}
func Test(wg *sync.WaitGroup, ch chan<- Respuesta, i int) {
	defer wg.Done()
	time.Sleep(time.Duration(i*3) * time.Second)
	ch <- Respuesta{Status: true, Valor: i}
}
