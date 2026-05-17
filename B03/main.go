package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
)

const (
	initBufSize = 1024 * 1024
	maxBufSize = 1024 * 1024 * 1024
)

func init() {
	buf := make([]byte, initBufSize)
	sc.Buffer(buf, maxBufSize)
}

func searchTriple(a []int, target int) bool {
	n := len(a)
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				if a[i]+a[j]+a[k] == target {
					return true
				}
			}
		}
	}
	return false
}

func main() {
	defer wr.Flush()

	sc.Scan()
	N, _ := strconv.Atoi(sc.Text())

	sc.Scan()
	numStr := strings.Fields(sc.Text())
	A := make([]int, N)
	for i, s := range numStr {
		A[i], _ = strconv.Atoi(s)
	}

	if searchTriple(A, 1000) {
		fmt.Fprintln(wr, "Yes")
	} else {
		fmt.Fprintln(wr, "No")
	}
}
