package reqid

import (
	"context"
	"crypto/rand"
	"fmt"

	"github.com/oklog/ulid/v2"
)

type ReqID string

const RequestIDKey ReqID = "request_id"

func NewRequestID() (ReqID, error) {
	id, err := ulid.New(ulid.Now(), rand.Reader)
	if err != nil {
		return "", fmt.Errorf("generate ulid: %w", err)
	}
	return ReqID(id.String()), nil
}

func RequestID(c context.Context) (string, bool) {
	id, ok := c.Value(RequestIDKey).(ReqID)
	return string(id), ok
}
