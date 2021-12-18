package main

// func getRow(rowIndex int) []int {
// 	rst := make([][]int, rowIndex+1)
// 	rst[0] = []int{1}
// 	for i := 1; i < rowIndex+1; i++ {
// 		row := make([]int, i+1)
// 		for j := 0; j < i+1; j++ {
// 			top := 0
// 			topleft := 0
// 			if j < i {
// 				top = rst[i-1][j]
// 			}
// 			if j-1 >= 0 {
// 				topleft = rst[i-1][j-1]
// 			}
// 			row[j] = top + topleft
// 		}
// 		rst[i] = row
// 	}
// 	return rst[rowIndex]
// }
// func main() {
// 	rst := getRow(3)
// 	fmt.Print(rst)
// }
