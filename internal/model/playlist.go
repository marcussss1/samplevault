package model

type Playlist struct {
	ID          string   `json:"id"`
	AuthorID    string   `json:"author_id"`
	IconURL     string   `json:"icon_url"`
	CreatedAt   string   `json:"created_at"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	SoundIDs    []string `json:"sound_ids"`
}
