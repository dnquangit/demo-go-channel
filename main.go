package main

import (
	"fmt"
	"time"
)

func main() {
	///* Without receiver from main goroutine */
	//chanelAtMainRoutine := make(chan int)
	//chanelAtMainRoutine <- 5
	//fmt.Println("Completing ...")

	///* Without receiver from anonymous goroutine */
	//chanelAtMainRoutine := make(chan int)
	//go func() {
	//	chanelAtMainRoutine <- 5
	//}()
	//fmt.Println("Completing ...")

	///* With receiver from main goroutine */
	//chanelAtMainRoutine := make(chan int)
	//go func() {
	//	fmt.Println("Start getting value from channel ...")
	//	<-chanelAtMainRoutine
	//	fmt.Println("Complete getting value from channel ...")
	//}()
	//chanelAtMainRoutine <- 5
	//fmt.Println("Completing ...")

	/* Increase time at receiver goroutine */
	chanelAtMainRoutine := make(chan int)
	go doBigJob(chanelAtMainRoutine)
	chanelAtMainRoutine <- 5
	fmt.Println("Completing ...")
}

func doBigJob(ch <-chan int) {
	fmt.Println("Start getting value from channel ...")
	<-ch
	time.Sleep(time.Second * 1)
	fmt.Println("Complete getting value from channel ...")
}
