package model

type Sound struct {
	ID                string `json:"id"`
	AuthorID          string `json:"author_id"`
	AudioURL          string `json:"audio_url"`
	IconURL           string `json:"icon_url"`
	FileName          string `json:"file_name"`
	CreatedAt         string `json:"created_at"`
	Title             string `json:"title"`
	MusicalInstrument string `json:"musical_instrument"`
	Genre             string `json:"genre"`
	Mood              string `json:"mood"`
	Tonality          string `json:"tonality"`
	Tempo             string `json:"tempo"`
	Style             string `json:"style"`
	IsGenerated       bool   `json:"is_generated"`
}
