package main

import (
	"bufio"
	"os"

	"github.com/defts/advent-of-code/utils"
)

var lines []string

func loadFile(f string) {
	file, _ := os.Open(f)
	buff := bufio.NewReader(file)
	lines = utils.ReaderToLines(buff)
}
