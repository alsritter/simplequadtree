package quad_tree

import (
	"fmt"
)

const MAX_ELE_NUM = 10

// Region 表示节点保存元素的范围
type Region struct {
	up     int
	bottom int
	left   int
	right  int
}

// ElePoint 保存真实的数据
type ElePoint struct {
	x    int
	y    int
	data string
}

// QuadTreeNode 一个四叉树节点，里面包含若干个元素
type QuadTreeNode struct {
	depth   int                    // 节点深度
	isLeaf  bool                   // 是否是叶子节点
	region  *Region                // 区域范围
	LU      *QuadTreeNode          // 左上子结点指针
	LB      *QuadTreeNode          // 左下子结点指针
	RU      *QuadTreeNode          // 右上子结点指针
	RB      *QuadTreeNode          // 右下子结点指针
	eleNum  int                    // 位置点数
	eleList [MAX_ELE_NUM]*ElePoint // 位置点列表
}

func NewNode(depth int, region *Region) *QuadTreeNode {
	node := &QuadTreeNode{}
	node.depth = depth
	node.isLeaf = true
	node.eleNum = 0
	node.region = region
	return node
}

func NewRegion(bottom, up, left, right int) *Region {
	return &Region{
		up:     up,
		bottom: bottom,
		left:   left,
		right:  right,
	}
}

func NewElement(x, y int, data string) *ElePoint {
	return &ElePoint{x: x, y: y, data: data}
}

/**
 * 插入元素
 * 1.判断是否已分裂，已分裂的选择适合的子结点，插入；
 * 2.未分裂的查看是否过载，过载的分裂结点，重新插入；
 * 3.未过载的直接添加
 *
 * @param node
 * @param ele
 * todo 使用元素原地址，避免重新分配内存造成的效率浪费
 */
func InsertEle(node *QuadTreeNode, ele ElePoint) {
	if node.isLeaf {
		// 如果该节点包含的元素已经大于最大的元素数量则分裂节点
		if node.eleNum+1 > MAX_ELE_NUM {
			splitNode(node)
			InsertEle(node, ele)
		} else {
			node.eleList[node.eleNum] = &ele
			node.eleNum++
		}
		return
	}

	midVertical := (node.region.up + node.region.bottom) / 2
	midHorizontal := (node.region.left + node.region.right) / 2

	// 把元素插入到对应的节点
	if ele.y > midVertical {
		if ele.x > midHorizontal {
			InsertEle(node.RU, ele)
		} else {
			InsertEle(node.LU, ele)
		}
	} else {
		if ele.x > midHorizontal {
			InsertEle(node.RB, ele)
		} else {
			InsertEle(node.LB, ele)
		}
	}
}

func DeleteEle(node *QuadTreeNode, ele *ElePoint) {
	/**
	 * 1.遍历元素列表，删除对应元素
	 * 2.检查兄弟象限元素总数，不超过最大量时组合兄弟象限
	 */
}

func DeleteNode(node *QuadTreeNode) {
	/**
	 * 遍历四个子象限的点，添加到象限点列表
	 * 释放子象限的内存
	 */
}

// 找到元素的所在的 Node
func QueryNodeByElement(node *QuadTreeNode, ele *ElePoint) {
	fmt.Printf("%+v %+v %+v %+v \n\n", node.depth, node.region, node.isLeaf, node.eleNum)
	if node.isLeaf {
		fmt.Printf("当前 Node 附近点有 %d 个，分别是：\n", node.eleNum)
		for i := 0; i < node.eleNum; i++ {
			// fmt.Printf("(%f,%f) \n", node.eleList[i].x, node.eleList[i].y)
		}
		return
	}

	midVertical := (node.region.up + node.region.bottom) / 2
	midHorizontal := (node.region.left + node.region.right) / 2

	if ele.y > midVertical {
		if ele.x > midHorizontal {
			QueryNodeByElement(node.RU, ele)
		} else {
			QueryNodeByElement(node.LU, ele)
		}
	} else {
		if ele.x > midHorizontal {
			QueryNodeByElement(node.RB, ele)
		} else {
			QueryNodeByElement(node.LB, ele)
		}
	}
}

/**
 * 分裂结点
 * 1.通过父结点获取子结点的深度和范围
 * 2.生成四个结点，挂载到父结点下
 *
 * @param node
 */
func splitNode(node *QuadTreeNode) {
	midVertical := (node.region.up + node.region.bottom) / 2
	midHorizontal := (node.region.left + node.region.right) / 2

	fmt.Printf("(%d,%d) (%d,%d)-(%d,%d) \n\n", midHorizontal, midVertical, node.region.left, node.region.right, node.region.bottom, node.region.up)

	node.isLeaf = false
	node.RU = NewNode(node.depth+1, NewRegion(midVertical, node.region.up, midHorizontal, node.region.right))
	node.LU = NewNode(node.depth+1, NewRegion(midVertical, node.region.up, node.region.left, midHorizontal))
	node.RB = NewNode(node.depth+1, NewRegion(node.region.bottom, midVertical, midHorizontal, node.region.right))
	node.LB = NewNode(node.depth+1, NewRegion(node.region.bottom, midVertical, node.region.left, midHorizontal))

	// 遍历结点下的位置点，将其插入到子结点中
	for i := 0; i < node.eleNum; i++ {
		InsertEle(node, *node.eleList[i])
		node.eleList[i] = nil // 释放空间
		node.eleNum--
	}
}
