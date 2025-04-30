package main

// Fulkerson algorithm implementation
func ff(primaryNode string, listNodes nodes, arrivalPoints []string) int {
	res := 0

	for _, ap := range arrivalPoints {
		for {
			path := dfs(primaryNode, ap, listNodes, nil, make(map[string]bool))
			if len(path) == 0 {
				break
			}

			mq := listNodes.minQuota(path)
			for i := 1; i < len(path); i++ {
				listNodes[path[i-1]][path[i]] -= mq
				listNodes[path[i]][path[i-1]] += mq
			}
			res += mq
		}
	}

	return res
}

func dfs(src string, dest string, listNodes nodes, path []string, visited map[string]bool) []string {
	path = append(path, src)

	if src == dest {
		return path
	}

	visited[src] = true

	for adjNode, quota := range listNodes[src] {
		if quota == 0 || visited[adjNode] {
			continue
		}
		if res := dfs(adjNode, dest, listNodes, path, visited); res != nil {
			return res
		}
	}

	return nil
}

// func bfs(src, dst string, listNodes nodes) []string {
// 	type vertex struct {
// 		name      string
// 		traversal []string
// 	}
//
// 	visited := map[string]bool{}
//
// 	queue := []vertex{{
// 		name:      src,
// 		traversal: []string{src},
// 	}}
//
// 	for len(queue) != 0 {
// 		v := queue[0]
// 		queue = queue[1:]
//
// 		if v.name == dst {
// 			return v.traversal
// 		}
//
// 		for next, quota := range listNodes[v.name] {
// 			if quota == 0 || visited[next] {
// 				continue
// 			}
// 			visited[next] = true
//
// 			queue = append(queue, vertex{
// 				name:      next,
// 				traversal: append(slices.Clone(v.traversal), next),
// 			})
// 		}
// 	}
//
// 	return nil
// }
