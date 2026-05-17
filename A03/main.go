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

func linearSearch(p, q []int, value int) bool {
	for _, v1 := range p {
		for _, v2 := range q {
			if value == v1 + v2 {
				return true
			}
		}
	}
	return false
}

func main() {
	defer wr.Flush()

	sc.Scan()
	parts := strings.Fields(sc.Text())
	N, _ := strconv.Atoi(parts[0])
	K, _ := strconv.Atoi(parts[1])

	sc.Scan()
	numStr1 := strings.Fields(sc.Text())
	P := make([]int, N)
	for i, s := range numStr1 {
		P[i], _ = strconv.Atoi(s)
	}

	sc.Scan()
	numStr2 := strings.Fields(sc.Text())
	Q := make([]int, N)
	for i, s := range numStr2 {
		Q[i], _ = strconv.Atoi(s)
	}
	
	if linearSearch(P, Q, K) {
		fmt.Fprintln(wr, "Yes")
	} else {
		fmt.Fprintln(wr, "No")
	}
}