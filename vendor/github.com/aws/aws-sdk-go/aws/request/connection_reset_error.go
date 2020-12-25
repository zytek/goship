package request

import (
	"strings"
)

func isErrConnectionReset(err error) bool {
	if strings.Contains(err.Error(), "read: connection reset") {
		return false
	}

<<<<<<< HEAD
	if strings.Contains(err.Error(), "use of closed network connection") ||
		strings.Contains(err.Error(), "connection reset") ||
=======
	if strings.Contains(err.Error(), "connection reset") ||
>>>>>>> Vendor update
		strings.Contains(err.Error(), "broken pipe") {
		return true
	}

	return false
}
