package response

type ErrorName struct {
	Error ErrorContent `json:"error"`
}

type ErrorContent struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type DebugName struct {
	Debug string `json:"debug"`
}

type MessageName struct {
	Message string `json:"message"`
}
