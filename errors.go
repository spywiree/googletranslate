package googletranslate

import "errors"

var NoTranslatedDataErr = errors.New("no translated data in response")
var MaxTextLengthExceededErr = errors.New("the maximum text length of 5000 characters has been exceeded")

// https://github.com/nidhaloff/deep-translator/blob/master/deep_translator/exceptions.py#L120
var TooManyRequestsErr = errors.New("too many requests. you can make up to 5 requests per second and up to 200k per day")
