package internal

//linked list for a free list level
type freeListElement struct {
	freeAddr uint64
	next *freeListElement
}

func pushBack(list *freeListElement, fAddr uint64) {
	for list.next != nil {
		list = list.next
	}
	list.next = &freeListElement{fAddr,nil}
}