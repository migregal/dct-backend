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
	Fromcase *struct {
		Fromcaseid string `json:"fromcaseid"`
	} `json:"fromcase,omitempty"`
	Fromloc *struct {
		Fromlocid string `json:"fromlocid"`
	} `json:"fromloc,omitempty"`
	Sku *struct {
		Skubarcode string `json:"skubarcode"`
	} `json:"sku,omitempty"`
	Tocase *struct {
		Tocaseid string `json:"tocaseid"`
	} `json:"tocase,omitempty"`
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
		Name string `json:"name,omitempty"`
		Guid string `json:"guid,omitempty"`
	} `json:"user,omitempty"`
	Operations []struct {
		OperationCode string `json:"code,omitempty"`
		OperationName string `json:"name,omitempty"`
	} `json:"operations,omitempty"`
	Task *struct {
		Taskname string `json:"taskname,omitempty"`
		Taskguid string `json:"taskguid,omitempty"`
		Totalqty string `json:"totalqty,omitempty"`
		Execqty  string `json:"execqty,omitempty"`
	} `json:"task,omitempty"`
	Tasklist []struct {
		Taskname string `json:"taskname,omitempty"`
		Taskguid string `json:"taskguid,omitempty"`
	} `json:"tasklist,omitempty"`
	Caselist []struct {
		Caseid    string `json:"caseid,omitempty"`
		Casename  string `json:"casename,omitempty"`
		Locid     string `json:"locid,omitempty"`
		Loc       string `json:"loc,omitempty"`
		Qty       int    `json:"qty,omitempty"`
		Deviation bool   `json:"deviation,omitempty"`
	} `json:"caselist,omitempty"`
	Case *struct {
		Casename string `json:"casename,omitempty"`
		Caseid   string `json:"caseid,omitempty"`
	} `json:"case,omitempty"`
	Skucaselist []struct {
		Skubarcode string `json:"skubarcode,omitempty"`
		Skuname    string `json:"skuname,omitempty"`
		RequiredKM bool   `json:"requiredKM,omitempty"`
		Qty        int    `json:"qty,omitempty"`
		Deviation  bool   `json:"deviation,omitempty"`
	} `json:"skucaselist,omitempty"`
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
