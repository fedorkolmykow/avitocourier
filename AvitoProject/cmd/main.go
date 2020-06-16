package main

import (
	"fmt"
	"net/http"

	"AvitoProject/pkg/calculator"
	"AvitoProject/pkg/clientdb"
	"AvitoProject/pkg/limiter"
	"AvitoProject/pkg/orderservice"
	"AvitoProject/pkg/orderservice/httpserver"
)


func main() {
	dbCon := clientdb.NewDb()
    calc := calculator.NewCalculator()
    service := orderservice.NewService(calc, dbCon)
    server := httpserver.NewServer(service)
    lim := limiter.NewCustomLimiter(10)

	fmt.Println("starting server at :8080")
	err := http.ListenAndServe(":8080", lim.Limit(server))
	if err != nil{
		fmt.Println(err)
	}
}