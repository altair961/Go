package main

import "log"

func main() {
	myMap := make(map[string]string)
	myMap["dog"] = "Samson"
	myMap["other-dog"] = "Cassie"
	myMap["dog"] = "fido"

	log.Println(myMap["dog"])
	log.Println(myMap["other-dog"])
}
