package common

import (
	// "log"
	// "net/http"

	"github.com/gin-gonic/gin"
)

// BindJSONAndValidate binds incoming JSON to a given struct and handles errors.
func BindJSONAndValidate(ctx *gin.Context, obj interface{}) error {
	err := ctx.BindJSON(obj)
	if err != nil {
		// SendError(ctx, http.StatusBadRequest, 0, "binding JSON failed", err)
		if err := HandleRequestError(ctx, err, "binding JSON failed"); err != nil {
			return err // Exit if an error occurred (response is already sent)
		}
		return err
	}
	return nil
}
