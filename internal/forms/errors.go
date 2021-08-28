package forms

type errors map[string][]string

//Add adds new error related to a specific field in the form
func (e errors) Add(field, errorMessage string) {
	e[field] = append(e[field], errorMessage)
}

//Get returns the first error message of a specific field in the form
func (e errors) Get(field string) string {
	errorField := e[field]
	if len(errorField) == 0 {
		return ""
	}
	return errorField[0]
}
