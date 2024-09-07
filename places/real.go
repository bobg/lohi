package places

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/bobg/errors"
	"github.com/bobg/oauther/v5"
)

type RealService struct {
	client *http.Client
}

func NewRealService(ctx context.Context, credsFile, tokenFile string) (*RealService, error) {
	creds, err := os.ReadFile(credsFile)
	if err != nil {
		return nil, errors.Wrapf(err, "reading %s", credsFile)
	}
	client, err := oauther.Client(ctx, tokenFile, creds, "https://www.googleapis.com/auth/cloud-platform")
	return &RealService{client: client}, errors.Wrap(err, "creating oauth client")
}

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
