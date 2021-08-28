package forms

import "net/url"

type Form struct {
	url.Values
	Errors errors
}

//New creates and initializes a form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string]string{}),
	}
}
