package samples

import (
	"fmt"
)

func (s Service) GetUrl() (string, error) {
	url, err := s.minioRepository.GetUrl()
	if err != nil {
		return "", fmt.Errorf("get all samples from tarantool repository: %w", err)
	}

	return url, nil
}
