package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/lib/pq"
)

type Service struct {
	db *sqlx.DB
}

func main() {

	fmt.Println("Привет?")

	db, err := sqlx.Connect("postgres", "host=localhost user=postgres password=postgres dbname=coursework sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalln(err)
	}

	service := Service{
		db: db,
	}

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.POST("/create", service.create)
	e.GET("/all", service.all)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

type createRequest struct {
	Title     string `json:"title"`
	Text      string `json:"text"`
	CreatorID int    `json:"creator_id"`
}

func (s *Service) create(c echo.Context) error {
	var req createRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	fmt.Printf("%#v\n", req)

	_, err := s.db.Exec("INSERT INTO posts (title, text, creator_id) VALUES ($1, $2, $3)", req.Title, req.Text, req.CreatorID)
	if err != nil {
		return err
	}

	return nil
}

type Request struct {
	ID        int
	CreatedAt time.Time `db:"created_at"`
	CreatorID int       `db:"creator_id"`
	Title     string
	Text      string
}

func (s *Service) all(c echo.Context) error {
	requests := []Request{}
	err := s.db.Select(&requests, "SELECT * FROM posts")
	if err != nil {
		return err
	}

	return c.JSON(200, requests)
}
