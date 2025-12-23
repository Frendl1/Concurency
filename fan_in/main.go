package main

import (
	"fmt"
	"sync"
)
func MergeChanels[T any](chanels ...<-chan T)<-chan T{
	var wg sync.WaitGroup
	wg.Add(len(chanels))


	outputCh:= make(chan T)
	 for _, chanel:= range chanels{
		go func(){
			defer wg.Done()
			for value:= range chanel{
			outputCh<- value
		}
		}()
	 }
	go func ()  {
		wg.Wait()
		close(outputCh)	
	}()



	return outputCh
}

func main(){
	chanel:= make(chan int)
	chanel1:= make(chan int)
	chanel2:= make(chan int)

	go func(){
		defer func(){
			close(chanel)
			close(chanel1)
			close(chanel2)
		}()

		for i:=0; i<100; i+=3{
			chanel <- i
			chanel1 <- i+1
			chanel2 <- i+2
		}
	}()

	for value := range MergeChanels(chanel, chanel1, chanel2){
		fmt.Println(value)
	}


}