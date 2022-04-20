package main

import "testing"

func gridPoints(dimensions int, sizes []int) (points [][]int) {
	numCombos := 1
	for i := 0; i < len(sizes); i++ {
		numCombos *= sizes[i]
	}
	for i := 0; i < numCombos; i++ {
		points = append(points, []int{})
		for d := 0; d < dimensions; d++ {
			for s := 0; s < sizes[d]; s++ {
				points[i] = append(points[i], s)
			}
		}
	}
	return
}

func TestExpandGrid(t *testing.T) {
	grid := map[[4]int]bool{
		[4]int{0, 0, 0, 0}: false,
		[4]int{1, 0, 0, 0}: false,
		[4]int{2, 0, 0, 0}: false,
		[4]int{0, 1, 0, 0}: false,
		[4]int{1, 1, 0, 0}: false,
		[4]int{2, 1, 0, 0}: false,
		[4]int{0, 2, 0, 0}: false,
		[4]int{1, 2, 0, 0}: false,
		[4]int{2, 2, 0, 0}: false,
	}

	want := gridPoints(4, []int{5, 5, 3, 3})

	got := ExpandGrid(grid)

	if len(want) != len(got) {
		t.Errorf("Length of expanded grid (%v) does not match expected (%v)\n",
			len(got), len(want))
	}
	for _, pos := range want {
		arr := [4]int{}
		copy(arr[:], pos)
		found := false
		for kk, _ := range got {
			if kk == arr {
				found = true
			}
		}
		if !found {
			t.Errorf("Value (%v) not found in expanded grid!\n", arr)
		}
	}
}
