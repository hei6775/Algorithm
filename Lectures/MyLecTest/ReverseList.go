package MyLecTest

import "fmt"

type Point struct {
	Val int
	Next *Point
}
//反转链表
func ReverseP(p *Point) (*Point) {
	if p == nil || p.Next == nil {
		return p
	}
	var pre *Point
	pre = nil
	next := new(Point)
	next = nil
	//a b c
	//next=b  p.Next=nil
	for p != nil {
		next = p.Next
		p.Next = pre
		pre = p
		p = next
	}
	return pre
}

func PrintP(p *Point){
	for p != nil {
		fmt.Print(p.Val)
		p = p.Next
	}
}