package ml

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io"
	"net/http"
	"os"
	"time"
)

func (c Client) GenerateSoundByText(ctx context.Context, text, duration, sessionID string) (*os.File, error) {
	// todo вынести в общую
	req, err := http.NewRequestWithContext(context.Background(), "GET", c.Host+"generate_by_text", nil)
	if err != nil {
		return nil, fmt.Errorf("create request 1: %w", err)
	}

	q := req.URL.Query()
	q.Add("input_text", text)
	q.Add("duration", duration)
	q.Add("session_id", sessionID)
	req.URL.RawQuery = q.Encode()

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		Timeout: 10 * time.Minute,
	}

	resp1, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request 1: %w", err)
	}
	defer resp1.Body.Close()

	body, err := io.ReadAll(resp1.Body)
	if err != nil {
		return nil, fmt.Errorf("read all resp 1: %w", err)
	}

	fmt.Println()
	fmt.Println(string(body)) // Выводим тело ответа в формате JSON
	fmt.Println()

	type ml1Response struct {
		AudioURL string `json:"audio_url"`
	}
	var mlResp ml1Response
	err = json.Unmarshal(body, &mlResp)
	if err != nil {
		return nil, fmt.Errorf("decode response 1: %w", err)
	}

	req, err = http.NewRequestWithContext(context.Background(), "GET", mlResp.AudioURL, nil)
	if err != nil {
		return nil, fmt.Errorf("create request 2: %w", err)
	}

	resp2, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request 2: %w", err)
	}
	defer resp2.Body.Close()

	file, err := os.Create(uuid.NewString() + ".wav")
	if err != nil {
		return nil, fmt.Errorf("create file 2: %w", err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp2.Body)
	if err != nil {
		return nil, fmt.Errorf("copy file 2: %w", err)
	}

	return file, nil
}
