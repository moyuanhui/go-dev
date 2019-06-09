package main

func main() {
	var intLink Link
	for i := 0; i < 10; i++ {
		// intLink.InsertHead(i)
		intLink.InsertTail(i)
	}
	intLink.Trans()
}
