package sounds

import (
	"fmt"
	"github.com/marcussss1/simplevault/internal/model"
	"github.com/tarantool/go-tarantool"
	"strconv"
)

func toSounds(resp *tarantool.Response) []model.Sound {
	var sounds []model.Sound

	for _, tuples := range resp.Tuples() {
		for _, tuple := range tuples {
			sound := tuple.([]interface{})

			isGenerated, err := strconv.ParseBool(fmt.Sprint(sound[13]))
			if err != nil {
				panic(err)
			}

			sounds = append(sounds, model.Sound{
				ID:                fmt.Sprint(sound[0]),
				AuthorID:          fmt.Sprint(sound[1]),
				AudioURL:          fmt.Sprint(sound[2]),
				IconURL:           fmt.Sprint(sound[3]),
				FileName:          fmt.Sprint(sound[4]),
				CreatedAt:         fmt.Sprint(sound[5]),
				Title:             fmt.Sprint(sound[6]),
				MusicalInstrument: fmt.Sprint(sound[7]),
				Genre:             fmt.Sprint(sound[8]),
				Mood:              fmt.Sprint(sound[9]),
				Tonality:          fmt.Sprint(sound[10]),
				Tempo:             fmt.Sprint(sound[11]),
				Style:             fmt.Sprint(sound[12]),
				IsGenerated:       isGenerated,
			})
		}
	}

	return sounds
}
