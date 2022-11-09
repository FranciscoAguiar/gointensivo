package exemplo

import (
	"fmt"
	"time"
)

type Carro struct {
	Marca string
}

func (c *Carro) MudaMarca(marca string) {
	c.Marca = marca
	fmt.Println(c.Marca)
}

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running \n", i, name)
		time.Sleep(time.Second)
	}
}

func main() {
	fmt.Println("Hello, World!")
	a := 1
	fmt.Println(a)
	fmt.Println(&a)
	carro := Carro{Marca: "Fiat"}
	carro.MudaMarca("GM")
	fmt.Println(carro.Marca)
	go task("B") //goroutine - green threads
	go task("C")
	task("D")

	canal := make(chan string)
	//T2
	go func() {
		canal <- "veio da T2"
	}()

	//T1
	msg := <-canal
	fmt.Println(msg)

}
