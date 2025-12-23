package main

import "fmt"

//TODO: Сделай канал с буфером, отправь в него 5 чисел без отдельной горутины. Пойми, почему программа не зависает.

func main(){
	ch:=make(chan int, 3)
	ch<-1
	fmt.Println("первое сообщение записалось")
	ch<-2
	fmt.Println("второе сообщение записалось")
	empty:= <-ch
	fmt.Println("канал запихнули в переменную")
	ch<-3
	fmt.Println("3 сообщение записалось")
	ch<-4
	fmt.Println("4 сообщение записалось")
	reload:=<-ch
	fmt.Println("reload")

	ch<-5
	fmt.Println("5 сообщение отправилось")
	reload2 :=<-ch
	ch<-6
		fmt.Println("6 сообщение отправилось")



	fmt.Println(empty, <-ch,reload,reload2, <-ch, <-ch)
}