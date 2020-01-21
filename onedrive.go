package onedrive

import (
	"net/http"
	"time"
)

const (
	version = "0.1"
	baseURL = "https://api.onedrive.com/v1.0"
)

// OneDrive is the entry point for the client. It manages the communication with
// Microsoft OneDrive API
type OneDrive struct {
	Client *http.Client
	// When debug is set to true, the JSON response is formatted for better readability
	Debug   bool
	BaseURL string
	// Services
	Drives *DriveService
	Items  *ItemService
	// Private
	throttle time.Time
}

// New returns a new OneDrive client to enable you to communicate with
// the API
func New(c *http.Client) *OneDrive {
	drive := OneDrive{
		Client:   c,
		BaseURL:  baseURL,
		throttle: time.Now(),
	}
	drive.Drives = &DriveService{&drive}
	drive.Items = &ItemService{&drive}
	return &drive
}

func (od *OneDrive) throttleRequest(time time.Time) {
	od.throttle = time
}
