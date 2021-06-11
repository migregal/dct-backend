package net

import "github.com/valyala/fasthttp"

func (h *Handler) getTasks(reqStr []byte) (*ResponseBody, error) {
	nreq := fasthttp.AcquireRequest()
	nreq.SetRequestURI(h.Url)
	nreq.Header.Add("Authorization", h.Auth)
	nreq.Header.SetMethodBytes(POST)
	nreq.Header.SetContentType("application/json")
	nreq.SetBody(reqStr)

	nres := fasthttp.AcquireResponse()
	if err := fasthttp.Do(nreq, nres); err != nil {
		return nil, err
	}
	fasthttp.ReleaseRequest(nreq)

	data := &Response{}
	if err := data.UnmarshalJSON(nres.Body()); err != nil {
		return nil, err
	}
	fasthttp.ReleaseResponse(nres)

	return data.ResponseBody, nil
}
