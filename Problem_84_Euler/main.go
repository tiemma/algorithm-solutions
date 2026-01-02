package main

import (
	"fmt"
	"math"
	"sort"
)

// https://projecteuler.net/problem=84

// Abbreviations:
// CC: Community Chest
// CH: Chance

const (
	SIZE           = 40.0
	MAX_ROLLS      = 3 // Number of doubles in a row
	MAX_STATES     = SIZE * MAX_ROLLS
	MAX_ITERATIONS = 15_000
	NUM_SIDES_DICE = 4
	GO_TO_JAIL     = 30
	JAIL           = 10
)

var (
	// Special tiles
	CC        = []int{2, 17, 23}
	CH        = []int{7, 22, 36}
	RAILS     = []int{5, 15, 25, 35}
	UTILITIES = []int{12, 28}
)

// For markov chains, we mark every possible combination as a separate state
// Therefore a default no doubles is one state
// 1 double marks the next square
// 2 doubles marks the next square and the next square after that
// 3 doubles moves to JAIL so we can ignore that
// [[NO_DOUBLES, DOUBLE_1, DOUBLE_2], ...SAME PATTERN FOR EACH SQUARE INDEX]
// We index with (square * MAX_ROLLS) + doubles
func stateIndex(square, doubles int) int {
	return (square * MAX_ROLLS) + doubles
}

func hasArr(arr []int, x int) bool {
	for _, n := range arr {
		if n == x {
			return true
		}
	}

	return false
}

type transit struct {
	square      int
	probability float64
}

// Generate map of transition squares for each square
// We will use this to generate the Markov chain matrix
func generateTransitMap() map[int][]transit {
	transitMap := map[int][]transit{}

	// G2J -> JAIL, this has a probability of 1 as it will happen every time
	transitMap[30] = append(transitMap[30], transit{square: 10, probability: 1})

	for _, index := range CC {
		// Jump to
		// GO
		// JAIL
		transitMap[index] = []transit{
			{
				square:      0,
				probability: 1 / 16.0,
			},
			{
				square:      10,
				probability: 1 / 16.0,
			},
		}
	}

	for _, index := range CH {
		// Go back three spaces
		transitMap[index] = append(transitMap[index], transit{(index - 3) % 40, 1 / 16.0})

		// Jump to
		// GO
		// JAIL
		// C1
		// E3
		// H2
		// R1
		transitMap[index] = append(transitMap[index],
			transit{square: 0, probability: 1.0 / 16.0},
			transit{square: 5, probability: 1.0 / 16.0},
			transit{square: 10, probability: 1.0 / 16.0},
			transit{square: 11, probability: 1.0 / 16.0},
			transit{square: 24, probability: 1.0 / 16.0},
			transit{square: 39, probability: 1.0 / 16.0},
		)

		// Next R, we have two next rail cards
		nextR := nextTransition(RAILS, index)
		transitMap[index] = append(transitMap[index], transit{square: nextR, probability: 2 / 16.0})

		// Next U
		nextU := nextTransition(UTILITIES, index)
		transitMap[index] = append(transitMap[index], transit{square: nextU, probability: 1 / 16.0})
	}

	for _, arr := range transitMap {
		sort.Slice(arr, func(i, j int) bool { return arr[i].square < arr[j].square })
	}

	return transitMap
}

// Find the next square that has a transition for CC and CH cards
func nextTransition(states []int, breakPoint int) int {
	for _, index := range states {
		if index > breakPoint {
			return index
		}
	}

	return states[0]
}

type roll struct {
	sum         int
	isDouble    bool
	probability float64
}

// Generate probability of each outcome
func diceProbability() []roll {
	var probabilities []roll
	for i := 1; i <= NUM_SIDES_DICE; i++ {
		for j := 1; j <= NUM_SIDES_DICE; j++ {
			probabilities = append(probabilities, roll{isDouble: i == j, sum: i + j, probability: 1 / 36.0})
		}
	}

	return probabilities
}

// Generate Markov chain matrix of iterations with probability
// https://en.wikipedia.org/wiki/Markov_chain
// M[i][j] = P(S_{t+1} = j | S_t = i)

// We will consider each double roll as a separate state
// We can have no doubles, 1 double or 2 doubles ensuring 3 states
// 3 doubles lead to a transition to JAIL square (10)

// We start off by generating the probability for each dice roll outcome
// Each dice outcome is 1 / 36, we check for doubles and mark the sum for each outcome
// i + roll % 40 -> next position

// To ensure we can track the transition where landing on a square moves to another square
// we need to keep track of the square index
// We capture the probability of landing on that square and add it to the next state

