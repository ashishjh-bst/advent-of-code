package day8

import (
	"strconv"
	"strings"
)

type Node struct {
	left  string
	right string
}

type NodeMap map[string]Node

func (nm *NodeMap) TraverseInstructions(instruction string, startNode string, count int) int {
	for _, i := range instruction {
		nextNode := (*nm)[startNode]
		if i == 'L' {
			startNode = nextNode.left
		} else {
			startNode = nextNode.right
		}
		count++
	}
	if startNode == "ZZZ" {
		return count
	}
	return nm.TraverseInstructions(instruction, startNode, count)
}

func (nm *NodeMap) TraverseInstructions2(instruction string, startNode string, count int) int {
	for _, i := range instruction {
		nextNode := (*nm)[startNode]
		if i == 'L' {
			startNode = nextNode.left
		} else {
			startNode = nextNode.right
		}
		count++
	}
	if startNode[2:] == "Z" {
		return count
	}
	return nm.TraverseInstructions2(instruction, startNode, count)
}

func parseNode(data string) Node {
	nodes := strings.Split(data, ",")
	node := &Node{}
	node.left = strings.Trim(nodes[0], " ()")
	node.right = strings.Trim(nodes[1], " ()")
	return *node
}

func Part1(input *string) string {
	nodeMap := make(NodeMap)
	data := strings.Split(*input, "\n\n")
	instructions := data[0]
	for _, line := range strings.Split(data[1], "\n") {
		nodeData := strings.Split(line, "=")
		nodeMap[strings.Trim(nodeData[0], " ")] = parseNode(nodeData[1])
	}
	count := nodeMap.TraverseInstructions(instructions, "AAA", 0)
	return strconv.Itoa(count)
}

func Part2(input *string) string {
	nodeMap := make(NodeMap)
	data := strings.Split(*input, "\n\n")
	instructions := data[0]
	var startNodes []string
	for _, line := range strings.Split(data[1], "\n") {
		nodeData := strings.Split(line, "=")
		node := strings.Trim(nodeData[0], " ")
		nodeMap[node] = parseNode(nodeData[1])
		if node[2:] == "A" {
			startNodes = append(startNodes, node)
		}
	}
	var nodeCounters []int
	for _, node := range startNodes {
		count := nodeMap.TraverseInstructions2(instructions, node, 0)
		nodeCounters = append(nodeCounters, count)
	}
	var lcm int
	if len(nodeCounters) == 2 {
		lcm = LCM(nodeCounters[0], nodeCounters[1])
	} else {
		lcm = LCM(nodeCounters[0], nodeCounters[1], nodeCounters[2:]...)
	}

	return strconv.Itoa(lcm)
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
