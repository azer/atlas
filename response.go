package atlas

import "encoding/json"

type Response struct {
	Code   int
	Result interface{}
	Output []byte
	IsJSON bool
}

type SuccessResponseContext struct {
	Result interface{} `json:"result"`
	Ok     bool        `json:"ok"`
}

type ErrorResponseContext struct {
	Error interface{} `json:"error"`
}

func (response *Response) JSON() (string, error) {
	parsed, err := json.MarshalIndent(response.Result, "", "	")

	if err != nil {
		return "", err
	}

	return string(parsed), nil
}
