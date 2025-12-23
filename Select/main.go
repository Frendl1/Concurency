package main

//TODO: Сделай два канала — один тикер (`time.Tick`), другой "stop".
//TODO: В цикле через `select` выводи “tick”, пока не придёт stop.

import (
	"fmt"
	"time"
)

func main(){
	ticker:= time.NewTicker(time.Second*6) 	//Создаём тикер
	defer ticker.Stop()						//Выключим его после окончания мэин
	done:= make(chan bool)					// Создаём канал приёмник
	go func(){								//Имитипуем работу
		time.Sleep(time.Second*5)
		done<-true 
	}()
	for {									//Запускаем цикл для селект
		select{								
		case <-done:						//Вариант с выполнением задачи
			fmt.Println("Done!")
			return
		case t:=<-ticker.C:					//Вариант с выводом текущего времени
			fmt.Println("Current time: ", t)
			return
		}
	}

}