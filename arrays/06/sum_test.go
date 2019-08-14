package arrays

import (
	"fmt"
	"reflect"
	"testing"
)

func checkSum(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d given", got, want)
	}
}

func checkSums(t *testing.T, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		checkSum(t, got, want)
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		checkSum(t, got, want)
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}
	checkSums(t, got, want)
}

func TestSumAllTails(t *testing.T) {

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9}, []int{5, 8, 10})
		want := []int{2, 9, 18}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum([]int{1, 2, 3, 4, 5})
	}
}

func BenchmarkSumAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAll([]int{1, 2}, []int{0, 9})
	}
}

func BenchmarkSumAllTails(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAllTails([]int{1, 2}, []int{0, 9}, []int{5, 8, 10})
	}
}

func ExampleSum() {
	sum := Sum([]int{1, 2, 3, 4, 5})
	fmt.Println(sum)
	// Output: 15
}

func ExampleSumAll_case1() {
	sum1 := SumAll([]int{1, 2}, []int{0, 9})
	fmt.Println(sum1)
	// Output: [3 9]
}

func ExampleSumAll_case2() {
	sum2 := SumAll([]int{}, []int{0, 9})
	fmt.Println(sum2)
	// Output: [0 9]
}

func ExampleSumAllTails() {
	sum := SumAllTails([]int{1, 2}, []int{0, 9}, []int{5, 8, 10})
	fmt.Println(sum)
	// Output: [2 9 18]
}
