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

func main() {
	defer wr.Flush()

	sc.Scan()
	parts := strings.Fields(sc.Text())
	N, _ := strconv.Atoi(parts[0])
	K, _ := strconv.Atoi(parts[1])

	count := 0
	for r := 1; r <= N; r++ {
		for b := 1; b <= N; b++ {
			w := K - r - b
			if w >= 1 && w <= N {
				count++
			}
		}
	}

	fmt.Fprintln(wr, count)
}