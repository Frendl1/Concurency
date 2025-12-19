package main

//TODO:  Каналы – базовый обмен:
// - Создай канал `chan int`.
// - Запусти произвольное число горутин, которые шлют в канал числа.
// - В главной горутине читай из канала и выводи числа.
// - Закрой канал и корректно завершай чтение (`for v := range ch`).

import (
	"fmt"
	"time"
)

func sender(ch chan string, senderDone chan bool){
	for i:=1; i<4; i++{
		ch<-fmt.Sprintf("messege %d\n", i)
	}

	close(ch)
	time.Sleep(100*time.Millisecond)
	senderDone <- true
}

func receive(ch chan string, receiverDone chan bool){
	for msg:= range ch{
		fmt.Println("Received: ", msg)
	}
	receiverDone <- true

}

func main(){
	ch:=make(chan string)
	senderDone:= make(chan bool)
	receiverDone:=make(chan bool)

	go sender(ch, senderDone)
	go receive(ch, receiverDone)

	<- senderDone
	<- receiverDone

	fmt.Println("All messege sended")

}