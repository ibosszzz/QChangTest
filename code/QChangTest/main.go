package main

import (
	"net/http"
	"fmt"
	"log"
	"time"
	"github.com/gorilla/mux"
	"QChangTest/controller"
)

func main(){
	r := mux.NewRouter()
	r.HandleFunc("/findValue", controller.FindValue).Methods("GET")
	r.HandleFunc("/cashier", controller.Cashier).Methods("POST")
	r.HandleFunc("/checkCashierDesk", controller.CheckCashierDesk).Methods("GET")
	r.HandleFunc("/addCash", controller.AddCash).Methods("POST")
	

	srv := &http.Server{
        Handler:      r,
        Addr:         "127.0.0.1:8900",
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }
	fmt.Println("Server Starting...")
	log.Fatal(srv.ListenAndServe())
}
