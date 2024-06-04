package repo

import "context"

type Atomic interface {
	Atomic(ctx context.Context, f func(tx Atomic) error) error
}
