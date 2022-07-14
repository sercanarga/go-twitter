package twitter

import (
	"net/http"

	"github.com/dghubble/sling"
)

// AccountService provides a method for account credential verification.
type MediaService struct {
	sling *sling.Sling
}

// newAccountService returns a new AccountService.
func newMediaService(sling *sling.Sling) *MediaService {
	return &MediaService{
		sling: sling.Path("media/"),
	}
}

type MediaUploadPhotoParams struct {
	MediaData string `url:"media_data,omitempty"`
}

type MediaUploadResponse struct {
	MediaID  int `json:"media_id"`
	MediaIDStr  string `json:"media_id_string"`
	MediaKey  string `json:"media_key"`
}

func (s *MediaService) MediaUpload(params *MediaUploadPhotoParams) (*MediaUploadResponse, *http.Response, error) {
	media := new(MediaUploadResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Base("https://upload.twitter.com/1.1/media/").Post("upload.json").BodyForm(params).Receive(media, apiError)
	return media, resp, relevantError(err, *apiError)
}
