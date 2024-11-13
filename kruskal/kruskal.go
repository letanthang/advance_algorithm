package kruskal

import (
	"app/disjoint_set"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"os"
	"sort"
)

// Định nghĩa cạnh trong đồ thị
type Edge struct {
	U, V, Weight int
}

func Kruskal(n int, edges []Edge) []Edge {
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Weight < edges[j].Weight
	})

	ds := disjoint_set.NewDisjointSet(n)
	var mst []Edge

	for _, e := range edges {
		if !ds.IsConnected(e.U, e.V) {
			mst = append(mst, e)
			ds.Union(e.U, e.V)
		}

		if len(mst) == n-1 {
			break
		}
	}

	return mst
}

func DisplayEdges(edges []Edge) {
	nodes := []opts.GraphNode{}
	for i := range edges {
		nodes = append(nodes, opts.GraphNode{Name: string(i)})
	}

	var links []opts.GraphLink

	for _, e := range edges {
		link := opts.GraphLink{Source: string(e.U), Target: string(e.V)}
		links = append(links, link)
	}

	graph := charts.NewGraph()
	graph.SetGlobalOptions(charts.WithTitleOpts(opts.Title{Title: "Graph Visualization"}))

	graph.AddSeries("graph", nodes, links).
		SetSeriesOptions(
			charts.WithGraphChartOpts(opts.GraphChart{Layout: "force"}),
			charts.WithLabelOpts(opts.Label{}))
	// Save as HTML
	f, _ := os.Create("graph.html")
	graph.Render(f)
}
