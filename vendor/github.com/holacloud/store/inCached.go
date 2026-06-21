package store

import (
	"context"
)

type StoreCached[T Identifier] struct {
	persistence Storer[T] // Persistent storage (e.g., StoreDisk)
	cache       Storer[T] // Caching layer (e.g., StoreMemory)
}

func NewStoreCached[T Identifier](persistence Storer[T], cache Storer[T]) (*StoreCached[T], error) {

	if cache == nil {
		cache = NewStoreMemory[T]()
	}

	// Warm up cache from persistence
	items, err := persistence.List(context.Background())
	if err != nil {
		return nil, err
	}

	for _, item := range items {
		_ = cache.Put(context.Background(), item)
	}

	return &StoreCached[T]{
		persistence: persistence,
		cache:       cache,
	}, nil
}

func (s *StoreCached[T]) List(ctx context.Context) ([]*T, error) {
	return s.cache.List(ctx)
}

func (s *StoreCached[T]) Put(ctx context.Context, item *T) error {
	// 1. Persist first (source of truth)
	if err := s.persistence.Put(ctx, item); err != nil {
		return err
	}
	// 2. Update cache
	return s.cache.Put(ctx, item)
}

func (s *StoreCached[T]) Get(ctx context.Context, id string) (*T, error) {
	// 1. Check cache
	item, err := s.cache.Get(ctx, id)
	if err == nil && item != nil {
		return item, nil
	}

	// 2. Fallback to persistence
	item, err = s.persistence.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	if item == nil {
		return nil, nil // Not found
	}

	// 3. Update cache (read repair / populate)
	_ = s.cache.Put(ctx, item)

	return item, nil
}

func (s *StoreCached[T]) Delete(ctx context.Context, id string) error {
	// 1. Delete from persistence
	if err := s.persistence.Delete(ctx, id); err != nil {
		return err
	}
	// 2. Delete from cache
	return s.cache.Delete(ctx, id)
}
