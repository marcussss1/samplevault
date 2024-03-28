package sounds

import (
	"context"
	"fmt"
	"github.com/marcussss1/simplevault/internal/model"
)

func (r Repository) GetRandomSounds(ctx context.Context) ([]model.Sound, error) {
	var sounds []model.Sound

	resp, err := r.conn.Eval("return box.space.sounds:select()", []interface{}{})
	if err != nil {
		return nil, fmt.Errorf("select all sounds from tarantool storage: %w", err)
	}

	for _, tuple := range resp.Tuples() {
		for _, a := range tuple {
			//fmt.Println(a[0])
			b := a.(model.Sound)
			fmt.Println(b)
		}
		//sounds = append(sounds, model.Sound{
		//	ID:                fmt.Sprint(tuple[0]),
		//	AuthorID:          fmt.Sprint(tuple[1]),
		//	AudioURL:          fmt.Sprint(tuple[2]),
		//	IconURL:           fmt.Sprint(tuple[3]),
		//	FileName:          fmt.Sprint(tuple[4]),
		//	CreatedAt:         fmt.Sprint(tuple[5]),
		//	Title:             fmt.Sprint(tuple[6]),
		//	MusicalInstrument: fmt.Sprint(tuple[7]),
		//	Genre:             fmt.Sprint(tuple[8]),
		//	Mood:              fmt.Sprint(tuple[9]),
		//	Tonality:          fmt.Sprint(tuple[10]),
		//	Tempo:             fmt.Sprint(tuple[11]),
		//	Style:             fmt.Sprint(tuple[12]),
		//	IsGenerated:       false,
		//})
	}

	return sounds, nil
	//var sounds []model.Sound
	//
	//r.conn.Eval("", interface{}{})
	//
	//err := r.conn.SelectTyped("sounds", "primary", 0, 50,
	//	tarantool.IterAll, tarantool.StringKey{""}, &sounds,
	//)
	//if err != nil {
	//	return nil, fmt.Errorf("select all sounds from tarantool storage: %w", err)
	//}
	//
	//return sounds, nil
}
