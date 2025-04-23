package main

import (
	"fmt"
	"log"
)

func main() {
	listNodes, primaryNode, arrivalPoint, err := readInput("input.log")
	if err != nil {
		log.Fatal("read input:", err)
	}

	maxPassengers := 0

	var allPaths [][]string
	dfs(primaryNode, arrivalPoint, listNodes, []string{}, &allPaths, make(map[string]bool))
	if count := listNodes.maxPassengersCount(allPaths); maxPassengers < count {
		maxPassengers = count
	}

	fmt.Println(allPaths)
	fmt.Println(len(allPaths))

	log.Println("maximum number of passengers:", maxPassengers)
}
