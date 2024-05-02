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
	GenerateSoundByText(ctx context.Context, text string) (model.Sound, error)
	GenerateSoundByImageURL(ctx context.Context, imageURL string) (model.Sound, error)
	GenerateSoundByAudioURL(ctx context.Context, audioURL string) (model.Sound, error)
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
	GetUserBySessionID(ctx context.Context, sessionID string) (model.GetUserBySessionID, error)
}
