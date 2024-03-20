package playlists

import (
	"context"
	"fmt"
	"github.com/marcussss1/simplevault/internal/model"
)

func (r Repository) GetAllPlaylists(ctx context.Context, userID string) ([]model.Playlist, error) {
	var playlists []model.Playlist

	err := r.conn.EvalTyped(`
		local result = {}
		local playlists = box.space.playlists:select({}, {limit = 3})
		for _, playlist in ipairs(playlists) do
			local playlist_id = playlist[1]
			local author_id = playlist[2]
			local icon_url = playlist[3]
			local created_at = playlist[4]
			local title = playlist[5]
			local description = playlist[6]
			local sound_ids = playlist[7]
		
			local playlistData = {
				id = playlist_id,
				author_id = author_id,
				icon_url = icon_url,
				created_at = created_at,
				title = title,
				description = description,
				sounds = {}
			}
		
			for _, sound_id in ipairs(sound_ids) do
				local sound = box.space.sounds:select{sound_id}
				table.insert(playlistData.sounds, {
					id = sound[1],
					author_id = sound[2],
					audio_url = sound[3],
					icon_url = sound[4],
					file_name = sound[5],
					created_at = sound[6],
					title = sound[7],
					musical_instrument = sound[8],
					genre = sound[9],
					mood = sound[10],
					tonality = sound[11],
					tempo = sound[12],
					style = sound[13]
				})
			end
		
			table.insert(result, playlistData)
		end

		return result
	`, nil, &playlists)
	if err != nil {
		return nil, fmt.Errorf("select all playlists from tarantool storage: %w", err)
	}

	return playlists, nil
}
