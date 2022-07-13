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
	Media []byte `url:"media,omitempty"`
}

func (s *MediaService) MediaUpload(params *MediaUploadPhotoParams) (*MediaUploadPhotoParams, *http.Response, error) {
	user := new(User)
	apiError := new(APIError)
	resp, err := s.sling.New().Base("https://upload.twitter.com/1.1/media/").Post("upload.json").BodyForm(params).Receive(user, apiError)
	return user, resp, relevantError(err, *apiError)
}
