package app

import (
	"customerLabs/dto"
	"customerLabs/errs"
	Service "customerLabs/service"
	"encoding/json"
	"fmt"
	"net/http"
)
//struct to connect with service
type CHandlers struct {
	CService Service.ConversionService
}

// this function used to validate the request body and return the response from channel
func (h *CHandlers) ConversionHandler(ch chan dto.ConversionRequestDto, ch2 chan dto.ConversionResponseDto, w http.ResponseWriter, r *http.Request) {

	var requestBody dto.ConversionRequestDto

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)

	//Added error handling to send the error response from json decoder
	if err != nil {
		errormessage := errs.ErrorResponse{
			Errors: struct {
				Code    string `json:"code"`
				Message string `json:"message"`
			}{
				Code:    "CONV001",
				Message: err.Error(),
			},
		}
		writeResponse(w, http.StatusBadRequest, errormessage)
	} else {
		// Send the request to the channel for processing by a worker in a separate goroutine
		go func() {
			ch <- requestBody
		}()

		resBody, ok := <-ch2
		if !ok {
			fmt.Println("Channel closed. Exiting worker.")
			return
		}
		writeResponse(w, http.StatusOK, resBody)
	}

}

// Added a seperate writeResponse function for write the response
func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
