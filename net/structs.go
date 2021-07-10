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
	Task          *struct {
		Taskguid string `json:"taskguid"`
	} `json:"task,omitempty"`
	Case *struct {
		Caseid string `json:"caseid"`
	} `json:"case,omitempty"`
	Previouscase *struct {
		Previouscaseid string `json:"previouscaseid"`
	} `json:"previouscase,omitempty"`
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
	Operations []struct {
		OperationCode string `json:"code"`
		OperationName string `json:"name"`
	} `json:"operations,omitempty"`
	Task *struct {
		Taskname string `json:"taskname"`
		Taskguid string `json:"taskguid"`
		Totalqty string `json:"totalqty"`
		Execqty  string `json:"execqty"`
	} `json:"task,omitempty"`
	Tasklist []struct {
		Taskname string `json:"taskname"`
		Taskguid string `json:"taskguid"`
	} `json:"tasklist,omitempty"`
	Caselist []struct {
		Caseid    string `json:"caseid"`
		Casename  string `json:"casename"`
		Locid     string `json:"locid"`
		Loc       string `json:"loc"`
		Qty       int    `json:"qty"`
		Deviation bool   `json:"deviation"`
	} `json:"caselist,omitempty"`
	Case *struct {
		Casename string `json:"casename"`
		Caseid   string `json:"caseid"`
	} `json:"case,omitempty"`
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
