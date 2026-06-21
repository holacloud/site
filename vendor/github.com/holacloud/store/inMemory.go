package store

import (
	"context"
	"encoding/json"
	"sync"
	"sync/atomic"
)

type node[T any] struct {
	item atomic.Pointer[T]
	next atomic.Pointer[node[T]]
}

type StoreMemory[T Identifier] struct {
	// head uses atomic.Pointer to ensure memory visibility and ordering.
	// While a pointer assignment on 64-bit systems is atomic in terms of machine instructions (no torn writes),
	// the Go Memory Model does not guarantee that other goroutines will see the updated value instantly
	// or in the correct order without synchronization primitives (like atomic or mutex).
	// atomic.Pointer provides the necessary "happens-before" edges to ensure that initialization of the node
	// happens before the head pointer is visible to readers.
	head  atomic.Pointer[node[T]]
	mutex sync.Mutex
}

func NewStoreMemory[T Identifier]() *StoreMemory[T] {
	return &StoreMemory[T]{}
}

func (f *StoreMemory[T]) List(ctx context.Context) ([]*T, error) {
	// No lock needed for readers
	var result []*T

	current := f.head.Load()
	for current != nil {
		// safe copy
		itemPtr := current.item.Load()
		if itemPtr != nil { // Should be non-nil generally
			var newItem *T
			remarshal(itemPtr, &newItem)
			result = append(result, newItem)
		}

		current = current.next.Load()
	}

	return result, nil
}

func (f *StoreMemory[T]) Put(ctx context.Context, item *T) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	id := (*item).GetId()
	version := (*item).GetVersion()

	// Check if exists
	current := f.head.Load()

	for current != nil {
		currentItem := current.item.Load()
		if (*currentItem).GetId() == id {
			// Found, check version
			if (*currentItem).GetVersion() != version {
				return ErrVersionGone
			}

			// Update
			(*item).SetVersion(version + 1)
			current.item.Store(item)
			return nil

		}
		current = current.next.Load()
	}

	// Not found, insert new
	(*item).SetVersion((*item).GetVersion() + 1)
	newNode := &node[T]{}
	newNode.item.Store(item)
	newNode.next.Store(f.head.Load())

	f.head.Store(newNode)

	return nil
}

func (f *StoreMemory[T]) Get(ctx context.Context, id string) (*T, error) {
	// No lock needed for readers
	current := f.head.Load()
	for current != nil {
		currentItem := current.item.Load()
		if (*currentItem).GetId() == id {
			// Copy
			var newItem *T
			remarshal(currentItem, &newItem)
			return newItem, nil
		}
		current = current.next.Load()
	}

	return nil, nil
}

func remarshal(in, out any) {
	b, _ := json.Marshal(in)
	_ = json.Unmarshal(b, &out)
}

func (f *StoreMemory[T]) Delete(ctx context.Context, id string) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	current := f.head.Load()
	var prev *node[T]

	for current != nil {
		currentItem := current.item.Load()
		if (*currentItem).GetId() == id {
			// Found, remove
			next := current.next.Load()
			if prev == nil {
				// Removing head
				f.head.Store(next)
			} else {
				// Removing from middle/end
				prev.next.Store(next)
			}
			return nil
		}
		prev = current
		current = current.next.Load()
	}

	return nil
}
