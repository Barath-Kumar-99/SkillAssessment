package dto

type ConversionResponseDto struct {
	Event           string               `json:"event"`
	EventType       string               `json:"event_type"`
	AppID           string               `json:"app_id"`
	UserID          string               `json:"user_id"`
	MessageID       string               `json:"message_id"`
	PageTitle       string               `json:"page_title"`
	PageURL         string               `json:"page_url"`
	BrowserLanguage string               `json:"browser_language"`
	ScreenSize      string               `json:"screen_size"`
	Attributes      map[string]Attribute `json:"attributes"`
	Traits          map[string]Trait     `json:"traits"`
}

type Attribute struct {
	Value string `json:"value"`
	Type  string `json:"type"`
}

type Trait struct {
	Value string `json:"value"`
	Type  string `json:"type"`
}

//type ConversionResponseDto struct {
//	Event           string `json:"event"`
//	EventType       string `json:"event_type"`
//	AppId           string `json:"app_id"`
//	UserId          string `json:"user_id"`
//	MessageId       string `json:"message_id"`
//	PageTitle       string `json:"page_title"`
//	PageUrl         string `json:"page_url"`
//	BrowserLanguage string `json:"browser_language"`
//	ScreenSize      string `json:"screen_size"`
//	Attributes      struct {
//		FormVarient struct {
//			Value string `json:"value"`
//			Type  string `json:"type"`
//		} `json:"form_varient"`
//		Ref struct {
//			Value string `json:"value"`
//			Type  string `json:"type"`
//		} `json:"ref"`
//	} `json:"attributes"`
//	Traits struct {
//		Age struct {
//			Value string `json:"value"`
//			Type  string `json:"type"`
//		} `json:"age"`
//		Email struct {
//			Value string `json:"value"`
//			Type  string `json:"type"`
//		} `json:"email"`
//		Name struct {
//			Value string `json:"value"`
//			Type  string `json:"type"`
//		} `json:"name"`
//	} `json:"traits"`
//}
