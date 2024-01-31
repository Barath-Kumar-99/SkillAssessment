package Service

import (
	"customerLabs/domain"
	"customerLabs/dto"
	"fmt"
	"sync"
)

//initialization of interface for service layer
type ConversionService interface {
	Worker(ch chan dto.ConversionRequestDto, ch2 chan dto.ConversionResponseDto, wg *sync.WaitGroup) *dto.ConversionResponseDto
}
//worker function to convert the request to ConversionRequestDto type
func (r DefaultConversionService) Worker(ch chan dto.ConversionRequestDto, ch2 chan dto.ConversionResponseDto, wg *sync.WaitGroup) *dto.ConversionResponseDto {
	defer wg.Done()

	for {
		// Receive request from channel
		requestBody, ok := <-ch
		if !ok {
			fmt.Println("Channel closed. Exiting worker.")
			return nil
		}
		// Convert the request to the new format
		converted := convertToNewFormat(requestBody)
		// Send response back to channel
		ch2 <- converted
	}
}
//convertToNewFormat to convert the request to dto.ConversionResponseDto  format
func convertToNewFormat(req dto.ConversionRequestDto) dto.ConversionResponseDto {

	// Attribute represents an attribute in the converted format
	type Attribute struct {
		Value string `json:"value"`
		Type  string `json:"type"`
	}

	// Trait represents a trait in the converted format
	type Trait struct {
		Value string `json:"value"`
		Type  string `json:"type"`
	}

	//converting the request to dto.ConversionResponseDto format
	converted := dto.ConversionResponseDto{
		Event:           req["ev"].(string),
		EventType:       req["et"].(string),
		AppID:           req["id"].(string),
		UserID:          req["uid"].(string),
		MessageID:       req["mid"].(string),
		PageTitle:       req["t"].(string),
		PageURL:         req["p"].(string),
		BrowserLanguage: req["l"].(string),
		ScreenSize:      req["sc"].(string),
		Attributes:      make(map[string]dto.Attribute),
		Traits:          make(map[string]dto.Trait),
	}
	//added the for loop to iterate the req and assign the values of attributes and traits
	for key, value := range req {
		if len(key) >= 4 && key[:4] == "atrk" {
			attrNum := key[4:]
			attrValueKey := "atrv" + attrNum
			attrTypeKey := "atrt" + attrNum

			converted.Attributes[value.(string)] = dto.Attribute{
				Value: req[attrValueKey].(string),
				Type:  req[attrTypeKey].(string),
			}
		} else if len(key) >= 5 && key[:5] == "uatrk" {
			traitNum := key[5:]
			traitValueKey := "uatrv" + traitNum
			traitTypeKey := "uatrt" + traitNum

			converted.Traits[value.(string)] = dto.Trait{
				Value: req[traitValueKey].(string),
				Type:  req[traitTypeKey].(string),
			}
		}
	}

	return converted
}

//establishing the connection between service and repository incase of db connection
type DefaultConversionService struct {
	repo domain.ConversionRepository
}


func NewConversionService(repository domain.ConversionRepository) DefaultConversionService {
	return DefaultConversionService{repository}
}
