package utils

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type ErrorFeild struct {
	Message string `json:"message"`
	Feild string	`json:"feild"`
}
type CustomError struct{
	StatusCode int `json:"statusCode"`
	Errors []ErrorFeild `json:"errors"`
}

func RequestValidationError(ve validator.ValidationErrors) CustomError{
	errs := []ErrorFeild{}
	err := ErrorFeild{}

	
	for _, verr := range ve {
		err.Message = verr.Error()
		err.Feild = verr.Field()
		errs = append(errs, err)
	}


	return CustomError{
		StatusCode: http.StatusBadRequest,
		Errors: errs,
	}
}

func BadRequestError(msg string) CustomError{
	return CustomError{
		StatusCode: http.StatusBadRequest,
		Errors: []ErrorFeild{
			{Message: msg},
		},
	}
}

func InternalServerError(msg string) CustomError{
	return CustomError{
		StatusCode: http.StatusInternalServerError,
		Errors: []ErrorFeild{
			{Message: msg},
		},
	}
}

func NotAuthorisedError() CustomError{
	return CustomError{
		StatusCode: http.StatusUnauthorized,
		Errors: []ErrorFeild{
			{Message: "Not authorized"},
		},
	}
}

func NotFoundError(r string) CustomError{
	return CustomError{
		StatusCode: http.StatusNotFound,
		Errors: []ErrorFeild{
			{Message: fmt.Sprintf("%s Not found", r)},
		},
	}
}

func DatabaseConnectionError() CustomError{
	return CustomError{
		StatusCode: http.StatusInternalServerError,
		Errors: []ErrorFeild{
			{Message: "error connecting to DB"},
		},
	}
}