package main

import (
	"fmt"
	"sync"
	"time"
)

// (1) use goroutine
func hello(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep((100 * time.Millisecond))
		fmt.Println(s)
	}
}

// (2) multi goroutine
var start time.Time

func init() {
	start = time.Now()
}

func getChars(sentence string) {
	for _, c := range sentence {
		fmt.Printf("%c at time %v\n", c, time.Since(start))
		time.Sleep(10 * time.Millisecond)
	}
}

func getDigits(digits []int) {
	for _, d := range digits {
		fmt.Printf("%d at time %v\n", d, time.Since(start))
		time.Sleep(20 * time.Millisecond)
	}
}

// (3) use channel
func greet(c chan string) {
	data := <-c //storing data from channel c
	fmt.Println("hello", data)
}

// (6) select
func getAverage(numbers []int, ch chan float64) {
	sum := 0

	for _, e := range numbers {
		sum += e
	}

	ch <- float64(sum) / float64(len(numbers))
}

func getMax(numbers []int, ch chan int) {
	max := numbers[0]

	for _, e := range numbers {
		if max < e {
			max = e
		}
	}

	ch <- max
}

// (7) Race Condition
// cek dengan go run -race main.go
func getNumbers() int {
	// data i di perebutkan untuk di print atau di isi 5
	var i int

	go func() {
		i = 5
	}()

	return i
}

// (8) fix race codition with waitgroups
func getNumbersWaitgroups() int {
	// data i di perebutkan untuk di print atau di isi 5
	var i int
	var wg sync.WaitGroup
	wg.Add(1) //menambahkan 1 process

	go func() {
		i = 5
		wg.Done() // menandakan wg selesai
	}()

	// menunggu wg selesai
	wg.Wait()

	return i
}

// (8) fix race codition with channel blocking
func getNumbersChannelBlocking() int {
	var i int

	done := make(chan interface{})
	go func() {
		i = 5
		done <- struct{}{} //isi struct kosong agar tak muncul apapun
	}()

	// menunggu hingga channel diisi
	<-done
	return i
}

// (8) fix race codition with mutex
type safeNumber struct {
	i int
	m sync.Mutex
}

func (safeNumber *safeNumber) Set(number int) {
	safeNumber.m.Lock()
	defer safeNumber.m.Unlock() //defer jalankan terakhir
	safeNumber.i = number       // isi data i
}

func (safeNumber *safeNumber) Get() int {
	safeNumber.m.Lock()
	defer safeNumber.m.Unlock()
	return safeNumber.i
}

func getNumbersMutex() int {
	var i = safeNumber{}

	go func() {
		i.Set(5)
	}()

	time.Sleep(1 * time.Second) //untuk menunggu diisi
	return i.Get()
}

// main salah satu adalah go concurrency
func main() {
	// // (1) hello dengan go routine
	// go hello("world")
	// hello("hello")

	// // (2) multiple goroutine
	// fmt.Println("main start at time", time.Since(start))
	// fmt.Println()

	// // getchars goroutine
	// go getChars("helllo")
	// // getdigits goroutine
	// go getDigits([]int{1, 2, 3})

	// // schedule another goroutine
	// time.Sleep(200 * time.Millisecond)
	// fmt.Println("\nmain stopped at time", time.Since(start))

	// // (3) using channel
	// fmt.Println("main() started")
	// // membuat channel dengan tipe string
	// c := make(chan string)

	// go greet(c)

	// c <- "John" //write data to channel c
	// fmt.Println("main() stopped")

	// // (4) another use channel
	// theMine := [5]string{"ore1", "ore2", "ore3"}
	// oreChan := make(chan string)

	// // finder
	// go func(mine [5]string) {
	// 	for _, item := range mine {
	// 		oreChan <- item
	// 	}
	// }(theMine)

	// // ore breaker
	// go func() {
	// 	for i := 0; i < 3; i++ {
	// 		foundOre := <-oreChan
	// 		fmt.Println("Received", foundOre, "from finder")
	// 	}
	// }()

	// <-time.After(5 * time.Second) //again ignore this for now

	// // (5) buffered channel (FIFO) and without goroutine
	// message := make(chan string, 2) //buffered channel
	// message <- "buffered"
	// message <- "channel"

	// fmt.Println(<-message) // store
	// fmt.Println(<-message)

	// // (6) use select
	// numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}
	// chan1 := make(chan float64)

	// go getAverage(numbers, chan1)

	// chan2 := make(chan int)
	// go getMax(numbers, chan2)

	// for i := 0; i < 2; i++ {
	// 	// memilih chan mana yang akan dikeluarkan
	// 	select {
	// 	case avg := <-chan1:
	// 		fmt.Println("Average:", avg)
	// 	case max := <-chan2:
	// 		fmt.Println("Max:", max)
	// 	}
	// }

	// // (7) Race Condotion
	// fmt.Println(getNumbers())

	// // (8) fix race codition with waitgroups
	// fmt.Println(getNumbersWaitgroups())

	// // (8) fix race codition with channel blocking
	// fmt.Println(getNumbersChannelBlocking())

	// (8) fix race codition with mutex
	fmt.Println(getNumbersMutex())
}
