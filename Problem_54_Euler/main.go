package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

// https://projecteuler.net/problem=54

var (
	order = []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}
)

func index(card string) int {
	for i, c := range order {
		if c == card {
			return i
		}
	}

	panic("Not found")
}

func sortCards(cards []string) {
	sort.Slice(cards, func(i, j int) bool {
		return index(cards[i]) < index(cards[j])
	})
}

func formatHands(hands []string) map[string][]string {
	cards := map[string][]string{}
	for _, hand := range hands {
		suits := strings.Split(hand, "")
		num := suits[0]
		suit := suits[1]
		if _, ok := cards[suit]; !ok {
			cards[suit] = []string{}
		}
		cards[suit] = append(cards[suit], num)
	}

	for _, c := range cards {
		sortCards(c)
	}

	return cards
}

func fullHouse(hands []string) int {
	answer := nOfAKind(hands, 3, 1)
	verifier := nOfAKind(hands, 2, 1)
	if verifier*answer == 0 {
		return 0
	}

	return answer
}

func getCardCounts(cards map[string][]string) map[string]int {
	cardCounts := map[string]int{}
	for _, card := range cards {
		for _, suit := range card {
			cardCounts[suit]++
		}
	}

	return cardCounts
}

func nOfAKind(hands []string, n int, occurrences int) int {
	cards := formatHands(hands)
	cardCounts := getCardCounts(cards)
	value := float64(math.MinInt64)

	for k, c := range cardCounts {
		if c == n {
			occurrences -= 1
			value = math.Max(float64(index(k)), value)
		}
	}

	if occurrences == 0 {
		return int(value)
	}

	return 0
}

func fourOfAKind(hands []string) int {
	return nOfAKind(hands, 4, 1)
}

func threeOfAKind(hands []string) int {
	return nOfAKind(hands, 3, 1)
}

func twoPairs(hands []string) int {
	return nOfAKind(hands, 2, 2)
}

func onePair(hands []string) int {
	return nOfAKind(hands, 2, 1)
}

func highestCard(hands []string) int {
	cards := formatHands(hands)
	start := index("2")
	cardCounts := getCardCounts(cards)
	for _, card := range cards {
		for _, suit := range card {
			if cardCounts[suit] > 1 {
				continue
			}
			if index(suit) > start {
				start = index(suit)
			}
		}
	}

	return start
}

func straight(hands []string) int {
	cards := formatHands(hands)
	cardCounts := getCardCounts(cards)
	hasTwoAndAce := (cardCounts["A"] * cardCounts["2"]) != 0
	prevCard := ""
	for _, card := range order {
		if _, ok := cardCounts[card]; ok {
			if prevCard == "" {
				prevCard = card
				continue
			}
			if card == "A" && hasTwoAndAce {
				return 1
			}
			if index(card)-index(prevCard) != 1 {
				return 0
			}
			prevCard = card
		}
	}

	return highestCard(hands)
}

func flush(hands []string) int {
	cards := formatHands(hands)

	if len(cards) != 1 {
		return 0
	}

	return highestCard(hands)
}

func straightFlush(hands []string) int {
	cards := formatHands(hands)
	if len(cards) != 1 {
		return 0
	}

	for _, card := range cards {
		firstIndex := index(card[0])
		for i := 1; i < len(card); i++ {
			if index(card[i])-firstIndex != 1 {
				return 0
			}
			firstIndex = index(card[i])
		}
	}

	return highestCard(hands)
}

func isRoyalFlush(hands []string) int {
	cards := formatHands(hands)
	firstIndex := index("T")
	if len(cards) != 1 {
		return 0
	}

	for _, card := range cards {
		for _, c := range card {
			if index(c) != firstIndex {
				return 0
			}
			firstIndex += 1
		}
	}

	return 1
}

func winningHand(p1 []string, p2 []string) bool {
	p1Stats := []int{isRoyalFlush(p1), straightFlush(p1), fourOfAKind(p1), fullHouse(p1), flush(p1), straight(p1), threeOfAKind(p1), twoPairs(p1), onePair(p1), highestCard(p1)}
	p2Stats := []int{isRoyalFlush(p2), straightFlush(p2), fourOfAKind(p2), fullHouse(p2), flush(p2), straight(p2), threeOfAKind(p2), twoPairs(p2), onePair(p2), highestCard(p2)}

	fmt.Println(p1Stats)
	fmt.Println(p2Stats)

	for i := 0; i <= len(p1Stats); i++ {
		if p1Stats[i] == 0 && p2Stats[i] == 0 || p1Stats[i] == p2Stats[i] {
			continue
		}
		return p1Stats[i] > p2Stats[i]
	}

	return false
}

func main() {
	content, err := os.ReadFile("poker.txt")
	if err != nil {
		panic(err)
	}

	count := 0
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		fmt.Println(line)
		options := strings.Split(line, " ")
		p1 := options[:5]
		p2 := options[5:]

		if winningHand(p1, p2) {
			fmt.Println("p1")
			count += 1
		} else {
			fmt.Println("p2")
		}
	}

	fmt.Println("--------------")
	fmt.Println(count)
}