// Every Chance (CH), Community Chest (CC) square has a transition to a different square
// Transit squares are counted as 10 for CH and 2 for CC, there are 16 cards in total
// This leaves 14 cards with no transition for CC cards and 6 for CH cards
// We can then add the probability of landing on that square to the next state

// We can then build the matrix by iterating over all possible states
// If we move to JAIL, we transition to that state directly
// Ignore any doubles as the end state is independent
// Otherwise, we add the probability of landing on that square
// Any addition for doubles will be done before we generate the next step
func buildMatrix() [][]float64 {
	arr := make([][]float64, MAX_STATES)
	for i := 0; i < MAX_STATES; i++ {
		arr[i] = make([]float64, MAX_STATES)
	}

	// Fixed constants
	rollOutcomes := diceProbability()
	transitionMap := generateTransitMap()

	// We ignore the doubles for JAIL as it is independent of the states
	// We do not consider 3 doubles as a separate state in the markov chain
	// We can store it in any state for JAIL, I chose 0 (no doubles)
	jailIndex := stateIndex(JAIL, 0)

	for i := 0; i < SIZE; i++ {
		for k := 0; k < MAX_ROLLS; k++ {
			from := stateIndex(i, k)

			for _, rollOutcome := range rollOutcomes {
				// Step that will never happen
				if rollOutcome.probability == 0 {
					continue
				}

				doubles := 0
				if rollOutcome.isDouble {
					doubles = k + 1
				}

				// Assume the starting position is anywhere on the board
				// Along with the sum of the rolls, round that up to the next 40 if the user passes GO
				nextPos := (rollOutcome.sum + i) % SIZE

				if doubles == MAX_ROLLS || nextPos == GO_TO_JAIL {
					// Go to jail
					arr[from][jailIndex] += rollOutcome.probability
					continue
				}

				// Index taking into consideration the number of doubles in the state
				aggregateIndex := stateIndex(nextPos, doubles)

				if hasArr(CC, nextPos) {
					// 14 because those are the cards left that do not move the player
					arr[from][aggregateIndex] += 14 * rollOutcome.probability / (16.0)
				} else if hasArr(CH, nextPos) {
					// 6 because those are the cards left that do not move the player
					arr[from][aggregateIndex] += 6 * rollOutcome.probability / (16.0)
				} else {
					// If normal, just add the probability
					arr[from][aggregateIndex] += rollOutcome.probability
				}

				for _, nextStep := range transitionMap[nextPos] {
					nextStepIndex := stateIndex(nextStep.square, doubles)
					arr[from][nextStepIndex] += nextStep.probability * rollOutcome.probability
				}
			}
		}
	}

	return arr
}

// Uses power iteration to find stationary distribution
// https://datascience.oneoffcoder.com/markov-chain-stationary-distribution.html
// https://en.wikipedia.org/wiki/Power_iteration
func stationaryDistribution(M [][]float64) []float64 {
	a := make([]float64, MAX_STATES)
	for i := 0; i < MAX_STATES; i++ {
		a[i] = 1 / MAX_STATES
	}

	for iter := 0; iter < MAX_ITERATIONS; iter++ {
		newA := make([]float64, MAX_STATES)

		// Dot product of M * a
		for i := 0; i < MAX_STATES; i++ {
			for j := 0; j < MAX_STATES; j++ {
				newA[j] += M[i][j] * a[i]
			}
		}

		// The normalization step is not necessary for convergence
		// It can help with stability, but the ratios are the same
		norm := 0.0
		for _, n := range newA {
			norm += n * n
		}

		norm = math.Sqrt(norm)

		for i := 0; i < MAX_STATES; i++ {
			newA[i] /= norm
		}

		a = newA
	}

	return a
}

func main() {
	// Generate Markov chain matrix
	arr := buildMatrix()

	// Find stationary distribution using the power iteration method
	dist := stationaryDistribution(arr)

	// Generate distribution for each square
	// Add up the probabilities for each square and double state
	var newArr [SIZE]float64
	for idx, prob := range dist {
		square := idx / MAX_ROLLS
		newArr[square] += prob
	}

	// Sort the distribution by probability
	var sortableDistribution [][2]float64
	for idx, prob := range newArr {
		sortableDistribution = append(sortableDistribution, [2]float64{float64(idx), prob})
	}

	sort.Slice(sortableDistribution, func(i, j int) bool {
		return sortableDistribution[i][1] > sortableDistribution[j][1]
	})

	result := ""
	for i := 0; i < 3; i++ {
		result += fmt.Sprintf("%02d", int(sortableDistribution[i][0]))
	}

	fmt.Println(result)
}
