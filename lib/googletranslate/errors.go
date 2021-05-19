package googletranslate

import "errors"

var errBadNetwork = errors.New("bad network, please check your internet connection")
var errBadRequest = errors.New("bad request, request on google Translate api isn't working")
