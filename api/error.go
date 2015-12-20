package api

import "strconv"

type ErrorContainer struct {
	Error `json:"error"`
}

type Error struct {
	Errors  []ErrorDetail `json:"errors"`
	Code    uint16        `json:"code"`
	Message string        `json:"message"`
}

func (e *Error) IsZero() bool { return len(e.Errors) == 0 && e.Code == 0 && e.Message == "" }

func (e *Error) Error() string { return e.String() }
func (e *Error) String() string {
	s := ""
	for _, ed := range e.Errors {
		s += ed.String() + "\n"
	}
	return s +
		"Code: " + strconv.FormatUint(uint64(e.Code), 10) + "\n" +
		"Message: " + e.Message
}

type ErrorDetail struct {
	Domain       string `json:"domain"`
	Reason       string `json:"reason"`
	Message      string `json:"message"`
	ExtendedHelp string `json:"extendedHelp"`
}

func (e *ErrorDetail) String() string {
	return "Domain: " + e.Domain + "\n" +
		"Reason: " + e.Reason + "\n" +
		"Message: " + e.Message + "\n" +
		"ExtendedHelp: " + e.ExtendedHelp
}
