package dto

type ConvertedFormat struct {
	Event       string               `json:"event"`
	EventType   string               `json:"event_type"`
	AppID       string               `json:"app_id"`
	UserID      string               `json:"user_id"`
	MessageID   string               `json:"message_id"`
	PageTitle   string               `json:"page_title"`
	PageURL     string               `json:"page_url"`
	BrowserLang string               `json:"browser_language"`
	ScreenSize  string               `json:"screen_size"`
	Attributes  map[string]Attribute `json:"attributes"`
	UserTraits  map[string]Trait     `json:"traits"`
}
