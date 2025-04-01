package googletranslate

import (
	"errors"
	"strconv"
)

type HttpError int

func (e HttpError) Error() string {
	return "received non-OK HTTP status: " + strconv.Itoa(int(e))
}

//nolint:staticcheck
var (
	NoTranslatedDataErr      = errors.New("no translated data in response")
	MaxTextLengthExceededErr = errors.New("the maximum text length of 5000 characters has been exceeded")
	// https://github.com/nidhaloff/deep-translator/blob/master/deep_translator/exceptions.py#L120
	TooManyRequestsErr = errors.New("too many requests. you can make up to 5 requests per second and up to 200k per day")
)
