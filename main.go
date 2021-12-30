package main

import (
	"fmt"
	"math/rand"
	"myquadtree/quad_tree"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	region := quad_tree.NewRegion(0, 360, 0, 360)
	root := quad_tree.NewNode(1, *region)

	start := time.Now()
	for i := 0; i < 100000; i++ { // 插入 10 万条数据
		x := rand.Intn(360)
		y := rand.Intn(360)
		quad_tree.InsertEle(root, *quad_tree.NewElement(x, y, ""))
	}
	t := time.Now()
	fmt.Println(t.Sub(start).String()) // 计算插入的速度 40.621264ms

	// 查询
	testPoint := quad_tree.NewElement(124, 145, "")
	quad_tree.QueryNodeByElement(root, testPoint)

	// 绘制节点
	quad_tree.SetBackgroundColor()
	quad_tree.PrintAllQuadTree(root)
	quad_tree.PrintNodeByQuadTree(root, testPoint)
}
