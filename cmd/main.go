package main

import (
	"TestTest/infrastructure"
	"TestTest/interfaces"
	"TestTest/user_cases"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
)

func main() {

	// инициализация переменной
	connStr := "user=postgres password=ihavetoget5588 dbname=postgres sslmode=disable"
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	dbimage := infrastructure.NewDbimageConnect(db)

	libImage := interfaces.NewLibraryImages()

	repoImage := interfaces.NewRepositoryImages(dbimage)

	resizeImager := user_cases.NewService(libImage, repoImage)

	handlers := infrastructure.NewHandlers(resizeImager)

	// изменить размер картинки
	http.HandleFunc("/struct/", handlers.HandleResizeImage)

	// история по измененным картинкам
	http.HandleFunc("/historyimages/", handlers.HandleHistoryImages)

	// данные картинки по id
	http.HandleFunc("/getimage/{id}/", handlers.HandleGetImageById)

	// изменить данные картинки по id
	http.HandleFunc("/updateimage/{id}/", handlers.HandleUpdateImageById)

	//for check THEN: delete
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "form.html")
	})

	fmt.Println("Server is listening...")
	http.ListenAndServe(":45998", nil)
}
