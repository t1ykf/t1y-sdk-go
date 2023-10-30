package t1y

import (
	"strings"

	"github.com/google/uuid"
)

func UUID() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}
