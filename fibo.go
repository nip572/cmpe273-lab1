package main
 
import "fmt"
 
func main(){
var input int
fmt.Println("Please enter a number ")
fmt.Scanf("%d", &input)
fmt.Println(fibo(input))
}

func fibo(a int) int{
if a == 0 {
	return 0
} else if a == 1 {
	return 1
} else {
	return fibo(a-1) + fibo(a-2)
}

}