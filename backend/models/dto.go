package models

type GenerictResult[T any] struct {
	Data         T      `json:"data"`
	ErrorMessage string `json:"errorMessage"`
}
