package main

import (
	"fmt"
	"strings"
	"strconv"
	"bufio"
	"os"
)

func main() {
	alice := strings.Split(strings.TrimSpace(readline()), " ")
	bob := strings.Split(strings.TrimSpace(readline()), " ")
	aliceTotalScore := 0
	bobTotalScore := 0
	for i := 0; i < len(alice); i++ {
		aliceScore, _ := strconv.Atoi(alice[i])
		bobScore, _ := strconv.Atoi(bob[i])
		aliceTotalScore += compare(aliceScore, bobScore)
		bobTotalScore += compare(bobScore, aliceScore)
	}
	fmt.Printf("%d %d", aliceTotalScore, bobTotalScore)
}

func compare(a, b int) int {
	if a > b {
		return 1
	}
	return 0
}

var reader *bufio.Reader

func readline() string {
	if reader == nil {
		reader = bufio.NewReader(os.Stdin)
	}
	line, _ := reader.ReadString('\n')
	return line
}
