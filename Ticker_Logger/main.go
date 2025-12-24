package main

//TODO: **Проект: “Ticker Logger”**
//Сделай программу, которая:
//- каждую секунду печатает сообщение,
//- но автоматически завершает программу через 5 секунд (через select + time.After)

import (
	"fmt"
	"time"
)

func main(){
	ticker:=time.NewTicker(1*time.Second)
	defer ticker.Stop()
	done:=make(chan struct{})
	timeout:=time.After(5*time.Second)

	go func(){
		for {
			select{
			case t:=<-ticker.C:
				fmt.Println("tick ", t.Format("05"))
			case <-done:
				return
			}
		}
	}()

	<-timeout
	close(done)
	fmt.Println("Done")

}