package storage

import (
	"bytes"
	"context"
)

type Uploader interface {
	Upload(c context.Context, buff bytes.Buffer) (string, error)
}
