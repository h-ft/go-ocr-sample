package functions

import (
	"context"

	"github.com/otiai10/gosseract"
)

type Functions struct {
	Gosseract *gosseract.Client
}

func New(gosseract *gosseract.Client) *Functions {
	return &Functions{
		Gosseract: gosseract,
	}
}

func (f *Functions) ReadImage(ctx context.Context) error {

	return nil
}
