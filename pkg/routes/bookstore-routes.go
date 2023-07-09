package routes

import (
	"github.com/gorilla/mux"
	"github.com/frankie-frank-frank/go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router){
	var bookWithId = "/book/{bookId}"

	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")	
	router.HandleFunc(bookWithId, controllers.GetBookById).Methods("GET")	
	router.HandleFunc(bookWithId, controllers.UpdateBook).Methods("PUT")	
	router.HandleFunc(bookWithId, controllers.DeleteBook).Methods("DELETE")	
}