package hw04lrucache

import "fmt"

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Prev  *ListItem
	Next  *ListItem
}

type list struct {
	First *ListItem
	Last  *ListItem
	Size  int
}

func (l *list) Len() int {
	return l.Size
}

func (l *list) Front() *ListItem {
	return l.First
}

func (l *list) Back() *ListItem {
	return l.Last
}

func (l *list) PushFront(v interface{}) *ListItem {
	defer fmt.Printf("В начало добавлен элемент %s\n", v)
	if l.First == nil {
		item := &ListItem{v, nil, nil}
		l.First = item
		l.Last = item
		l.Size++
		return l.First
	}

	secondItem := l.First
	item := ListItem{
		Value: v,
		Next:  secondItem,
		Prev:  nil,
	}
	secondItem.Prev = &item
	l.First = &item
	l.Size++
	return l.First
}

func (l *list) PushBack(v interface{}) *ListItem {
	defer fmt.Printf("В конец добавлен элемент %s\n", v)
	if l.Last == nil {
		item := &ListItem{v, nil, nil}
		l.First = item
		l.Last = item
		l.Size++
		return l.Last
	}

	secondToLast := l.Last
	item := ListItem{
		Value: v,
		Next:  nil,
		Prev:  secondToLast,
	}
	secondToLast.Next = &item
	l.Last = &item
	l.Size++
	return l.Last
}

func (l *list) Remove(i *ListItem) {
	switch {
	case l.Size == 1:
		l.First = nil
		l.Last = nil
	case i == l.First:
		l.First = l.First.Next
		l.First.Prev = nil
	case i == l.Last:
		l.Last = l.Last.Prev
		l.Last.Next = nil
	default:
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
	}
	l.Size--
	fmt.Printf("Удален элемент %s\n", i.Value)
}

func (l *list) MoveToFront(i *ListItem) {
	defer fmt.Printf("Элемент %s переставлен в начало\n", i.Value)
	if i == l.First {
		return
	}
	if i.Next != nil {
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
	} else {
		i.Prev.Next = nil
		l.Last = i.Prev
	}
	i.Next = l.First
	l.First.Prev = i
	i.Prev = nil
	l.First = i
}

func NewList() List {
	return new(list)
}
