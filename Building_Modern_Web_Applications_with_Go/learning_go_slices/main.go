package main

import "log"

func main() {
	var mySlice []string

	mySlice = append(mySlice, "Trevor")
	mySlice = append(mySlice, "John")
	mySlice = append(mySlice, "Mary")

	log.Println(mySlice)
}
