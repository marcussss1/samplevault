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

	// Создаем переменную для распакованных данных
	//var unpackedData map[string]interface{}
	//
	//err = msgpack.Unmarshal([]byte(resp.String()), &unpackedData)
	//if err != nil {
	//	return nil, err
	//}

	fmt.Println(resp.Tuples())

	for _, tuple := range resp.Tuples() {
		fmt.Println()
		fmt.Println(tuple)
		fmt.Println()
	}

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
