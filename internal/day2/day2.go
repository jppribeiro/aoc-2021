package day2

import (
	"aoc-2021/internal/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type pos struct {
	h int
	d int
}

type posWithAim struct {
	h int
	d int
	a int
}

func Solution() {
	lines := utils.ReadInput("internal/day2/input2")

	pos := pos{0,0}



	posAim := posWithAim{0,0,0}

	for _, m := range lines {
		move(&pos, m)
		moveWithAim(&posAim, m)
	}

	fmt.Printf("First movement type results in %d\n", pos.h * pos.d)
	fmt.Printf("Second movement type results in %d\n", posAim.h * posAim.d)
}

func move(current *pos, mov string) {
	m := strings.Split(mov, " ")

	dir := m[0]
	dist, err := strconv.Atoi(m[1])

	if err != nil {
		fmt.Errorf("%s", err)
		os.Exit(1)
	}

	switch dir {
	case "forward":
		current.h = current.h + dist
	case "down":
		current.d = current.d + dist
	case "up":
		current.d = current.d - dist
	default:
		fmt.Errorf("wrong direction %s", dir)
	}
}

func moveWithAim(cur *posWithAim, move string) {
	m := strings.Split(move, " ")

	dir := m[0]
	dist, err := strconv.Atoi(m[1])

	if err != nil {
		fmt.Errorf("%s", err)
		os.Exit(1)
	}

	switch dir {
	case "forward":
		cur.h = cur.h + dist
		cur.d = cur.d + cur.a * dist
	case "down":
		cur.a = cur.a + dist
	case "up":
		cur.a = cur.a - dist
	default:
		fmt.Errorf("wrong direction %s", dir)
	}
}