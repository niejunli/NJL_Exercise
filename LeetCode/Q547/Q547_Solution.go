package Q547

/*
并查集 --- 省份数量
注意路径压缩
*/

type UnionFind struct {
    count int
    parent []int
}

func NewUnionFind(n int) *UnionFind {
    parent := make([]int, n)
    for i:=0; i<n; i++ {
        parent[i] = i // 初始化，每个城市的各自独立
    }
    return &UnionFind{
        count:n,
        parent:parent,
    }
}


func(u *UnionFind) Union(i, j int) {
    pi := u.Find(i) // 找到节点i的parent
    pj := u.Find(j) // 找到节点j的parenty
    if pi != pj {
        u.parent[pi] = pj
        u.count--
    }
}

func (u *UnionFind) Find(i int) int {
    root := i
    for u.parent[root] != root {
        root = u.parent[root] // 找到根节点，只有根节点的parent是自己
    }
    // for u.parent[i] != i { // 遍历至根节点
    //     i, u.parent[i] = u.parent[i], root // 路径压缩，让路径上的节点都挂在根节点上
    // }
    if u.parent[i] != i {
        u.parent[i] = root
    }
    return root
}

func findCircleNum(isConnected [][]int) int {
    if len(isConnected) == 0 {
        return 0
    }
    length := len(isConnected) // 城市个数
    ufind := NewUnionFind(length)

    for i:=0; i<length; i++ {
        for j:=i+1; j<length; j++ {
            if isConnected[i][j] == 1 {
                ufind.Union(i,j)
            }
        }
    }
    return ufind.count
}