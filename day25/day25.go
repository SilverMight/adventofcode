package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Edge struct {
    source, dest string
}

func getEdgeFrequency(graph map[string][]string) map[Edge]int {
    timesVisited := make(map[Edge]int)
    for start := range graph {
        bfs(graph, timesVisited, start)
    }

    fmt.Println(timesVisited)
    return timesVisited
}

func bfs(graph map[string][]string, timesVisited map[Edge]int, start string) {

    queue := []string{start}
    visited := make(map[string]bool)

    visited[start] = true
    for len(queue) > 0 {
        curr := queue[0]
        queue = queue[1:]


        for _, neighbor := range graph[curr] {
            if _, hasVisited := visited[neighbor]; !hasVisited {
                queue = append(queue, neighbor)
                visited[neighbor] = true
                // Order doesn't matter
                var key Edge
                if curr < neighbor {
                    key = Edge{curr, neighbor}
                } else {
                    key = Edge{neighbor, curr}
                }
                timesVisited[key]++
            }
        }


    }
}

func getMostVisitedEdge(edgeFrequency map[Edge]int) Edge  {
    best := -1
    var ret Edge 
    for edge, freq := range edgeFrequency {
        if freq > best {
            ret = edge 
            best = freq
        }
    }

    return ret
}

func removeEdge(graph map[string][]string, edge Edge) {
    destIndex := slices.Index(graph[edge.source], edge.dest)
    graph[edge.source] = slices.Delete(graph[edge.source], destIndex, destIndex + 1)

    sourceIndex := slices.Index(graph[edge.dest], edge.source)
    graph[edge.dest] = slices.Delete(graph[edge.dest], sourceIndex, sourceIndex + 1)

}

func numNodes(graph map[string][]string, start string) (numNodes int) {
    queue := []string{start}
    visited := make(map[string]bool)
    visited[start] = true

    for len(queue) > 0 {
        curr := queue[0]
        queue = queue[1:]


        for _, neighbor := range graph[curr] {
            if _, hasVisited := visited[neighbor]; !hasVisited {
                queue = append(queue, neighbor)
                visited[neighbor] = true

            }
        }
        numNodes++


    }
    return
}

func parse(fileName string) map[string][]string {
    file, err := os.Open(fileName)
    if err != nil {
        return nil
    }
    scanner := bufio.NewScanner(file)

    graph := make(map[string][]string)

    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.Split(line, ":")

        node := strings.TrimSpace(parts[0])
        neighbors := strings.Fields(parts[1])

        for _, neighbor := range neighbors {
            graph[node] = append(graph[node], neighbor)
            graph[neighbor] = append(graph[neighbor], node)
        }
    }

    return graph
}

func main()  {
    adjacencyMap := parse("input.txt") 

    var firstEdgeSource string
    
    // Remove most visited edges 3 times
    for i := 0; i < 3; i++ {
        edgeFrequency := getEdgeFrequency(adjacencyMap)
        maxEdge := getMostVisitedEdge(edgeFrequency)
        if i == 0 {
            firstEdgeSource = maxEdge.source
        }
        removeEdge(adjacencyMap, maxEdge)
    }

    group1 := numNodes(adjacencyMap, firstEdgeSource)
    group2 := len(adjacencyMap) - group1
    fmt.Println(group1 * group2)
    
}
