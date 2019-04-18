package list

import (
	"fmt"
	"testing"
)

var l *SinglyLinkedList

func TestSinglyLinkedList_InsertFirst(t *testing.T) {
	l = NewSinglyLinkedList()
	l.InsertFirst(110)
	l.InsertFirst(111)
	l.InsertFirst(112)
	fmt.Println(l)
}

func TestSinglyLinkedList_InsertLast(t *testing.T) {
	l.InsertLast(113)
	fmt.Println(l)
}

func TestSinglyLinkedList_GetFirst(t *testing.T) {
	if i, ok := l.GetFirst(); ok {
		fmt.Println(i)
	}
}

func TestSinglyLinkedList_GetLast(t *testing.T) {
	if i, ok := l.GetLast(); ok {
		fmt.Println(i)
	}
}

func TestSinglyLinkedList_GetSize(t *testing.T) {
	fmt.Println(l.GetSize())
}


func TestSinglyLinkedList_GetItems(t *testing.T) {
	fmt.Println(l.GetItems())
}

func TestSinglyLinkedList_RemoveByValue(t *testing.T) {
	fmt.Println(l.RemoveByValue(110))
}

func TestSinglyLinkedList_GetItems2(t *testing.T) {
	fmt.Println(l.GetItems())
}

func TestSinglyLinkedList_RemoveByIndex(t *testing.T) {
	fmt.Println(l.RemoveByIndex(2))
}

func TestSinglyLinkedList_GetItems3(t *testing.T) {
	fmt.Println(l.GetItems())
}

func TestSinglyLinkedList_SearchValue(t *testing.T) {
	fmt.Println(l.SearchValue(110))
	fmt.Println(l.SearchValue(111))
}