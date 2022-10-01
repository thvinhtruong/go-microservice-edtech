package presenter

import (
	"SepFirst/UserService/app/apperror"
	"errors"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

var (
	Date = [...]string{"Mon", "Tues", "Weds", "Thu", "Fri", "Sat", "Sun"}
)

func ValidateDate(str string) error {
	for _, v := range Date {
		if str == v {
			return nil
		}
	}

	return apperror.ErrorWrongDateFormat
}

func CheckStringLength(strs []string, min int, max int, checkSpace bool) error {
	for _, str := range strs {
		err := validation.Validate(str, validation.Required, validation.Length(min, max), validation.When(checkSpace, validation.By(SpaceCheck)))
		if err == apperror.ErrorSpaceDetected {
			return err
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func SpaceCheck(value interface{}) error {
	s, _ := value.(string)
	if index := strings.Index(s, " "); index >= 0 && index <= len(s)-1 {
		return apperror.ErrorSpaceDetected
	}

	return nil
}

func ValidateUserAuth(username string, email string, password string) error {
	if len(username) == 0 && len(email) == 0 || len(password) == 0 {
		return errors.New("username or password is empty")
	}

	return nil
}

func ValidateUserEmptyField(req UserRequest) error {
	if len(req.Fullname) == 0 || len(req.Phone) == 0 ||
		len(req.Password) == 0 || len(req.Email) == 0 || len(req.Gender) == 0 {
		return apperror.ErrorEmptyField
	}

	return nil
}
