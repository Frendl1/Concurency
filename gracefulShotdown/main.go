package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func gracefulShotdon(){
	sigChan:= make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	timer:= time.After(time.Second*10)
	select{
	case <- timer:
		fmt.Println("times out")
		return
	case sig:=<-sigChan:
		fmt.Println("signal: ", sig)
	}
}



func main(){
	gracefulShotdon()
}



























/*func main(){
	timer:= time.After(time.Second*1)
	resultCh:=make(chan int)
	go func(){
		defer close(resultCh)
		for i:=1; i<=1000; i++{
			select{
			case <- timer: 
				fmt.Println("time Out")
				return
			default: 
				time.Sleep(time.Nanosecond)
				resultCh<-i
			}
		}
	}()
		for v:=range resultCh{
			fmt.Println(v)
		}

}*/