package quad_tree

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"strconv"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

var img = image.NewRGBA(image.Rect(0, 0, 1006, 1006))
var col = color.RGBA{255, 255, 0, 255} // Green

// HLine draws a horizontal line
func HLine(x1, y, x2 int) {
	for ; x1 <= x2; x1++ {
		img.Set(x1, y, col)
	}
}

// VLine draws a veritcal line
func VLine(x, y1, y2 int) {
	for ; y1 <= y2; y1++ {
		img.Set(x, y1, col)
	}
}

// Rect draws a rectangle utilizing HLine() and VLine()
func Rect(x1, y1, x2, y2 int) {
	if x1 > x2 {
		x1, x2 = x2, x1
	}

	if y1 > y2 {
		y1, y2 = y2, y1
	}

	HLine(x1, y1, x2)
	HLine(x1, y2, x2)
	VLine(x1, y1, y2)
	VLine(x2, y1, y2)
}

func SetBackgroundColor() {
	draw.Draw(img, img.Bounds(), &image.Uniform{color.RGBA{50, 50, 50, 255}}, image.ZP, draw.Src)
}

// PrintQuadTree 打印整颗四叉树
func PrintAllQuadTree(node *QuadTreeNode) {
	travels(node)
	genImage()
}

// 找到元素的所在的 Node
func PrintNodeByQuadTree(node *QuadTreeNode, ele *ElePoint) {
	travelsByEle(node, ele)
	genImage()
}

func getRect(node *QuadTreeNode) (int, int, int, int) {
	return node.region.left, node.region.bottom, node.region.right, node.region.up
}

// 生成图片
func genImage() {
	f, err := os.Create("draw.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, img)
}

// 遍历
func travels(node *QuadTreeNode) {
	col = color.RGBA{255, 0, 0, 255} // Red
	Rect(getRect(node))
	col = color.RGBA{0, 255, 0, 255} // Green
	for i := 0; i < node.eleNum; i++ {
		if node.eleList[i] != nil {
			img.Set(int(node.eleList[i].x), int(node.eleList[i].y), col)
		}
	}

	if node.isLeaf {
		return
	}

	travels(node.LB)
	travels(node.RB)
	travels(node.LU)
	travels(node.RU)
}

// 遍历
func travelsByEle(node *QuadTreeNode, ele *ElePoint) {
	col = color.RGBA{0, 0, 255, 255} // Blue
	Rect(getRect(node))

	addLabel(img,
		(node.region.left+node.region.right)/2,
		(node.region.up+node.region.bottom)/2,
		strconv.Itoa(node.depth)+"-"+
			fmt.Sprintf("L:%d,B:%d,R:%d,U:%d",
				node.region.left,
				node.region.bottom,
				node.region.right,
				node.region.up))

	if node.isLeaf {
		for i := 0; i < node.eleNum; i++ {
			if node.eleList[i] != nil {
				col = color.RGBA{0, 255, 0, 255} // Green
				img.Set(int(node.eleList[i].x), int(node.eleList[i].y), col)
			}
		}
		return
	}

	midVertical := (node.region.up + node.region.bottom) / 2
	midHorizontal := (node.region.left + node.region.right) / 2

	if ele.y > midVertical {
		if ele.x > midHorizontal {
			travelsByEle(node.RU, ele)
		} else {
			travelsByEle(node.LU, ele)
		}
	} else {
		if ele.x > midHorizontal {
			travelsByEle(node.RB, ele)
		} else {
			travelsByEle(node.LB, ele)
		}
	}
}

// draw text
func addLabel(img *image.RGBA, x, y int, label string) {
	col := color.RGBA{255, 255, 255, 255}
	point := fixed.Point26_6{fixed.Int26_6(x * 64), fixed.Int26_6(y * 64)}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(label)
}
