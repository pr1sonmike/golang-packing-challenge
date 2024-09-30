package service

import (
	"context"
	"reflect"
	"testing"
)

func TestFulfillNumber(t *testing.T) {
	tests := []struct {
		name     string
		number   int
		packs    []int
		expected map[int]int
	}{
		{
			name:     "Test case 1",
			number:   1,
			packs:    []int{5000, 2000, 1000, 500, 250},
			expected: map[int]int{250: 1},
		},
		{
			name:     "Test case 2",
			number:   250,
			packs:    []int{500, 250},
			expected: map[int]int{250: 1},
		},
		{
			name:     "Test case 3",
			number:   251,
			packs:    []int{300, 250},
			expected: map[int]int{300: 1},
		},
		{
			name:     "Test case 4",
			number:   501,
			packs:    []int{1000, 500, 250},
			expected: map[int]int{500: 1, 250: 1},
		},
		{
			name:     "Test case 5",
			number:   1201,
			packs:    []int{500, 200, 100, 50, 25},
			expected: map[int]int{500: 2, 200: 1, 25: 1},
		},
		{
			name:     "Test case 6",
			number:   1249,
			packs:    []int{500, 200, 100, 50, 25},
			expected: map[int]int{500: 2, 200: 1, 50: 1},
		},
		{
			name:     "Test case 7",
			number:   1250,
			packs:    []int{500, 200, 100, 50, 25},
			expected: map[int]int{500: 2, 200: 1, 50: 1},
		},
		{
			name:     "Test case 8",
			number:   1251,
			packs:    []int{500, 200, 100, 50, 25},
			expected: map[int]int{500: 2, 200: 1, 50: 1, 25: 1},
		},
		{
			name:     "Test case 9",
			number:   999,
			packs:    []int{1000, 250},
			expected: map[int]int{1000: 1},
		},
		{
			name:   "Test case 10",
			number: 11,
			packs:  []int{3, 7, 6},
			expected: map[int]int{
				6: 2,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := NewPacksService()
			result, _ := ps.FulfillItems(context.Background(), tt.number, tt.packs)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("FulfillItems(%d, %v) = %v; expected %v", tt.number, tt.packs, result, tt.expected)
			}
		})
	}
}
