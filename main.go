package main

import (
	"app/disjoint_set"
	"app/kruskal"
	"fmt"
)

func main() {
	playKruskal()
}

func playKruskal() {
	edges := []kruskal.Edge{
		{0, 1, 4},
		{0, 2, 3},
		{0, 3, 3},
		{1, 5, 1},
		{1, 6, 5},
		{2, 3, 2},
		{3, 7, 9},
		{5, 6, 3},
		{7, 1, 5},
	}

	fmt.Println("Edges:", edges)

	mst := kruskal.Kruskal(8, edges)

	fmt.Println("Minimum Spanning Tree:", mst)
}

func testDisjointSet() {

	ds := disjoint_set.NewDisjointSet(10)
	ds.Print()

	ds.Union(0, 1)
	ds.Print()

	ds.Union(2, 3)
	ds.Print()

	fmt.Println("Find parent 3 :", ds.Find(3))

	ds.Union(1, 2)
	ds.Print()

	fmt.Println("Find parent 3 :", ds.Find(3))

	fmt.Println(ds.GetSets())
	ds.PrintSets()
}
