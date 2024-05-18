package sounds

import (
	"context"
	"github.com/marcussss1/simplevault/internal/model"
	"mime/multipart"
	"os"
)

type tarantoolRepository interface {
	GetAllSounds(ctx context.Context) ([]model.Sound, error)
	GetRandomSounds(ctx context.Context) ([]model.Sound, error)
	GetLastGeneratedUserSounds(ctx context.Context, userID string) ([]model.Sound, error)
	StoreSound(ctx context.Context, sound model.Sound) error
	DeleteSoundByID(ctx context.Context, soundID string) error
	EditSoundByID(ctx context.Context, sound model.EditSound) error
	LikeSoundByID(ctx context.Context, authorID, soundID string) error
	GetLikedSounds(ctx context.Context, userID string) ([]model.Sound, error)
}

type mlClient interface {
	GenerateSoundByText(ctx context.Context, text, duration string) (*os.File, error)
	GenerateSoundByImageURL(ctx context.Context, imageURL string) (*os.File, error)
	GenerateSoundByAudioURL(ctx context.Context, audioURL string) (*os.File, error)
}

type minioRepository interface {
	UploadSound(ctx context.Context, file multipart.File, filename string, size int64) error
}
