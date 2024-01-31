package main

import (
	"customerLabs/app"
	"customerLabs/domain"
	"customerLabs/dto"
	Service "customerLabs/service"
	"fmt"
	"log"
	"net/http"
	"sync"
)

func main() {

	requestChannel := make(chan dto.ConversionRequestDto)
	responseChannel := make(chan dto.ConversionResponseDto)
	// Use a WaitGroup to wait for all workers to finish before exiting
	var wg sync.WaitGroup

	log.Println("Starting the application...")
	//router := mux.NewRouter()

	conversionRepository := domain.NewConversionRepositoryDb()

	HC := app.CHandlers{CService: Service.NewConversionService(conversionRepository)}

	// Start multiple worker goroutines (adjust the number as needed)

	numWorkers := 5
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go HC.CService.Worker(requestChannel, responseChannel, &wg)
	}

	http.HandleFunc("/averagekootam/conversion", func(w http.ResponseWriter, r *http.Request) {
		HC.ConversionHandler(requestChannel, responseChannel, w, r)
	})

	port := 8080
	fmt.Printf("Server is listening on port %d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}

	// Close the channel and wait for all workers to finish
	close(requestChannel)
	close(responseChannel)
	wg.Wait()
}
