package identity_client

import (
	"fmt"
	"strings"
)

// AddTokenPrefix first checks to see if a `Bearer` prefix exists &
// if it doesn't then will add one.
func AddTokenPrefix(t string) string {
	tokenSlices := strings.Split(t, "Bearer ")
	if len(tokenSlices) == 1 {
		return fmt.Sprintf("Bearer %s", t)
	}
	return t
}
