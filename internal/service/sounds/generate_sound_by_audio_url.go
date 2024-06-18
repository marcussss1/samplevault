package sounds

import (
	"context"
	"fmt"
	"github.com/marcussss1/simplevault/internal/model"
)

func (s Service) GenerateSoundByAudioURL(ctx context.Context, audioURL, userID string) (model.Sound, error) {
	// todo вынести в общую
	_, err := s.mlClient.GenerateSoundByAudioURL(ctx, audioURL)
	if err != nil {
		return model.Sound{}, fmt.Errorf("genereate sound by audio url from ml client: %w", err)
	}

	return model.Sound{}, nil
	//audioFile, err = os.Open(audioFile.Name())
	//if err != nil {
	//	return model.Sound{}, fmt.Errorf("open audio file: %w", err)
	//}
	//
	//audioFileInfo, err := audioFile.Stat()
	//if err != nil {
	//	return model.Sound{}, fmt.Errorf("get audio file info: %w", err)
	//}
	//defer audioFile.Close()
	//
	//err = s.minioRepository.UploadSound(ctx, audioFile, audioFile.Name(), audioFileInfo.Size())
	//if err != nil {
	//	return model.Sound{}, fmt.Errorf("upload sound from minio repository: %w", err)
	//}
	//
	//sample := newSample(userID, audioFile.Name())
	//err = s.tarantoolRepository.StoreSound(ctx, sample)
	//if err != nil {
	//	return model.Sound{}, fmt.Errorf("store sound from tarantool repository: %w", err)
	//}
	//
	//return sample, nil
}
