package sounds

import (
	"context"
	"fmt"
	"github.com/marcussss1/simplevault/internal/model"
)

const a = `function get_random_records()
			local space = box.space.sounds
			local all_records = space:select()
			local random_records = {}
			local random_indexes = {}
		
			for i = 1, 5 do
				table.insert(random_indexes, math.random(1, #all_records))
			end
		
			for _, index in ipairs(random_indexes) do
				table.insert(random_records, all_records[index])
			end
		
			return random_records
		end`

func (r Repository) GetRandomSounds(ctx context.Context) ([]model.Sound, error) {
	var sounds []model.Sound

	resp, err := r.conn.Call("get_random_records", nil)
	if err != nil {
		return nil, fmt.Errorf("select random sounds from Tarantool storage: %w", err)
	}

	fmt.Println(resp.String())
	//// разбить input на массивы строк
	//items := strings.Split(resp.String(), " ")
	//
	//// пройти по каждому элементу и создать структуру Sound
	//for _, item := range items {
	//	parts := strings.Split(item, " ") // предполагаем, что элементы разделены пробелами
	//	if len(parts) != 13 {
	//		return nil, errors.New("неправильный формат элемента")
	//	}
	//
	//	sound := Sound{
	//		ID:                parts[0],
	//		AuthorID:          parts[1],
	//		AudioURL:          parts[2],
	//		IconURL:           parts[3],
	//		FileName:          parts[4],
	//		CreatedAt:         parts[5],
	//		Title:             parts[6],
	//		MusicalInstrument: parts[7],
	//		Genre:             parts[8],
	//		Mood:              parts[9],
	//		Tonality:          parts[10],
	//		Tempo:             parts[11],
	//		Style:             parts[12],
	//	}
	//
	//	sounds = append(sounds, sound)
	//}

	//fmt.Println(resp.Data)

	//for _, item := range resp.Data {
	//	soundData, ok := item.([]interface{})
	//	if !ok || len(soundData) < 2 {
	//		return nil, errors.New("invalid sound data format")
	//	}
	//
	//	sound := model.Sound{
	//		ID:                soundData[0].(string),
	//		AuthorID:          soundData[1].(string),
	//		AudioURL:          soundData[2].(string),
	//		IconURL:           soundData[3].(string),
	//		FileName:          soundData[4].(string),
	//		CreatedAt:         soundData[5].(string),
	//		Title:             soundData[6].(string),
	//		MusicalInstrument: soundData[7].(string),
	//		Genre:             soundData[8].(string),
	//		Mood:              soundData[9].(string),
	//		Tonality:          soundData[10].(string),
	//		Tempo:             soundData[11].(string),
	//		Style:             soundData[12].(string),
	//	}
	//
	//	sounds = append(sounds, sound)
	//}

	return sounds, nil
}
