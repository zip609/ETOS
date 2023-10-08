package tools

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Tool struct {
	ID        int
	ToolsName string
	ToolsID   int
	Info      string
}

func GetScanToolsFromDB(db *sql.DB) ([]Tool, error) {
	var tools []Tool

	rows, err := db.Query("SELECT id, toolsname, toolsid, info FROM Tools_list")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var tool Tool
		if err := rows.Scan(&tool.ID, &tool.ToolsName, &tool.ToolsID, &tool.Info); err != nil {
			return nil, err
		}
		tools = append(tools, tool)
	}

	return tools, rows.Err()
}
