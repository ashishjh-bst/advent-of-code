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
	return ""
}
