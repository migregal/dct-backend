package net

const (
	wrongJSONFormat = "Wrong JSON format"
)

const (
	statusFail    = "Fail"
	statusSuccess = "Success"
)

const (
	incorrectToken  = "Incorrect token"
	incorrectMethod = "Incorrect method"
)

func (h *Handler) msgSuccessResponse(msg string) *ResponseBody {
	return &ResponseBody{
		Status:  statusSuccess,
		Message: msg,
		Error:   false,
	}
}

func (h *Handler) msgErrorResponse(msg string) *ResponseBody {
	return &ResponseBody{
		Status:  statusFail,
		Message: msg,
		Error:   true,
	}
}

func (h *Handler) processRequest(request []byte) []byte {
	h.Log.Infoln(string(request))

	response := Response{}

	req := &Request{}
	if err := req.UnmarshalJSON(request); err != nil {
		h.Log.Error(err)
		response.ResponseBody = h.msgErrorResponse(wrongJSONFormat)
		str, _ := response.MarshalJSON()
		return str
	}
	response.Header = req.Header
	response.RequestBody = req.Body

	if req.Header.Token != h.AccessToken {
		h.Log.Error(incorrectToken)
		response.ResponseBody = h.msgErrorResponse(wrongJSONFormat)
		str, _ := response.MarshalJSON()
		return str
	}

	tempToken := req.Header.Token
	req.Header.Token = h.RedirectToken
	reqStr, _ := req.MarshalJSON()

	switch req.Header.Method {
	case login:
		body, err := h.login(reqStr)

		if err != nil {
			h.Log.Error(err)
			response.ResponseBody = h.msgErrorResponse(err.Error())
		} else {
			response.ResponseBody = body
		}
	case tasks:
		body, err := h.getTasks(reqStr)

		if err != nil {
			h.Log.Error(err)
			response.ResponseBody = h.msgErrorResponse(err.Error())
		} else {
			response.ResponseBody = body
		}
	default:
		h.Log.Error(req.Header.Method + " " + incorrectMethod)
		response.ResponseBody = h.msgErrorResponse(wrongJSONFormat)
	}

	response.Header.Token = tempToken

	str, _ := response.MarshalJSON()
	return str
}
