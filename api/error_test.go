package api

import (
	"encoding/json"
	"testing"

	"github.com/nordsieck/defect"
)

const exampleError = `{
	"error": {
		"errors": [
			{
				"domain": "usageLimits",
				"reason": "dailyLimitExceededUnreg",
				"message": "Daily Limit for Unauthenticated Use Exceeded. Continued use requires signup.",
				"extendedHelp": "https://code.google.com/apis/console"
			}
		],
		"code": 403,
		"message": "Daily Limit for Unauthenticated Use Exceeded. Continued use requires signup."
	}
}`

func TestErrorUnmarshal(t *testing.T) {
	var container ErrorContainer
	err := json.Unmarshal([]byte(exampleError), &container)
	defect.Equal(t, err, nil)
	defect.DeepEqual(t, container, ErrorContainer{
		Error: Error{
			Errors: []ErrorDetail{{
				Domain:       "usageLimits",
				Reason:       "dailyLimitExceededUnreg",
				Message:      "Daily Limit for Unauthenticated Use Exceeded. Continued use requires signup.",
				ExtendedHelp: "https://code.google.com/apis/console",
			}},
			Code:    403,
			Message: "Daily Limit for Unauthenticated Use Exceeded. Continued use requires signup.",
		},
	})
}
