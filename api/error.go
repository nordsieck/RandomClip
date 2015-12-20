package api

type ErrorContainer struct {
	Error `json:"error"`
}

type Error struct {
	Errors  []ErrorDetail `json:"errors"`
	Code    uint16        `json:"code"`
	Message string        `json:"message"`
}

func (e *Error) IsZero() bool { return len(e.Errors) == 0 && e.Code == 0 && e.Message == "" }

type ErrorDetail struct {
	Domain       string `json:"domain"`
	Reason       string `json:"reason"`
	Message      string `json:"message"`
	ExtendedHelp string `json:"extendedHelp"`
}
