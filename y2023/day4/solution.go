package day4

import (
	"strconv"
	"strings"
)

type Card struct {
	mynums   []int
	winnings []int
}

func Part1(input *string) string {
	cards := strings.Split(*input, "\n")
	sum := 0
	for _, sCard := range cards {
		card := parseCard(sCard)
		sum += card.score()
	}
	return strconv.Itoa(sum)
}

func Part2(input *string) string {
	cards := strings.Split(*input, "\n")
	var cardSlice []Card
	cardCountMap := make(map[int]int)
	for index, sCard := range cards {
		card := parseCard(sCard)
		cardSlice = append(cardSlice, card)
		cardCountMap[index] = 1
	}
	for index, card := range cardSlice {
		count := cardCountMap[index]
		for count > 0 {
			matches := card.matches()
			addCardCounter := 1
			for addCardCounter <= matches {
				cardCountMap[index+addCardCounter]++
				addCardCounter++
			}
			count--
		}
	}
	totalCards := 0
	for _, value := range cardCountMap {
		totalCards += value
	}
	return strconv.Itoa(totalCards)
}

func parseCard(card string) Card {
	myCard := &Card{}
	nums := strings.Trim(strings.Split(card, ":")[1], " ")
	nums_split := strings.Split(nums, "|")
	smyNums := strings.Split(strings.Trim(nums_split[0], " "), " ")
	for _, num := range smyNums {
		if num == "" {
			continue
		}
		i, _ := strconv.Atoi(strings.Trim(num, " "))
		myCard.mynums = append(myCard.mynums, i)
	}
	smyWinnings := strings.Split(strings.Trim(nums_split[1], " "), " ")
	for _, num := range smyWinnings {
		if num == "" {
			continue
		}
		i, _ := strconv.Atoi(strings.Trim(num, " "))
		myCard.winnings = append(myCard.winnings, i)
	}
	return *myCard
}

func (c *Card) score() int {
	score := 0
	for _, num := range c.mynums {
		for _, w := range c.winnings {
			if num == w {
				if score == 0 {
					score = 1
				} else {
					score *= 2
				}
			}
		}
	}
	return score
}

func (c *Card) matches() int {
	matches := 0
	for _, num := range c.mynums {
		for _, w := range c.winnings {
			if num == w {
				matches++
			}
		}
	}
	return matches
}
