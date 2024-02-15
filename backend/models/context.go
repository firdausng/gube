package models

type Context struct {
	ID          int64
	Name        string `binding:"required"`
	DisplayName string `binding:"required"`
	Active      bool   `binding:"required"`
}
