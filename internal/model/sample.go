package model

type Sample struct {
	ID                string `json:"id"`
	AuthorID          string `json:"author_id"`
	AudioURL          string `json:"audio_url"`
	IconURL           string `json:"icon_url"`
	FileName          string `json:"file_name"`
	Title             string `json:"title"`
	Duration          string `json:"duration"`
	MusicalInstrument string `json:"musical_instrument"`
	Genre             string `json:"genre"`
	IsFavourite       bool   `json:"is_favourite"`
	CreatedAt         string `json:"created_at"`
}
