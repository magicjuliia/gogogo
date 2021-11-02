/*
Pointers are not thread save so the functionality should not
be used with Go routines since they multiplex onto OS threads.

Actionable items:
- try to use RWMutex
- use array indexes instead of pointers
- read about tail-recursion in Go
*/

package main

import (
	"fmt"
)

type toDoItem struct {
	text     string
	nextItem *toDoItem
	prevItem *toDoItem
}

func test() {

	item1 := toDoItem{"call my dentist", nil, nil}
	item2 := toDoItem{"go grocery shopping", nil, nil}
	item3 := toDoItem{"clean kitchen", nil, nil}

	addItem(&item2, &item3)
	addItem(&item1, &item2)
	// removeItem(&item1)
	printToDoList(&item1)
	// printToDoListInReverce(&item3)
}

func main() {
	item1 := toDoItem{"call my dentist", nil, nil}
	item2 := toDoItem{"go grocery shopping", nil, nil}
	item3 := toDoItem{"clean kitchen", nil, nil}

	addItem(&item2, &item3)
	addItem(&item1, &item2)
	// removeItem(&item1)
	printToDoList(&item1)
	// printToDoListInReverce(&item3)
}

func printToDoList(item *toDoItem) {
	fmt.Println(item.text)

	if item.nextItem != nil {
		printToDoList(item.nextItem)
	}
}

func printToDoListInReverce(item *toDoItem) {
	isFirstItem := (item.nextItem != nil) && (item.prevItem == nil)

	if isFirstItem {
		fmt.Println(item.text)
		return
	}

	fmt.Println(item.text)
	printToDoListInReverce(item.prevItem)
}

func addItem(newItem *toDoItem, nextItem *toDoItem) {
	nextItem.prevItem = newItem
	if newItem.nextItem == nil {
		newItem.nextItem = nextItem
		return
	}

	nextItemOfTheNewItem := newItem.nextItem
	newItem.nextItem = nextItem
	nextItem.nextItem = nextItemOfTheNewItem
}

func removeItem(item *toDoItem) {
	isFirstItem := (item.nextItem != nil) && (item.prevItem == nil)
	isLastItem := (item.nextItem == nil) && (item.prevItem != nil)
	if isFirstItem {
		item.nextItem.prevItem = nil
		item.nextItem = nil
		return
	}
	if isLastItem {
		item.prevItem.nextItem = nil
		return
	}
	item.prevItem.nextItem = item.nextItem
	item.nextItem.prevItem = item.prevItem
}
