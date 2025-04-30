package main

import (
	"log"
)

func main() {
	listNodes, primaryNode, arrivalPoints, err := readInput("input.log")
	if err != nil {
		log.Fatal("read input:", err)
	}

	var allPaths [][]string
	dfsAll(primaryNode, arrivalPoints, listNodes, []string{}, &allPaths, make(map[string]bool))
	log.Println("maximum number of passengers:", listNodes.maxPassengersCount(allPaths))

	log.Printf(
		"maximum number of passengers (ford fulkerson algorithm): %d", ff(primaryNode, listNodes, arrivalPoints),
	)
}
