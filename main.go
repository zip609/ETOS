package main

import (
	"ETOS/tools"
	"database/sql"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB // ваша глобальная переменная для подключения к базе данных

type TemplateRenderer struct {
	templates *template.Template
}

type FormData struct {
	IPOrDomain string `form:"ip_or_domain"`
	Tool       string `form:"tool"`
	Profile    string `form:"profile"`
	Command    string `form:"command"`
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	db, err := initDB("etos.db")
	if err != nil {
		e.Logger.Fatal(err)
	}
	defer db.Close()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to the ETOS Brute Web Service!")
	})

	e.Renderer = &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}

	e.GET("/login", func(c echo.Context) error {
		return c.Render(http.StatusOK, "login.html", nil)
	})

	e.POST("/login", HandleLogin)

	e.GET("/tools", HandleGetTools)

	e.GET("/scan", func(c echo.Context) error {

		scantools, err := tools.GetScanToolsFromDB(db)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Error retrieving list of scan tools")
		}

		profiles, err := tools.GetNmapProfilesFromDB(db)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Error retrieving profiles")
		}

		nmaphelp, err := tools.GetNmapHelpFromDB(db)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Error retrieving Nmap help")
		}

		return c.Render(http.StatusOK, "scan.html", map[string]interface{}{
			"scantools": scantools,
			"profiles":  profiles,
			"nmaphelp":  nmaphelp, // передайте это значение в шаблон
		})
	})

	e.GET("/getNmapString", tools.GetNmapString(db))

	e.POST("/updateNmapString", tools.UpdateNmapString(db))

	e.Logger.Fatal(e.Start(":8080"))
}

func initDB(filepath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		return nil, err
	}
	if db == nil {
		return nil, fmt.Errorf("DB nil")
	}
	return db, nil
}

func CreateJwtToken(username string) (string, error) {
	claims := &jwt.StandardClaims{
		Subject:   username,
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := rawToken.SignedString([]byte("AHTFGSD45664hlds6576"))
	if err != nil {
		return "", err
	}

	return token, nil
}

func HandleLogin(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	db, err := sql.Open("sqlite3", "./etos.db")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Database connection error")
	}
	defer db.Close()

	stmt := `SELECT username FROM Users WHERE username = ? AND password = ?`
	row := db.QueryRow(stmt, username, password)

	var retrievedUsername string
	err = row.Scan(&retrievedUsername)

	if err != nil {
		if err == sql.ErrNoRows {
			return echo.ErrUnauthorized
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Database query error")
	}
	token, err := CreateJwtToken(username)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error creating JWT token")
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}

func HandleGetTools(c echo.Context) error {
	tools, err := tools.GetScanToolsFromDB(db)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error retrieving tools list"})
	}
	return c.JSON(http.StatusOK, tools)
}
