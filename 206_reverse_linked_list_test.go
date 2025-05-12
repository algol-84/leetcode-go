package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head
	for _, v := range nums[1:] {
		current.Next = &ListNode{Val: v}
		current = current.Next
	}
	return head
}

func ListToSlice(head *ListNode) []int {
	slice := make([]int, 0)
	currentNode := head
	for currentNode != nil {
		slice = append(slice, currentNode.Val)
		currentNode = currentNode.Next
	}
	return slice
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	curr := head

	for curr != nil {
		nextTemp := curr.Next
		curr.Next = prev
		prev = curr
		curr = nextTemp
	}
	return prev
}

func main() {
	var j string
	var s string

	fmt.Scanln(&j)
	fmt.Scanln(&s)

	m := make(map[byte]struct{})
	for _, i := range []byte(j) {
		m[i] = struct{}{}
	}

	cnt := 0
	for _, i := range []byte(s) {
		if _, ok := m[i]; ok {
			cnt++
		}
	}
}

func TestReverseLinkedList(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{1, 2, 3, 4},
			expected: []int{4, 3, 2, 1},
		},
		{
			input:    []int{1},
			expected: []int{1},
		},
	}

	for _, test := range tests {
		fmt.Println(test.input)
		list := NewList(test.input)
		reversedList := reverseList(list)
		fmt.Println(ListToSlice(reversedList))

		if !reflect.DeepEqual(ListToSlice(reversedList), test.expected) {
			t.Errorf("got %v, but wanted %v\n", ListToSlice(reversedList), test.expected)
		}
	}
}
