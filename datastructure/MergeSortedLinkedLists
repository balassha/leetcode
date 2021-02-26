package main

import (
	"fmt"
	"sort"
)

type linkedList struct{
    head *ListNode
}

type ListNode struct {
    Val int
    Next *ListNode
 }

func main() {
	fmt.Println("Hello, playground")
	a1 := []int{1,2,4}
	a2 := []int{1,3,4}
	l1 := createLinkedList(a1)
	l2 := createLinkedList(a2)
	res := mergeTwoLists(l1.head,l2.head)
	for res.Next != nil{
		fmt.Println(res.Val)
		res = res.Next
	}
}

func createLinkedList(ar []int) linkedList {
	var l linkedList
	for _,v := range ar{
		if l.head == nil{
			l.head = &ListNode{
				Val:v,
				Next:nil,
			}
		}else{
			t := l.head
			for t.Next != nil {
				t = t.Next
			}
			t.Next = &ListNode{
				Val:v,
				Next:nil,
			}
		}
	}
	return l
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    var ar []int
    for l1 != nil {
        ar = append(ar,l1.Val)
        l1 = l1.Next
    }
    for l2 != nil {
        ar = append(ar,l2.Val)
        l2 = l2.Next
    }
    sort.Ints(ar)
    var res linkedList
    for _,v := range ar{
        if res.head == nil{
            res.head = &ListNode{
                Val:v,
                Next:nil,
            }
        }else{
            t := res.head
            for t.Next != nil {
                t = t.Next
            }
            t.Next = &ListNode{
                Val:v,
                Next:nil,
            }
        }
    }
    return res.head
}
