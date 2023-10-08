package tools

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

func GetNmapProfilesFromDB(db *sql.DB) ([]string, error) {
	rows, err := db.Query("SELECT nmapprofile FROM Nmap")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var profiles []string

	for rows.Next() {
		var nmapprofile string
		if err := rows.Scan(&nmapprofile); err != nil {
			return nil, err
		}
		profiles = append(profiles, nmapprofile)
	}

	return profiles, nil
}

func GetNmapString(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		profile := c.QueryParam("profile")
		if profile == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Profile is required",
			})
		}

		var (
			id          int
			nmapprofile string
			nmapstring  string
			nmaphelp    string
		)
		err := db.QueryRow("SELECT id, nmapprofile, nmapstring, nmaphelp FROM Nmap WHERE nmapprofile = ?", profile).Scan(&id, &nmapprofile, &nmapstring, &nmaphelp)
		if err != nil {
			if err == sql.ErrNoRows {
				return c.JSON(http.StatusNotFound, map[string]string{
					"error": "Profile not found",
				})
			}
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Error retrieving nmap data",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"id":          id,
			"nmapprofile": nmapprofile,
			"nmapstring":  nmapstring,
			"nmaphelp":    nmaphelp,
		})
	}
}

func UpdateNmapString(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var requestData map[string]interface{}
		err := c.Bind(&requestData)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Bad request data",
			})
		}

		id, ok := requestData["id"].(float64) // Конвертируем из interface{} в float64, затем в int
		if !ok {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "ID is required",
			})
		}

		nmapprofile, _ := requestData["nmapprofile"].(string)
		nmapString, _ := requestData["nmapstring"].(string)
		nmaphelp, _ := requestData["nmaphelp"].(string)

		_, err = db.Exec("UPDATE Nmap SET nmapprofile=?, nmapstring=?, nmaphelp=? WHERE id=?", nmapprofile, nmapString, nmaphelp, int(id))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to update nmap data",
			})
		}

		return c.JSON(http.StatusOK, map[string]string{
			"message": "Successfully updated",
		})
	}
}

func GetNmapHelpFromDB(db *sql.DB) (string, error) {
	var nmaphelp string
	err := db.QueryRow("SELECT nmaphelp FROM Nmap LIMIT 1").Scan(&nmaphelp) // Предположим, что у вас только одна запись с помощью. Если нет, возможно, вам потребуется дополнительная логика.
	if err != nil {
		return "", err
	}
	return nmaphelp, nil
}
