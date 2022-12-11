package day11

import (
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

	round := 1
	totalRounds := 20
	for round <= totalRounds {
		for index := range monkeyList {
			currMonkey := &monkeyList[index]
			if len(currMonkey.items) == 0 {
				continue
			}
			for _, item := range currMonkey.items {
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
				new = new / 3
				thrownTo := 0
				if new%currMonkey.test.operand == 0 {
					thrownTo = currMonkey.test.trueMonkeyIndex
				} else {
					thrownTo = currMonkey.test.falseMonkeyIndex
				}
				monkeyList[thrownTo].items = append(monkeyList[thrownTo].items, new)
				currMonkey.items = currMonkey.items[1:]
			}
		}
		round++
	}
	sort.Slice(monkeyList, func(i, j int) bool {
		return monkeyList[i].inspections > monkeyList[j].inspections
	})
	return strconv.Itoa(monkeyList[0].inspections * monkeyList[1].inspections)
}

func Part2(input *string) string {
	lines := strings.Split(*input, "\n")
	monkeyList := make([]Monkey, 0)
	divisor := 1
	for i := 0; i < len(lines); i = i + 7 {
		monkey := Monkey{}
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
		divisor *= test.operand
		trueCondition := strings.Split(lines[i+4], "")
		test.trueMonkeyIndex, _ = strconv.Atoi(trueCondition[len(trueCondition)-1])
		falseCondition := strings.Split(lines[i+5], "")
		test.falseMonkeyIndex, _ = strconv.Atoi(falseCondition[len(falseCondition)-1])
		monkey.test = test
		monkey.inspections = 0
		monkeyList = append(monkeyList, monkey)
	}

	round := 1
	totalRounds := 10000
	for round <= totalRounds {
		for index := range monkeyList {
			currMonkey := &monkeyList[index]
			if len(currMonkey.items) == 0 {
				continue
			}

			for _, item := range currMonkey.items {
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
				new = new % divisor
				thrownTo := 0
				if new%currMonkey.test.operand == 0 {
					thrownTo = currMonkey.test.trueMonkeyIndex
				} else {
					thrownTo = currMonkey.test.falseMonkeyIndex
				}
				monkeyList[thrownTo].items = append(monkeyList[thrownTo].items, new)
				currMonkey.items = currMonkey.items[1:]
			}
		}
		round++
	}
	sort.Slice(monkeyList, func(i, j int) bool {
		return monkeyList[i].inspections > monkeyList[j].inspections
	})
	return strconv.Itoa(monkeyList[0].inspections * monkeyList[1].inspections)
}
