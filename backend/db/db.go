package db

import (
	"database/sql"
	"fmt"
	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	db, err := sql.Open("sqlite", "gube.sql")
	if err != nil {
		panic("Database could not connect: " + err.Error())
	}

	DB = db

	err = createTable()
	if err != nil {
		panic("Database could not connect: " + err.Error())
	}

	fmt.Println("Tables created successfully!")
}

func createTable() error {
	createWorkspacesTable := `
        CREATE TABLE IF NOT EXISTS workspaces (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT NOT NULL,
            description TEXT NOT NULL,            
            active INTEGER NOT NULL CHECK(active IN (0, 1))            
        )
    `
	_, err := DB.Exec(createWorkspacesTable)

	createContextTable := `
        CREATE TABLE IF NOT EXISTS contexts (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT NOT NULL,
            DisplayName TEXT NOT NULL,
            workspace_id INTEGER,
            FOREIGN KEY(workspace_id) REFERENCES workspaces(id)                              
        )
    `
	_, err = DB.Exec(createContextTable)
	return err
}
