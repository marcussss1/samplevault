package model

import "errors"

var (
	// ErrNotFound - сущность не найдена.
	ErrNotFound = errors.New("not found")
	// ErrAlreadyExist - сущность уже существует.
	ErrAlreadyExist = errors.New("already exist")
)
