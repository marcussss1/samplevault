package sounds

import (
	"context"
	"fmt"
	"github.com/marcussss1/simplevault/internal/model"
)

func (r Repository) GetRandomSounds(ctx context.Context) ([]model.Sound, error) {

	resp, err := r.conn.Eval("return 42", []interface{}{})
	if err != nil {
		return nil, fmt.Errorf("select all sounds from tarantool storage: %w", err)
	}

	fmt.Println(resp.Data[0].([]interface{})[0].(uint64))

	return nil, nil
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
