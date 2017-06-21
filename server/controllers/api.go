package controllers

// JsonResponse is the universal API response struct.
// The response will always respond with an ok field,
// but the other fields are optional.
type JsonResponse struct {
	Ok       bool        `json:"ok"`
	Username string      `json:"username,omitempty"`
	Message  string      `json:"message,omitempty"`
	Data     interface{} `json:"data,omitempty"`
}
