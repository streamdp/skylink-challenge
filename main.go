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
	dfs(primaryNode, arrivalPoints, listNodes, []string{}, &allPaths, make(map[string]bool))

	log.Println("maximum number of passengers:", listNodes.maxPassengersCount(allPaths))
}
