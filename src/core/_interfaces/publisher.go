package interfaces

import "context"

type Publisher interface {
	Publish(ctx context.Context, message string) error
}
