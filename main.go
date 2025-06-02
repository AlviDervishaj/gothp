package main

import (
	"database/sql"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data any, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	e.HideBanner = true

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}

	renderer.templates = template.Must(renderer.templates.ParseGlob("htmx/*.html"))

	e.Renderer = renderer

	// Server static files
	e.Static("/static", "static")
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	dbUrl := os.Getenv("DATABASE_URL")
	// log out dbUrl for debugging

	log.Printf("Database URL: %s", dbUrl)
	if dbUrl == "" {
		dbUrl = "postgres://user:password@localhost:5432/mydb?sslmode=disable"
	}
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatalf("DB Connection error: %v", err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatalf("DB Ping Error: %v", err)
	}
	log.Println("Connected to PostgreSQL")

	e.GET("/", func(c echo.Context) error {
		data := struct {
			Title string
		}{
			Title: "Home Page",
		}
		return c.Render(http.StatusOK, "layout", data)
	})

	e.GET("/htmx/hello", func(c echo.Context) error {

		return c.Render(http.StatusOK, "hello.html", map[string]string{
			"Name": "HTMX Visitor",
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on: %s...", port)
	e.Logger.Fatal(e.Start(":" + port))
}
