package service

import "context"

type HashPasswordService interface {
	HashPassword(ctx context.Context, plainPassword string) (string, error)
}
