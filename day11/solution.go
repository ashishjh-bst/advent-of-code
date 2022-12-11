package day11

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	op          Operation
	test        Test
	items       []int
	inspections int
}

type Operation struct {
	operator string
	operand  string
}

type Test struct {
	operator         string
	operand          int
	trueMonkeyIndex  int
	falseMonkeyIndex int
}

func Part1(input *string) string {
	lines := strings.Split(*input, "\n")
	monkeyList := make([]Monkey, 0)
	for i := 0; i < len(lines); i = i + 7 {
		monkey := Monkey{}
		//fetch Items
		items := strings.Split(strings.Split(lines[i+1], ": ")[1], ", ")
		for _, item := range items {
			it, _ := strconv.Atoi(item)
			monkey.items = append(monkey.items, it)
		}
		op := Operation{}
		operationString := strings.Split(lines[i+2], ": ")
		ops := strings.Split(operationString[1], " ")
		op.operator = ops[3]
		op.operand = ops[4]
		monkey.op = op

		test := Test{}
		conditionString := strings.Split(lines[i+3], " ")
		test.operator = conditionString[3]
		test.operand, _ = strconv.Atoi(conditionString[5])
		trueCondition := strings.Split(lines[i+4], "")
		test.trueMonkeyIndex, _ = strconv.Atoi(trueCondition[len(trueCondition)-1])
		falseCondition := strings.Split(lines[i+5], "")
		test.falseMonkeyIndex, _ = strconv.Atoi(falseCondition[len(falseCondition)-1])
		monkey.test = test
		monkey.inspections = 0

		monkeyList = append(monkeyList, monkey)
	}

	lastMonkey := len(monkeyList)
	fmt.Printf("\nTotal Monkeys: %d and lastMonkey is %d", len(monkeyList), lastMonkey)
	round := 1
	totalRounds := 20
	for round <= totalRounds {
		for index, _ := range monkeyList {
			currMonkey := &monkeyList[index]
			if len(currMonkey.items) == 0 {
				continue
			}
			fmt.Printf("\n Monkey %d:", index)
			for _, item := range currMonkey.items {
				fmt.Printf("\n\t\tMonkey inspects an item with a worry level of %d", item)
				currMonkey.inspections++
				operand := 0
				new := 0
				old := item
				if currMonkey.op.operand == "old" {
					operand = old
				} else {
					operand, _ = strconv.Atoi(currMonkey.op.operand)
				}
				switch currMonkey.op.operator {
				case "*":
					new = old * operand
				case "+":
					new = old + operand
				}
				fmt.Printf("\n\t\tWorry level is %s by %d to %d", currMonkey.op.operator, operand, new)
				new = new / 3
				fmt.Printf("\n\t\tMonkey gets bored with item, Worry level is divided by 3 to %d", new)
				thrownTo := 0
				if new%currMonkey.test.operand == 0 {
					fmt.Printf("\n\t\tCurrent worry level is divisable by %d", currMonkey.test.operand)
					thrownTo = currMonkey.test.trueMonkeyIndex
				} else {
					thrownTo = currMonkey.test.falseMonkeyIndex
					fmt.Printf("\n\t\tCurrent worry level is not divisable by %d", currMonkey.test.operand)
				}
				fmt.Printf("\n\t\t\tItem with worry level %d is thrown to monkey %d", new, thrownTo)
				monkeyList[thrownTo].items = append(monkeyList[thrownTo].items, new)
				currMonkey.items = currMonkey.items[1:]
			}
		}
		fmt.Printf("\n End of round %d", round)
		for index, monkey := range monkeyList {
			fmt.Printf("\n\t Monkey%d: %#v", index, monkey.items)
		}
		round++
	}
	sort.Slice(monkeyList, func(i, j int) bool {
		return monkeyList[i].inspections > monkeyList[j].inspections
	})
	return strconv.Itoa(monkeyList[0].inspections * monkeyList[1].inspections)
}

func Part2(input *string) string {
	return ""
}
