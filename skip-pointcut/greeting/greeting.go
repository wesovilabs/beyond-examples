package greeting

import (
	"errors"
	"fmt"
)

func Hello(firstName string) error{
	if firstName == "" {
		return errors.New("[ERR] invalid firstName")
	}
	fmt.Printf("Hey %s\n", firstName)
	return nil
}

func Bye(firstName string) error{
	if firstName == "" {
		return errors.New("[ERR] invalid firstName")
	}
	fmt.Printf("Bye %s\n", firstName)
	return nil
}


func Greetings(mode string, firstName string) error {
	switch mode {
	case "Hello":
		return Hello(firstName)
	case "Bye":
		return Bye(firstName)
	default:
		return errors.New("[ERR] unexpected greeting")
	}
}
