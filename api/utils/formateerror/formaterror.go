package formateerror

import (
	"errors"
	"strings"
)

/**
To format some error messages in a more readable manner, we need to a create a package to help us achieve that.
*/
func FormatError(err string) error {
	if strings.Contains(err, "name") {
		return errors.New("Name Already Taken")
	}

	if strings.Contains(err, "email") {
		return errors.New("Email Already Taken")
	}

	if strings.Contains(err, "title") {
		return errors.New("Title Already Taken")
	}

	if strings.Contains(err, "hashedPassword") {
		return errors.New("Incorect Password")
	}
	return errors.New("Incorrect Details")
}
