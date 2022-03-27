package main

import (
	"log"
)

func GenerateData() <-chan int {
	oc := make(chan int)
	go func() {
		i := 0
		log.Println("123tttttt")
		for {
			oc <- i
			i = i + 1
			if i == 15 {
				break
			}
		}
		close(oc)
	}()
	return oc
}

func PrintData(ic <-chan int) <-chan int {
	log.Println("123123123")
	oc := make(chan int)

	go func() {
		for data := range ic {
			log.Println(data)
			oc <- data
		}
		close(oc)
	}()
	return oc
}

func TotalSum(ic <-chan int) {
	total := 0
	for data := range ic {
		total += data
	}
	log.Println("total: ", total)
}

func main() {
	TotalSum(PrintData(GenerateData()))
}
