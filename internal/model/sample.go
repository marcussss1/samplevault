package model

// TODO ссылка с s3 на трек
type Sample struct {
	ID                string `json:"id"`
	AuthorID          string `json:"author_id"`
	AudioURL          string `json:"audio_url"`
	IconURL           string `json:"icon_url"`
	Title             string `json:"title"`
	Duration          string `json:"duration"`
	MusicalInstrument string `json:"musical_instrument"`
	Genre             string `json:"genre"`
	IsFavourite       bool   `json:"is_favourite"`
}
