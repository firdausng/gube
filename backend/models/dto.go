package models

type GenerictResult[T any] struct {
	Data         T
	ErrorMessage string
}
