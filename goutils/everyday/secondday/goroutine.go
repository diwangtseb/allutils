package main

import "fmt"

func main(){
    mychan := make(chan int)
    mychan <- 1
    fmt.Println(<-mychan)
}
