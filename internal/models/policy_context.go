package models


type PolicyContext struct {
    Method  string
    Path    string
    Headers map[string]string
    Client  string // api-key / client-id
}

func NewPolicyContext(method, path string, headers map[string]string, client string) *PolicyContext {
	return &PolicyContext{
		Method:  method,
		Path:    path,
		Headers: headers,
		Client:  client,
	}
}