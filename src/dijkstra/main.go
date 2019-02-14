package main

import (
	"fmt"
	"math"
	"os"
)

const INF = math.MaxInt32

func dijkstra(graph [][]int, src int) {
	var (
		vertex = len(graph[0])
		dist   []int
		sptSet []bool
	)

	// Initialize dist and sptSet
	for i := 0; i < vertex; i++ {
		dist[i] = INF
		sptSet[i] = false
	}
	// src should be zero
	dist[src] = 0

	// Find shortest path to fill dist

}

func main() {

	var src int

	graph := [][]int{
		{0, 4, 0, 0, 0, 0, 0, 8, 0},
		{4, 0, 8, 0, 0, 0, 0, 11, 0},
		{0, 8, 0, 7, 0, 4, 0, 0, 2},
		{0, 0, 7, 0, 9, 14, 0, 0, 0},
		{0, 0, 0, 9, 0, 10, 0, 0, 0},
		{0, 0, 4, 14, 10, 0, 2, 0, 0},
		{0, 0, 0, 0, 0, 2, 0, 1, 6},
		{8, 11, 0, 0, 0, 0, 1, 0, 7},
		{0, 0, 2, 0, 0, 0, 6, 7, 0},
	}

	fmt.Println("Create default graph : ")
	for _, i1 := range graph {
		for _, i2 := range i1 {
			fmt.Printf("%3d ", i2)
		}
		fmt.Printf("\n")
	}

	fmt.Print("Enter the src : ")
	fmt.Scanf("%d", &src)
	if src >= len(graph[0]) {
		fmt.Fprintf(os.Stderr, "Invalid number : %v\n", src)
		os.Exit(1)
	}
	dijkstra(graph, src)
}
