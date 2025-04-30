package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type nodes map[string]map[string]int

func (n nodes) maxPassengersCount(allPath [][]string) int {
	var (
		filter = map[string]int{}
		ret    int
	)

	for _, path := range allPath {
		quota := n.minQuota(path)

		tag := path[0] + path[1]
		if filter[tag] < quota {
			filter[tag] = quota
		}
	}

	for _, v := range filter {
		ret += v
	}

	return ret
}

func (n nodes) minQuota(path []string) int {
	ret := math.MaxInt
	for i := 1; i < len(path); i++ {
		if quota := n[path[i-1]][path[i]]; quota < ret {
			ret = quota
		}
	}

	return ret
}

func readInput(fileName string) (nodes, string, []string, error) {
	f, err := os.OpenInRoot("./", fileName)
	if err != nil {
		return nil, "", nil, fmt.Errorf("failed to open file: %w", err)
	}

	defer func(f *os.File) {
		if err = f.Close(); err != nil {
			log.Fatalf("failed to close file: %v", err)
		}
	}(f)

	var (
		primaryNode   string
		arrivalPoints []string
		listNodes     = make(nodes)
	)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if s := strings.Split(scanner.Text(), " "); len(s) > 5 {
			switch s[1] {
			case "TRANSMISSION:":
				q, _ := strconv.Atoi(s[8])
				if _, ok := listNodes[s[3]]; !ok {
					listNodes[s[3]] = make(map[string]int)
				}
				listNodes[s[3]][s[5]] = q
			case "ALERT:":
				primaryNode = s[5]
			case "CRITICAL:":
				for _, p := range s[6:] {
					arrivalPoints = append(arrivalPoints, strings.TrimSuffix(p, ","))
				}
			}
		}
	}

	if err = scanner.Err(); err != nil {
		return nil, "", nil, fmt.Errorf("failed to read file: %w", err)
	}

	return listNodes, primaryNode, arrivalPoints, nil
}

func dfsAll(src string, dest []string, listNodes nodes, path []string, allPaths *[][]string, visited map[string]bool) {
	path = append(path, src)
	visited[src] = true

	if slices.Contains(dest, src) {
		*allPaths = append(*allPaths, slices.Clone(path))
	}

	for adjNode, quota := range listNodes[src] {
		if quota == 0 || visited[adjNode] {
			continue
		}
		dfsAll(adjNode, dest, listNodes, path, allPaths, visited)
	}

	visited[src] = false
}
