package ml

import (
	"context"
	"os"
)

//go:generate mockgen -source=interfaces.go -destination=./mock/interfaces.go -package=mock

type mlClient interface {
	GenerateSoundByText(ctx context.Context, text string) (*os.File, error)
	GenerateSoundByImageURL(ctx context.Context, imageURL string) (*os.File, error)
	GenerateSoundByAudioURL(ctx context.Context, audioURL string) (*os.File, error)
}
