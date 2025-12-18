package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
    Запусти несколько горутин, чтобы выводить сообщения параллельно.
    Например, запуск `go func() { … }()` в цикле,
    и пусть каждая горутина печатает свой id + сообщение,
	потом главный поток ждёт (через `time.Sleep` или `WaitGroup`).

*/
func chat(i int, wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Printf("Привет я %d\n", i)
	time.Sleep(2*time.Second)
}
func main(){
	var wg sync.WaitGroup
	goroutine := rand.Intn(5)
	wg.Add(goroutine)
	for i := range goroutine{
		go chat(i, &wg)
	}
	wg.Wait()

}










/*

func randomWait(ch chan int, wg *sync.WaitGroup){
	defer wg.Done()
	//fmt.Println("Горутина начала работу")


	workSecond:=rand.Intn(5+1)
	time.Sleep(time.Duration(workSecond)*time.Second)


	//fmt.Println("Горутина закончила работу", workSecond)
	ch <- workSecond
}

func main(){
	wg:= &sync.WaitGroup{}
	totalWorkTime:= 0
	start:= time.Now()
	ch := make(chan int, 100)
	for range 100{
		wg.Add(1)
		go randomWait(ch, wg)
	}
	wg.Wait()
	close(ch)

	for workTime:= range ch{
		totalWorkTime += workTime
	}
	mainTime:= time.Since(start)
	fmt.Println("mainTime = ", mainTime.Seconds())
	fmt.Println("Total work time = ", totalWorkTime)

}
*/

