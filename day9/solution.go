package day9

import (
	"fmt"
	"strconv"
	"strings"
)

type Knot struct {
	next *Knot
	prev *Knot
	x    int
	y    int
}

type Rope struct {
	head *Knot
	tail *Knot
}

func Part1(input *string) string {
	moves := strings.Split(*input, "\n")
	size := 2
	rope := &Rope{}
	moveMap := make(map[string]int, 0)
	moveMap["0,0"] = 1
	for i := 0; i < size; i++ {
		rope.InsertKnot()
	}
	for _, move := range moves {
		moveParser := strings.Split(move, " ")
		dir := moveParser[0]
		steps, _ := strconv.Atoi(moveParser[1])
		fmt.Println(dir, steps)
		for i := 0; i < steps; i++ {
			rope.head.moveNode(dir)
			moveMap[fmt.Sprintf("%d,%d", rope.tail.x, rope.tail.y)] = 1
			//fmt.Printf("\n HEAD: %d,%d, TAIL %d,%d \n", rope.head.x, rope.head.y, rope.tail.x, rope.tail.y)
		}
	}
	return fmt.Sprintf("%d", len(moveMap))
}

func Part2(input *string) string {
	moves := strings.Split(*input, "\n")
	size := 10
	rope := &Rope{}
	moveMap := make(map[string]int, 0)
	moveMap["0,0"] = 1
	for i := 0; i < size; i++ {
		rope.InsertKnot()
	}
	for _, move := range moves {
		moveParser := strings.Split(move, " ")
		dir := moveParser[0]
		steps, _ := strconv.Atoi(moveParser[1])
		//fmt.Println(dir, steps)
		for i := 0; i < steps; i++ {
			rope.head.moveNode(dir)
			moveMap[fmt.Sprintf("%d,%d", rope.tail.x, rope.tail.y)] = 1
			//fmt.Printf("\n HEAD: %d,%d, TAIL %d,%d \n", rope.head.x, rope.head.y, rope.tail.x, rope.tail.y)
		}
	}
	return fmt.Sprintf("%d", len(moveMap))
}

func (rope *Rope) InsertKnot() {
	knot := &Knot{x: 0, y: 0, next: nil, prev: nil}
	if rope.head == nil {
		rope.head = knot
		rope.tail = knot
	} else {
		head := rope.head
		for head.next != nil {
			head = head.next
		}
		knot.prev = head
		head.next = knot
		rope.tail = knot
	}
}

func (k *Knot) moveNode(dir string) {
	switch dir {
	case "R":
		k.x++
	case "L":
		k.x--
	case "U":
		k.y++
	case "D":
		k.y--
	}

	if k.next != nil {
		//fmt.Println(dir, *k, *k.next)
		if k.x-k.next.x == 2 {
			k.next.y = k.y
			k.next.moveNode("R")

		}
		if k.x-k.next.x == -2 {
			k.next.y = k.y
			k.next.moveNode("L")
		}
		if k.y-k.next.y == +2 {
			k.next.x = k.x
			k.next.moveNode("U")
		}
		if k.y-k.next.y == -2 {
			k.next.x = k.x
			k.next.moveNode("D")
		}
		return
	}
}
