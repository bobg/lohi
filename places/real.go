package places

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/bobg/errors"
	"github.com/bobg/oauther/v5"
)

// RealService is an implementation of Service that talks to the Google Places API.
type RealService struct {
	client *http.Client
}

// NewRealService creates a new RealService.
// It reads the Google API credentials from creds, and the OAuth token from tokenFile.
// If there is no valid token in tokenFile, it will open the user's browser to obtain one and write it to the file.
func NewRealService(ctx context.Context, creds []byte, tokenFile string) (*RealService, error) {
	client, err := oauther.Client(ctx, tokenFile, creds, "https://www.googleapis.com/auth/cloud-platform")
	return &RealService{client: client}, errors.Wrap(err, "creating oauth client")
}

// GetPlace returns a place by its ID.
func (s *RealService) GetPlace(ctx context.Context, id string) (*Place, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", "https://places.googleapis.com/v1/places/"+id, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "creating request for place %s", id)
	}
	req.Header.Set("X-Goog-FieldMask", "*")
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "fetching place %s", id)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "reading response body for place %s", id)
	}

	var place Place
	err = json.Unmarshal(body, &place)
	return &place, errors.Wrapf(err, "decoding place %s", id)
}
