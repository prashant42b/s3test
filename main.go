package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/prashant42b/jqueryApp/s3"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func main() {

	newFiberApp()
	s3.InitAWSSession()
	s3.FetchS3BucketContent()
	renderJqueryContent()

}

func renderJqueryContent() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", indexHandler)
	log.Print("Server is listening on port 8080....")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func newFiberApp() {

	app := fiber.New()
	// Use the log package in Fiber middleware
	app.Use(func(c *fiber.Ctx) error {
		// Log the request method and path
		log.Printf("Request received - Method: %s, Path: %s", c.Method(), c.Path())

		// Continue to the next middleware or route handler
		return c.Next()
	})

	//router.SetupRoutes(app)

	//Listens on port 3000
	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
