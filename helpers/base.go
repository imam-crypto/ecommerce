package helpers

import "net/http"

var RESPONSE_AUTH = ConvDefaultResponse(http.StatusUnauthorized, false, "Failed", "unauthorized")
var RESPONSE_NOT_FOUND = ConvDefaultResponse(http.StatusNotFound, false, "Failed", "failed")
var RESPONSE_BAD_REQUEST = ConvDefaultResponse(http.StatusBadRequest, false, "Failed", "bad request")
var RESPONSE_INTERNAL_SERVER_ERROR = ConvDefaultResponse(http.StatusInternalServerError, false, "Failed", "internal server error")
