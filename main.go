package main

import (
	"fmt"
	"math/rand"
	"myquadtree/quad_tree"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	root := quad_tree.NewNode(1, quad_tree.NewRegion(0, 1000, 0, 1000))

	start := time.Now()
	for i := 0; i < 10000; i++ { // 插入 10 万条数据
		x := rand.Intn(1000)
		y := rand.Intn(1000)
		quad_tree.InsertEle(root, *quad_tree.NewElement(x, y, ""))
	}
	t := time.Now()
	fmt.Println(t.Sub(start).String()) // 计算插入的速度 40.621264ms

	// 查询
	testPoint := quad_tree.NewElement(300, 21, "")
	quad_tree.QueryNodeByElement(root, testPoint)

	// 绘制节点
	quad_tree.SetBackgroundColor()
	quad_tree.PrintAllQuadTree(root)
	quad_tree.PrintNodeByQuadTree(root, testPoint)
}
