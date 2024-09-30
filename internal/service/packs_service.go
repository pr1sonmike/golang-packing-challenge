package service

import (
	"context"
	"errors"
	"math"
	"sort"
)

// PacksService defines the interface for the service
type PacksService interface {
	FulfillItems(ctx context.Context, number int, packs []int) (map[int]int, error)
}

// packsService is the implementation of PacksService
type packsService struct{}

// NewPacksService creates a new instance of PacksService
func NewPacksService() PacksService {
	return packsService{}
}

// FulfillItems calculates the packs needed to fulfill the given number
func (ps packsService) FulfillItems(ctx context.Context, number int, packs []int) (map[int]int, error) {
	// Validate input
	if number <= 0 {
		return nil, errors.New("number must be greater than zero")
	}
	if len(packs) == 0 {
		return nil, errors.New("packs list cannot be empty")
	}

	// Calculate the maximum result value
	maxResult := number + minElement(packs) + 1
	answer := maxResult

	// Initialize the dp array with -1 and set dp[0] to 0
	dp := make([]int, maxResult)
	dp[0] = 0
	for i := 1; i < maxResult; i++ {
		dp[i] = -1
	}

	// Update the dp array based on the packs
	for _, i := range packs {
		for j := i; j < maxResult; j++ {
			if dp[j-i] == -1 {
				continue
			}
			if dp[j] == -1 {
				dp[j] = dp[j-i] + 1
			} else {
				dp[j] = int(math.Min(float64(dp[j]), float64(dp[j-i]+1)))
			}
			if j >= number {
				answer = int(math.Min(float64(answer), float64(j)))
			}
		}
	}

	// Construct the result map by backtracking through the dp array
	result := make(map[int]int)
	for answer != 0 {
		for _, i := range packs {
			if answer >= i && dp[answer-i]+1 == dp[answer] {
				answer = answer - i
				result[i]++
				break
			}
		}
	}

	return result, nil
}

func minElement(arr []int) int {
	sort.Ints(arr)
	return arr[0]
}
