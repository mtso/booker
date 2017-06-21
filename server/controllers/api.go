package controllers

// JsonResponse is the universal API response struct.
// The response will always respond with an ok field,
// but the other fields are optional.
type JsonResponse map[string]interface{}

func (r *JsonResponse) Set(key string, value interface{}) {
	(*r)[key] = value
}
