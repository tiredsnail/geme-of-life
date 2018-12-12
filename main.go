package main

import "fmt"

var board [][]int

func main() {
	board := createPaint()
	paintGosperGliderGun(board)

	for i:=0; i<len(board); i++ {
		for ii:=0; ii<len(board[i]); ii++ {
			if board[i][ii] == 1 {
				fmt.Print("#")
			}
		}
		fmt.Print("\n")

	}
}


//创建画布
func createPaint() (board [][]int) {
	n := 50
	board = make([][]int,n)
	for i := 0; i < n; i++ {
		board[i]= make([]int, n)
	}
	return board
}

// paint reload
func reloadPaint() {
	for y:=0; y<len(board); y++ {
		for x:=0; x<len(board[y]); x++ {
			reverse(y, x)
		}
	}
}

// reverse 生死判断
func reverse(y int, x int) {

}


// paint Gosper Glider Gun 高帕斯滑翔机枪
func paintGosperGliderGun(board [][]int) {
	board[0][1] = 1
}


// paint 繁殖者


//





//slic := [][]int{
//	0:{1,1,1,1,1,1,1,1,1,1},
//	1:{1,1,1,1,1,1,1,1,1,1},
//	2:{1,1,1,1,1,1,1,1,1,1},
//	3:{1,1,1,1,1,1,1,1,1,1},
//	4:{1,1,1,1,1,1,1,1,1,1},
//	5:{1,1,1,1,1,1,1,1,1,1},
//	6:{1,1,1,1,1,1,1,1,1,1},
//	7:{1,1,1,1,1,1,1,1,1,1},
//	8:{1,1,1,1,1,1,1,1,1,1},
//	9:{1,1,1,1,1,1,1,1,1,1},
//}
//fmt.Println("slic", slic)
//
//for i:=0; i<len(slic); i++ {
//	for ii:=0; ii<len(slic[i]); ii++ {
//		if slic[i][ii] == 0 {
//			fmt.Print("#")
//		}
//	}
//	fmt.Print("\n")
//
//}