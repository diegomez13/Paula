package main

import (
	"sync"
	"time"
)

type Respuesta struct {
	Status bool
	Valor  int
	Indice int
}

func main() {

	var wg sync.WaitGroup
	ch := make(chan []Respuesta)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go Test2(&wg, ch, i)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for x := range ch {
		Silence(x)
	}

}
func Silence(a []Respuesta) {

}
func Test(wg *sync.WaitGroup, ch chan<- Respuesta, i int) {
	defer wg.Done()
	time.Sleep(time.Duration(i*3) * time.Second)
	ch <- Respuesta{Status: true, Valor: i}
}

func Test2(wg *sync.WaitGroup, ch chan<- []Respuesta, i int) {
	defer wg.Done()
	var Total int
	for j := 0; j < (9-i)*1000000000; j++ {
		Total = Total + j
	}

	ch <- []Respuesta{Respuesta{Status: true, Valor: Total, Indice: i}, Respuesta{Status: true, Valor: Total, Indice: i}}

}
