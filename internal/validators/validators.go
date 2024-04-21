package validators

import "fmt"

func ErrorMessage(field string, tag string, param string) string {
	if tag == "min" {
		return fmt.Sprintf("Invalid \"%s\" must be a minimum of %s characters.\n", tag, param)
	} else if tag == "max" {
		return fmt.Sprintf("Invalid \"%s\" must be a maximum of %s characters.\n", tag, param)
	} else if tag == "email" {
		return fmt.Sprintf("\"%s\" must be a valid email address", field)
	} else if tag == "required" {
		return fmt.Sprintf("\"%s\" is required", field)
	} else {
		return fmt.Sprintf("Invalid \"%s\"", field)
	}
}
