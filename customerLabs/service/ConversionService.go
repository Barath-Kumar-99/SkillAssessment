package Service

import (
	"customerLabs/domain"
	"customerLabs/dto"
	"fmt"
	"sync"
)

type ConversionService interface {
	Worker(ch chan dto.ConversionRequestDto, ch2 chan dto.ConversionResponseDto, wg *sync.WaitGroup) *dto.ConversionResponseDto
}

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
		ch2 <- converted
	}
}

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

type DefaultConversionService struct {
	repo domain.ConversionRepository
}

func NewConversionService(repository domain.ConversionRepository) DefaultConversionService {
	return DefaultConversionService{repository}
}
