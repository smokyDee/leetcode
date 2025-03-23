package main

import "fmt"

func main() {
	fmt.Println("hello world")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carrier := 0
	counter := 0
	numberLists := 2
	begin := ListNode{}
	var lastNode *ListNode
	for {
		if numberLists == 0 {
			if carrier != 0 {
				var node ListNode
				node.Val = carrier
				node.Next = nil
				lastNode.Next = &node
			}
			break
		}
		var (
			v1 int
			v2 int
		)
		if l1 != nil {
			v1 = l1.Val
		}
		if l2 != nil {
			v2 = l2.Val
		}

		val := (v1 + v2 + carrier) % 10

		if counter == 0 {
			begin.Val = val
			begin.Next = nil

			lastNode = &begin

		} else {
			var node ListNode
			node.Val = val
			node.Next = nil

			lastNode.Next = &node
			lastNode = &node
		}

		if l1 != nil {
			if l1.Next == nil {
				numberLists--
			}
			l1 = l1.Next
		}

		if l2 != nil {
			if l2.Next == nil {
				numberLists--
			}
			l2 = l2.Next
		}

		carrier = (v1 + v2 + carrier) / 10
		counter++
	}
	return &begin
}
