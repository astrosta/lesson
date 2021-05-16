package main

import (
	"fmt"
	"runtime"
	"time"
)

func main(){
	runtime.GOMAXPROCS(1)
	for i:= 0; i<300; i++{
		i:=i
		go func(){

			fmt.Println("A:", i)
		}()
	}

	time.Sleep(time.Hour)
}
