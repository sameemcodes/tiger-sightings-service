// error.go

package durable

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomError struct {
	Code    int
	Message string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

func HandleError(ctx *gin.Context, err error) {
	var customErr *CustomError
	if errors.As(err, &customErr) {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error_code": customErr.Code, "error_message": customErr.Message})
	} else {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

var (
	ErrPhotoUpload            = &CustomError{Code: 4001, Message: "Error during photo upload and resizing"}
	ErrTigerSighting          = &CustomError{Code: 4002, Message: "Tiger sighting within 5 km"}
	ErrCreateNewTigerSighting = &CustomError{Code: 4003, Message: "Error during CreateNewTigerSighting"}
)
