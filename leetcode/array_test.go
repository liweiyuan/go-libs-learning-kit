package leetcode

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Two Sum
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := range nums {
		another := target - nums[i]
		if _, ok := m[another]; ok {
			return []int{m[another], i}
		}
		m[nums[i]] = i
	}
	return nil
}

func twoSum2(nums []int, target int) []int {
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
		{[]int{1, 2, 3, 4, 5}, 10, nil},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := twoSum(tt.nums, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTwoSum2(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
		{[]int{1, 2, 3, 4, 5}, 10, nil},
		{[]int{2, 4, 5}, 9, []int{1, 2}},
		{[]int{2, 4, 5}, 10, nil},
		{[]int{2, 4, 5}, 7, []int{0, 2}},
		{[]int{2, 4, 5}, 12, nil},
		{[]int{2, 4, 5}, 13, nil},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := twoSum2(tt.nums, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum2() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Three Sum
func threeSum(nums []int) [][]int {
	result := [][]int{}
	if len(nums) < 3 {
		return result
	}

	// 排序数组
	quickSort(nums, 0, len(nums)-1)

	for i := 0; i < len(nums)-2; i++ {
		// 跳过重复元素
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				// 跳过重复元素
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return result
}

func quickSort(nums []int, start, end int) {
	if start >= end {
		return
	}
	pivot := nums[start]
	left, right := start, end
	for left < right {
		for left < right && nums[right] >= pivot {
			right--
		}
		nums[left] = nums[right]
		for left < right && nums[left] <= pivot {
			left++
		}
		nums[right] = nums[left]
	}
	nums[left] = pivot
	quickSort(nums, start, left-1)
	quickSort(nums, left+1, end)
}

func TestThreeSum(t *testing.T) {
	tests := []struct {
		nums []int
		want [][]int
	}{
		{
			nums: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			nums: []int{},
			want: [][]int{},
		},
		{
			nums: []int{0},
			want: [][]int{},
		},
		{
			nums: []int{0, 0, 0},
			want: [][]int{{0, 0, 0}},
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := threeSum(tt.nums)
			assert.Equal(t, tt.want, got)
		})
	}
}

// Maximum Subarray
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxSum := nums[0]
	currentSum := nums[0]

	for i := 1; i < len(nums); i++ {
		currentSum = max(nums[i], currentSum+nums[i])
		maxSum = max(maxSum, currentSum)
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{1}, 1},
		{[]int{5, 4, -1, 7, 8}, 23},
		{[]int{-1}, -1},
		{[]int{-2, -1}, -1},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := maxSubArray(tt.nums)
			assert.Equal(t, tt.want, got)
		})
	}
}

// Merge Sorted Array
func merge(nums1 []int, m int, nums2 []int, n int) {
	p1, p2 := m-1, n-1
	p := m + n - 1

	for p2 >= 0 {
		if p1 >= 0 && nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			nums1[p] = nums2[p2]
			p2--
		}
		p--
	}
}

func TestMerge(t *testing.T) {
	tests := []struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
		want  []int
	}{
		{
			nums1: []int{1, 2, 3, 0, 0, 0},
			m:     3,
			nums2: []int{2, 5, 6},
			n:     3,
			want:  []int{1, 2, 2, 3, 5, 6},
		},
		{
			nums1: []int{1},
			m:     1,
			nums2: []int{},
			n:     0,
			want:  []int{1},
		},
		{
			nums1: []int{0},
			m:     0,
			nums2: []int{1},
			n:     1,
			want:  []int{1},
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			merge(tt.nums1, tt.m, tt.nums2, tt.n)
			assert.Equal(t, tt.want, tt.nums1)
		})
	}
}

// Remove Duplicates
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	k := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 1, 2}, 2},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{1, 1}, 1},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := removeDuplicates(tt.nums)
			assert.Equal(t, tt.want, got)
		})
	}
}

// Rotate Array
func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

func TestRotate(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
		{[]int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100}},
		{[]int{1}, 0, []int{1}},
		{[]int{1, 2}, 3, []int{2, 1}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			nums := make([]int, len(tt.nums))
			copy(nums, tt.nums)
			rotate(nums, tt.k)
			assert.Equal(t, tt.want, nums)
		})
	}
}

// Plus One
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	result := make([]int, len(digits)+1)
	result[0] = 1
	return result
}

func TestPlusOne(t *testing.T) {
	tests := []struct {
		digits []int
		want   []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 4}},
		{[]int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
		{[]int{9}, []int{1, 0}},
		{[]int{9, 9}, []int{1, 0, 0}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := plusOne(tt.digits)
			assert.Equal(t, tt.want, got)
		})
	}
}

// Move Zeroes
func moveZeroes(nums []int) {
	nonZero := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[nonZero], nums[i] = nums[i], nums[nonZero]
			nonZero++
		}
	}
}

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{[]int{0}, []int{0}},
		{[]int{1}, []int{1}},
		{[]int{0, 0, 1}, []int{1, 0, 0}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			nums := make([]int, len(tt.nums))
			copy(nums, tt.nums)
			moveZeroes(nums)
			assert.Equal(t, tt.want, nums)
		})
	}
}

// Intersection of Arrays
func intersection(nums1 []int, nums2 []int) []int {
	set := make(map[int]bool)
	for _, num := range nums1 {
		set[num] = true
	}

	result := []int{}
	seen := make(map[int]bool)
	for _, num := range nums2 {
		if set[num] && !seen[num] {
			result = append(result, num)
			seen[num] = true
		}
	}
	return result
}

func TestIntersection(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
		want  []int
	}{
		{[]int{1, 2, 2, 1}, []int{2, 2}, []int{2}},
		{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{9, 4}},
		{[]int{1, 2}, []int{3, 4}, []int{}},
		{[]int{}, []int{1}, []int{}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := intersection(tt.nums1, tt.nums2)
			// 由于结果顺序可能不同，需要先排序再比较
			quickSort(got, 0, len(got)-1)
			want := make([]int, len(tt.want))
			copy(want, tt.want)
			quickSort(want, 0, len(want)-1)
			assert.Equal(t, want, got)
		})
	}
}
