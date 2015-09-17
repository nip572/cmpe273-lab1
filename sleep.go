package main

import "time"
import "fmt"

func Sleep(q int) {

	<-time.After(time.Second * time.Duration(q))

}

func main() {

	t1 := time.Now().Second()
	fmt.Println(t1)
	Sleep(5)
	t2 := time.Now().Second()
	fmt.Println(t2)

	if t2-t1 == 5 {

		fmt.Println("You are done")
	}
}