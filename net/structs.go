package net

//easyjson:json
type Header struct {
	Method string `json:"method"`
	Token  string `json:"token"`
}

//easyjson:json
type RequestHeader struct {
	Header *Header `json:"header,omitempty"`
}

//easyjson:json
type Body struct {
	Barcode *string `json:"barcode,omitempty"`
	User    *struct {
		Guid *string `json:"guid,omitempty"`
	} `json:"user,omitempty"`
	OperationCode *string `json:"operationcode,omitempty"`
}

//easyjson:json
type RequestBody struct {
	Body *Body `json:"request,omitempty"`
}

type Request struct {
	Header *Header `json:"header,omitempty"`
	Body   *Body   `json:"request,omitempty"`
}

//easyjson:json
type Data struct {
	User *struct {
		Name string `json:"name"`
		Guid string `json:"guid"`
	} `json:"user,omitempty"`
	OperationList []struct {
		OperationCode string `json:"operationcode"`
		OperationName string `json:"operationname"`
	} `json:"operationlist,omitempty"`
	Tasklist []struct {
		Taskname string `json:"taskname"`
		Taskguid string `json:"taskguid"`
	} `json:"tasklist,omitempty"`
}

//easyjson:json
type ResponseBody struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Error   bool   `json:"error"`
	Data    *Data  `json:"data,omitempty"`
}

//easyjson:json
type Response struct {
	Header       *Header       `json:"header"`
	RequestBody  *Body         `json:"request"`
	ResponseBody *ResponseBody `json:"response"`
}
