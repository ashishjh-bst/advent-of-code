package day7

import (
	"fmt"
	"strconv"
	"strings"
)

type Operation struct {
	input     []string
	output    string
	processed bool
}

type Register map[string]uint16

func Part1(input *string) string {
	r := &Register{}
	Ops := parseInput(input, r)
	processed := 0
	for len(Ops) > processed {
		for i := 0; i < len(Ops); i++ {
			op := Ops[i]
			if op.processed {
				continue
			}
			isSkipped := op.process(r)
			if !isSkipped {
				processed++
			}
		}
	}
	return fmt.Sprintf("%d", (*r)["a"])
}

func Part2(input *string) string {
	resultA := Part1(input)
	aVal, _ := strconv.ParseUint(resultA, 10, 16)
	r := &Register{}
	Ops := parseInput(input, r)
	r.setValue("b", uint16(aVal))
	processed := 0
	for len(Ops) > processed {
		for i := 0; i < len(Ops); i++ {
			op := Ops[i]
			fmt.Printf("\n input: %v, output %s, isProcessed %t", op.input, op.output, op.processed)
			if op.processed {
				continue
			}
			isSkipped := op.process(r)
			fmt.Printf(" isSkipped %t", isSkipped)
			if !isSkipped {
				processed++
				fmt.Printf("\n processed: %d of %d", processed, len(Ops))
			}
		}
	}
	return fmt.Sprintf("%d", (*r)["a"])
}

func parseInput(input *string, r *Register) []*Operation {
	operations := make([]*Operation, 0)
	lines := strings.Split(*input, "\n")
	for _, line := range lines {
		parts := strings.Split(line, " -> ")
		input := parts[0]
		output := parts[1]
		inputs := strings.Split(input, " ")
		if len(inputs) == 1 && IsNumber(inputs[0]) {
			val, _ := strconv.ParseUint(inputs[0], 10, 16)
			r.setValue(output, uint16(val))
			continue
		}
		operations = append(operations, &Operation{input: inputs, output: output})
	}
	return operations
}

func (o *Operation) process(r *Register) bool {
	var result uint16
	switch len(o.input) {
	case 1:
		input := o.input[0]
		if !IsNumber(input) && !r.hasKey(input) {
			return true
		}
		result = r.getValue(input)
	case 2:
		input := o.input[1]
		if !IsNumber(input) && !r.hasKey(input) {
			return true
		}
		value := r.getValue(input)
		result = ^value
	case 3:
		a := o.input[0]
		if !IsNumber(a) && !r.hasKey(a) {
			return true
		}
		b := o.input[2]
		if !IsNumber(b) && !r.hasKey(b) {
			return true
		}
		gate := o.input[1]
		aVal := r.getValue(a)
		bVal := r.getValue(b)
		switch gate {
		case "AND":
			result = aVal & bVal
		case "OR":
			result = aVal | bVal
		case "LSHIFT":
			result = aVal << bVal
		case "RSHIFT":
			result = aVal >> bVal
		}
	}
	r.setValue(o.output, result)
	o.processed = true
	return false
}

func (r *Register) setValue(a string, value uint16) {
	(*r)[a] = value
}

func (r *Register) hasKey(a string) bool {
	_, ok := (*r)[a]
	return ok
}

func (r *Register) getValue(a string) uint16 {
	val, err := strconv.ParseUint(a, 10, 16)
	if err != nil {
		return (*r)[a]
	}
	return uint16(val)
}

func IsNumber(a string) bool {
	_, err := strconv.ParseUint(a, 10, 16)
	return err == nil
}
