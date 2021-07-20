package main

import (
	"fmt"
	"net/http"

	"github.com/Sairan-ds/golang-rest-api-course/internal/comment"
	"github.com/Sairan-ds/golang-rest-api-course/internal/database"
	transportHttp "github.com/Sairan-ds/golang-rest-api-course/internal/transport/http"
)

// App - the struct which contains things like pointers
// to database connections
type App struct{}

// Run - sets up our application
func (app *App) Run() error {
	fmt.Println("Setting Up Our APP")

	var err error
	db, err := database.NewDatabase()
	if err != nil {
		return err
	}

	err = database.MigrateDb(db)
	if err != nil {
		return err
	}

	commentService := comment.NewService(db)

	handler := transportHttp.NewHandler(commentService)
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}

	return nil
}

func main() {
	fmt.Println("GO REST API COURSE")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up our REST API")
		fmt.Println(err)
	}
}
