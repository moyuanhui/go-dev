package goroute

func Add(a int, b int, ch chan int)  {
	sum := a + b
	ch <- sum
}