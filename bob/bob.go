// Bob exercism.io golang solution
package bob

import (
	"strings"
)

// Bob's brain
func Hey(remark string) string {
	// trim whitespace
	remark = strings.TrimSpace(remark)

	upperRemark := strings.ToUpper(remark)
	lowerRemark := strings.ToLower(remark)
	lastIndex := len(remark) - 1

	if lastIndex < 0 {
		return "Fine. Be that way!"
	}

	// if all caps and contains alphas
	if upperRemark == remark && lowerRemark != remark {
		if !strings.Contains(remark, "?") {
			return "Whoa, chill out!"
		} else {
			return "Calm down, I know what I'm doing!"
		}
	// Ending with a question mark
	} else if string(remark[lastIndex]) == "?" {
		return "Sure."
	}

	return "Whatever."
}
