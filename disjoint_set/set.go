package disjoint_set

import "fmt"

// Cấu trúc dữ liệu Disjoint Set với Path Compression và Union by Rank
type DisjointSet struct {
	parent, rank []int
}

// Khởi tạo Disjoint Set với n phần tử
func NewDisjointSet(n int) *DisjointSet {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i // Ban đầu, mỗi phần tử là cha của chính nó
	}
	return &DisjointSet{parent, rank}
}

// Hàm Find với Path Compression
func (ds *DisjointSet) Find(x int) int {
	if ds.parent[x] != x {
		ds.parent[x] = ds.Find(ds.parent[x]) // Path Compression
	}
	return ds.parent[x]
}

// Hàm Union với Union by Rank
func (ds *DisjointSet) Union(x, y int) {
	rootX := ds.Find(x)
	rootY := ds.Find(y)
	if rootX != rootY {
		// Union by Rank
		if ds.rank[rootX] < ds.rank[rootY] {
			ds.parent[rootX] = rootY
		} else if ds.rank[rootX] > ds.rank[rootY] {
			ds.parent[rootY] = rootX
		} else {
			ds.parent[rootY] = rootX
			ds.rank[rootX]++
		}
	}
}

func (ds *DisjointSet) IsConnected(x, y int) bool {
	return ds.Find(x) == ds.Find(y)
}

func (ds *DisjointSet) Print() {
	fmt.Println(ds.parent)
	fmt.Println(ds.rank)
	fmt.Println("----------------------------")
}

func (ds *DisjointSet) GetSets() [][]int {
	var mapSets = make(map[int][]int)
	for i := 0; i < len(ds.parent); i++ {
		root := ds.Find(i)
		if _, ok := mapSets[root]; ok {
			mapSets[root] = append(mapSets[root], i)
		} else {
			mapSets[root] = []int{i}
		}
	}

	var sets [][]int
	for _, v := range mapSets {
		sets = append(sets, v)
	}

	return sets
}

func (ds *DisjointSet) PrintSets() {
	sets := ds.GetSets()

	for i := 0; i < len(sets); i++ {
		fmt.Println("root:", sets[i][0], "members:", sets[i])
	}
}
