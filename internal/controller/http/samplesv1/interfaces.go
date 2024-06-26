package samplesv1

import (
	"context"
	"mime/multipart"

	"github.com/marcussss1/simplevault/internal/model"
	"github.com/minio/minio-go/v7"
)

//go:generate mockgen -source=interfaces.go -destination=./mock/interfaces.go -package=mock

type soundsService interface {
	GetAllSounds(ctx context.Context) ([]model.Sound, error)
	GetRandomSounds(ctx context.Context) ([]model.Sound, error)
	GetLastGeneratedUserSounds(ctx context.Context, userID string) ([]model.Sound, error)
	GenerateSoundByText(ctx context.Context, text, duration, userID, sessionID string) (model.Sound, error)
	GenerateSoundByImageURL(ctx context.Context, imageURL string, userID string) (model.Sound, error)
	GenerateSoundByAudioURL(ctx context.Context, audioURL string, userID string) (model.Sound, error)
	DeleteSoundByID(ctx context.Context, soundID string) error
	EditSoundByID(ctx context.Context, sound model.EditSound) error
	LikeSoundByID(ctx context.Context, authorID, soundID string) error
	DislikeSoundByID(ctx context.Context, authorID, soundID string) error
	GetLikedSounds(ctx context.Context, userID string) ([]model.Sound, error)
}

type playlistsService interface {
	GetRandomPlaylists(ctx context.Context) ([]model.RandomPlaylist, error)
	GetPlaylist(ctx context.Context, id string) (model.FullPlaylist, error)
}

type filesService interface {
	DownloadSound(ctx context.Context, filename string) (*minio.Object, error)
	UploadSound(ctx context.Context, file multipart.File, header *multipart.FileHeader, userID string, uploadSound model.UploadSound) (model.Sound, error)
}

type authService interface {
	Signup(ctx context.Context, signupUserRequest model.SignupUserRequest) (model.SignupUserResponse, error)
	Login(ctx context.Context, loginUserRequest model.LoginUserRequest) (model.LoginUserResponse, error)
	DeleteSessionByID(ctx context.Context, sessionID string) error
	GetUserBySessionID(ctx context.Context, sessionID string) (model.GetUserBySessionID, error)
}
