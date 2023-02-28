package data_structures

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInt(in *bufio.Reader) int {
	nStr, _ := in.ReadString('\n')
	nStr = strings.ReplaceAll(nStr, "\r", "")
	nStr = strings.ReplaceAll(nStr, "\n", "")
	n, _ := strconv.Atoi(nStr)
	return n
}
func readLine(in *bufio.Reader) []string {
	line, _ := in.ReadString('\n')
	line = strings.ReplaceAll(line, "\r", "")
	line = strings.ReplaceAll(line, "\n", "")
	numbs := strings.Split(line, " ")
	return numbs
}

func readArray(in *bufio.Reader) []int {
	numbs := readLine(in)
	arr := make([]int, len(numbs))
	for i, n := range numbs {
		val, _ := strconv.Atoi(n)
		arr[i] = val
	}
	return arr
}

func example() {
	var matrix [][]int
	in := bufio.NewReader(os.Stdin)
	n := readInt(in)
	for i := 0; i < n; i++ {
		arr := readArray(in)
		matrix = append(matrix, arr)
	}
	fmt.Println(matrix)

}
