package models

import (
	"database/sql"
	"errors"
	"fmt"
	"gube/backend/db"
)

type Workspace struct {
	ID          int64
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Active      bool   `binding:"required"`
}

func SetDefaultWorkspace() error {
	defaultWorkspace, err := GetEventByName("default")

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			defaultWorkspace = &Workspace{
				Name:        "default",
				Description: "This is default workspace",
				Active:      true,
			}
			err = defaultWorkspace.Save()
			if err != nil {
				errorStr := fmt.Sprintf("Error saving default workspace: %s", err)
				return errors.New(errorStr)

			}
			return nil
		}
		return err
	}

	return nil
}

func (workspace Workspace) Save() error {
	query := `
	INSERT INTO workspaces(name, description, active) 
	VALUES (?,?,?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(workspace.Name, workspace.Description, workspace.Active)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	workspace.ID = id
	return err
}

func (workspace Workspace) GetAllWorkspaces() ([]Workspace, error) {
	query := "SELECT * FROM workspaces"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var workspaces []Workspace

	for rows.Next() {
		var e Workspace
		err := rows.Scan(&e.ID, &e.Name, &e.Description)

		if err != nil {
			return nil, err
		}

		workspaces = append(workspaces, e)
	}

	return workspaces, nil
}

func GetEventById(id int64) (*Workspace, error) {
	query := "SELECT * FROM workspaces WHERE id = ?"
	rows := db.DB.QueryRow(query, id)

	var workspace Workspace
	err := rows.Scan(&workspace.ID, &workspace.Name, &workspace.Description)
	if err != nil {
		return nil, err
	}
	return &workspace, nil
}

func GetEventByName(name string) (*Workspace, error) {
	query := "SELECT * FROM workspaces WHERE name = ?"
	rows := db.DB.QueryRow(query, name)

	var workspace Workspace
	err := rows.Scan(&workspace.ID, &workspace.Name, &workspace.Description, &workspace.Active)
	if err != nil {
		return nil, err
	}
	return &workspace, nil
}

func (workspace Workspace) Update() error {
	query := `
	UPDATE workspaces
	SET name=?, description=?, active=?
	WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(workspace.Name, workspace.Description)
	if err != nil {
		return err
	}

	return err
}

func (workspace Workspace) Delete() error {
	_, err := db.DB.Exec("DELETE FROM workspaces WHERE id = ?", workspace.ID)
	return err
}

func GetActiveWorkspace() (*Workspace, error) {
	query := "SELECT * FROM workspaces WHERE active = 1"
	rows := db.DB.QueryRow(query)

	var workspace Workspace
	err := rows.Scan(&workspace.ID, &workspace.Name, &workspace.Description, &workspace.Active)
	if err != nil {
		return nil, err
	}
	return &workspace, nil

	//var workspace Workspace
	//err := db.DB.QueryRow("SELECT * FROM workspaces WHERE active = 1").Scan(&workspace.ID, &workspace.Name, &workspace.Description)
	//if err != nil {
	//	return nil, err
	//}
	//return &workspace, nil
}
