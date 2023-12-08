package day7

import (
	"sort"
	"strconv"
	"strings"
)

type HandType int

const (
	HIGH_CARD HandType = iota
	ONE_PAIR
	TWO_PAIR
	THREE_OF_KIND
	FULL_HOUSE
	FOUR_OF_KIND
	FIVE_OF_KIND
)

func CardScore(card rune, part int) int {
	if part == 2 && card == 'J' {
		return 1
	}
	cardMap := map[rune]int{
		'A': 14,
		'K': 13,
		'Q': 12,
		'J': 11,
		'T': 10,
		'9': 9,
		'8': 8,
		'7': 7,
		'6': 6,
		'5': 5,
		'4': 4,
		'3': 3,
		'2': 2,
	}
	return cardMap[card]
}

type Hand struct {
	raw      string
	cards    [5]int
	handType HandType
	bid      int
	jCount   int
}

func (h *Hand) parseHand(cardData [2]string) {
	cards := cardData[0]
	bid := cardData[1]
	h.raw = cards
	for index, card := range cards {
		h.cards[index] = CardScore(card, 1)
	}
	h.bid, _ = strconv.Atoi(bid)
	h.determineHandType()
}

func (h *Hand) parseHand2(cardData [2]string) {
	cards := cardData[0]
	bid := cardData[1]
	h.raw = cards
	for index, card := range cards {
		h.cards[index] = CardScore(card, 2)
	}
	h.bid, _ = strconv.Atoi(bid)
	h.determineHandType()
	h.determineJIncrease()
}

func (h *Hand) determineHandType() {
	cards := make(map[int]int)
	for _, card := range h.cards {
		if _, ok := cards[card]; !ok {
			cards[card] = 1
		} else {
			cards[card]++
		}
	}

	h.jCount = cards[1]

	if len(cards) == 1 {
		h.handType = FIVE_OF_KIND
		return
	}
	if len(cards) == 5 {
		h.handType = HIGH_CARD
		return
	}
	if len(cards) == 4 {
		h.handType = ONE_PAIR
		return
	}
	if len(cards) == 2 {
		for _, value := range cards {
			if value == 4 || value == 1 {
				h.handType = FOUR_OF_KIND
				return
			}
			if value == 2 {
				h.handType = FULL_HOUSE
				return
			}
		}
	}
	if len(cards) == 3 {
		for _, value := range cards {
			if value == 2 {
				h.handType = TWO_PAIR
				return
			}
			if value == 3 {
				h.handType = THREE_OF_KIND
				return
			}
		}
	}
}

func (h *Hand) determineJIncrease() {
	switch h.handType {
	case FIVE_OF_KIND:
		h.handType = FIVE_OF_KIND
		return
	case FOUR_OF_KIND:
		if h.jCount != 0 {
			h.handType = FIVE_OF_KIND
			return
		}
	case FULL_HOUSE:
		if h.jCount != 0 {
			h.handType = FIVE_OF_KIND
			return
		}
	case THREE_OF_KIND:
		if h.jCount != 0 {
			h.handType = FOUR_OF_KIND
			return
		}
	case TWO_PAIR:
		if h.jCount == 2 {
			h.handType = FOUR_OF_KIND
			return
		}
		if h.jCount == 1 {
			h.handType = FULL_HOUSE
			return
		}
	case ONE_PAIR:
		if h.jCount != 0 {
			h.handType = THREE_OF_KIND
			return
		}
	case HIGH_CARD:
		if h.jCount != 0 {
			h.handType = ONE_PAIR
		}
	}
}

func Part1(input *string) string {
	cardLine := strings.Split(*input, "\n")
	var hands []Hand
	for _, cardStr := range cardLine {
		cardData := strings.Split(cardStr, " ")
		hand := &Hand{}
		hand.parseHand([2]string(cardData))
		hands = append(hands, *hand)
	}
	sort.Slice(hands, func(i, j int) bool {
		hand1 := hands[i]
		hand2 := hands[j]
		if hand1.handType == hand2.handType {
			cards1 := hand1.cards
			cards2 := hand2.cards
			for index := range cards1 {
				if cards1[index] == cards2[index] {
					continue
				}
				return cards1[index] < cards2[index]
			}
		}
		return hand1.handType < hand2.handType
	})
	sum := 0
	for index, hand := range hands {
		score := (index + 1) * hand.bid
		sum += score
	}
	return strconv.Itoa(sum)
}

func Part2(input *string) string {
	cardLine := strings.Split(*input, "\n")
	var hands []Hand
	for _, cardStr := range cardLine {
		cardData := strings.Split(cardStr, " ")
		hand := &Hand{}
		hand.parseHand2([2]string(cardData))
		hands = append(hands, *hand)
	}
	sort.Slice(hands, func(i, j int) bool {
		hand1 := hands[i]
		hand2 := hands[j]
		if hand1.handType == hand2.handType {
			cards1 := hand1.cards
			cards2 := hand2.cards
			for index := range cards1 {
				if cards1[index] == cards2[index] {
					continue
				}
				return cards1[index] < cards2[index]
			}
		}
		return hand1.handType < hand2.handType
	})
	sum := 0
	for index, hand := range hands {
		score := (index + 1) * hand.bid
		sum += score
	}
	return strconv.Itoa(sum)
}
