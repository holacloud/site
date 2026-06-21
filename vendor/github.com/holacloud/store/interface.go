package store

import (
	"context"
	"errors"
)

type Identifier interface {
	GetId() string
	GetVersion() int64
	SetVersion(version int64)
}

var ErrVersionGone = errors.New("version gone")

type Storer[T Identifier] interface {
	List(ctx context.Context) ([]*T, error)
	Put(ctx context.Context, item *T) error
	Get(ctx context.Context, id string) (*T, error)
	Delete(ctx context.Context, id string) error
}
