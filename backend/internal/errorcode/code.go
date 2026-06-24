package errorcode

const (
	Success                = 0
	MediaNotFound          = 10001
	MediaTypeInvalid       = 10002
	MediaAlreadyExists     = 10003
	FavoriteAlreadyExists  = 20001
	FavoriteNotFound       = 20002
	RatingInvalid          = 20003
	RatingNotFound         = 20004
	MediaAlreadyHidden     = 20005
	HiddenNotFound         = 20006
	SearchKeywordEmpty     = 30001
	ValidationError        = 40001
	ParameterInvalid       = 40002
	NotFound               = 40003
	ScanInProgress         = 50001
	ScanNotRunning         = 50002
	InternalError          = 90001
)

var messages = map[int]string{
	Success:               "success",
	MediaNotFound:         "media not found",
	MediaTypeInvalid:      "invalid media type",
	MediaAlreadyExists:    "media hash already exists",
	FavoriteAlreadyExists: "already favorited",
	FavoriteNotFound:      "favorite not found",
	RatingInvalid:         "rating out of range",
	RatingNotFound:        "rating not found",
	MediaAlreadyHidden:    "already hidden",
	HiddenNotFound:        "hidden record not found",
	SearchKeywordEmpty:    "search keyword cannot be empty",
	ValidationError:       "validation error",
	ParameterInvalid:      "invalid parameter",
	NotFound:              "resource not found",
	ScanInProgress:        "scan is already in progress",
	ScanNotRunning:        "no scan task is running",
	InternalError:         "internal server error",
}

func Message(code int) string {
	if msg, ok := messages[code]; ok {
		return msg
	}
	return "unknown error"
}
