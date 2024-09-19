package Remove_Element

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func Test(t *testing.T) {
	var nums = []int{2, 3, 3, 2}     // Input array
	var val = 3                      // Value to remove
	var expectedNums = []int{2, 2}   // The expected answer with correct length.
	var k = removeElement(nums, val) // Calls your implementation

	assert.Equal(t, k, len(expectedNums))
	sort.Ints(nums)
	for i := 0; i < len(expectedNums); i++ {
		assert.Equal(t, nums[i], expectedNums[i])
	}
}
func removeElement(nums []int, val int) int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[j] = nums[i]
			j++
		}
	}
	return j
}
