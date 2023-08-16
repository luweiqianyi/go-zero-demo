// Code generated by goctl. DO NOT EDIT.
package types

type Request struct {
	Name string `path:"name"`
}

type Response struct {
	Message string `json:"message"`
}

type CommonResp struct {
	Result  string `json:"result"`
	Message string `json:"message,omitempty"` // omitempty: allow to omit
}

type RegisterReq struct {
	AccountName string `json:"accountName"`
	Password    string `json:"password"`
}

type RegisterResp struct {
	CommonResp
}
