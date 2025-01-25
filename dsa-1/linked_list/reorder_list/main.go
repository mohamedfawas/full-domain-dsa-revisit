package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode)  {
	if head==nil || head.Next == nil || head.Next.Next==nil {
		return
	}

	slow, fast := head,head
	for fast.Next !=nil && fast.Next.Next !=nil{
		slow = slow.Next
		fast = fast.Next.Next 
	}

	secondHalf := slow.Next

	slow.Next = nil

	var prev *ListNode= nil
	curr := secondHalf
	for curr !=nil{
		nextTemp := curr.Next
		
	}
}
