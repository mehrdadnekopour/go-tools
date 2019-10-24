package helpers

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/mehrdadnekopour/go-tools/mypes"
	"github.com/mehrdadnekopour/go-tools/templates"
	validator "gopkg.in/go-playground/validator.v9"

	// "./templates"
	"github.com/labstack/echo"
)

// EchoContext ...
// var EchoContext echo.Context

const (
	// EchoHeaderUser ...
	EchoHeaderUser string = "x-user"
	// EchoHeaderRoles ...
	EchoHeaderRoles string = "x-rols"
	// EchoHeaderRedirectPath ...
	EchoHeaderRedirectPath string = "x-path"
)

// GetCurrentURL ...
func GetCurrentURL(ctx echo.Context) (url string) {
	r := ctx.Request()
	url = ctx.Scheme() + "://" + r.Host + r.URL.Path
	return
}

// GetModifyPathInfo ....
func GetModifyPathInfo(ctx echo.Context) (userInterface interface{}, method, path string) {
	r := ctx.Request()
	method = fmt.Sprintf("%s", r.Method)
	path = ctx.Get(EchoHeaderRedirectPath).(string)
	// path := fmt.Sprintf("%s", r.URL.RequestURI())

	userInterface = ctx.Get(EchoHeaderUser)

	return
}

// ParsePaginationParams ...
func ParsePaginationParams(ctx echo.Context) (limit, offset int, err mypes.Merror) {
	limit, _ = strconv.Atoi(ctx.QueryParam("limit"))
	offset, _ = strconv.Atoi(ctx.QueryParam("offset"))

	if limit < 0 {
		e := errors.New("Limit must be positive")
		err.Set(true, e, mypes.HTTPBadRequest)
	}

	if limit == 0 {
		limit = 50
	}

	if limit > 100 {
		limit = 50
	}

	return
}

// ParseQueryParamInt ...
func ParseQueryParamInt(ctx echo.Context, key string) (val int, merr mypes.Merror) {
	valStr := ctx.QueryParam(key)

	if valStr == "" {
		merr.Set(true, errors.New("invalid query param"), mypes.HTTPBadRequest)
		return 0, merr
	}

	val, e := strconv.Atoi(valStr)

	if e != nil {
		val = -1
		merr.SetInvalidQueryParamMrror()
	}
	return
}

// ParseQueryParamListInt ...
func ParseQueryParamListInt(ctx echo.Context, key, devider string) (output []int, merror mypes.Merror) {

	if key == "" {
		return
	}

	params := ctx.QueryParam(key)

	if params == "" {
		return
	}

	params = strings.Replace(params, "[", "", 1)
	params = strings.Replace(params, "]", "", 1)

	vals := strings.Split(params, devider)

	for _, val := range vals {

		v, e := strconv.Atoi(val)

		if e != nil {
			merror.SetInvalidQueryParamMrror()
			return nil, merror
		}

		output = append(output, v)
	}

	return
}

// ParseQueryParamListString ...
func ParseQueryParamListString(ctx echo.Context, key, devider string) (output []string, merr mypes.Merror) {
	if key == "" {
		return
	}

	params := ctx.QueryParam(key)

	if params == "" {
		return
	}

	params = strings.Replace(params, "[", "", 1)
	params = strings.Replace(params, "]", "", 1)

	output = strings.Split(params, devider)

	return
}

//

// Bind ...
func Bind(ctx echo.Context, i interface{}) (merr mypes.Merror) {
	e := ctx.Bind(i)

	if e != nil {
		merr.Set(true, e, mypes.HTTPUnprocessableEntity)
	}

	// Important! You must make mypes.NewValidate() before using this!
	if mypes.Validate != nil {
		e = mypes.Validate.Struct(i)
		if e != nil {

			if _, ok := e.(*validator.InvalidValidationError); ok {
				fmt.Println(e)
			} else {
				for _, err := range e.(validator.ValidationErrors) {

					fmt.Println(err.Namespace())
					fmt.Println(err.Field())
					fmt.Println(err.StructNamespace()) // can differ when a custom TagNameFunc is registered or
					fmt.Println(err.StructField())     // by passing alt name to ReportError like below
					fmt.Println(err.Tag())
					fmt.Println(err.ActualTag())
					fmt.Println(err.Kind())
					fmt.Println(err.Type())
					fmt.Println(err.Value())
					fmt.Println(err.Param())
					fmt.Println()
				}
			}

			merr.Set(true, e, http.StatusBadRequest)
		}
	}

	return
}

// UploadFile  ... Multipart
func UploadFile(ctx echo.Context, paramName, name, assetsPath, fileDirPath string) (dir, uniqueFileName, suffix, alias string, merr mypes.Merror) {
	uniqueFileName = ""

	//-----------
	// Read file
	//-----------

	// Source
	file, err := ctx.FormFile(paramName)
	if err != nil {
		merr.Set(true, err, mypes.HTTPInternalServerError)
		return
	}

	src, err := file.Open()
	if err != nil {
		merr.Set(true, err, mypes.HTTPInternalServerError)
		return
	}
	defer src.Close()

	_, dir, err = CreateDirIfNotExist(assetsPath, fileDirPath)

	if err != nil {
		merr.Set(true, err, mypes.HTTPInternalServerError)
		return
	}

	nameParts := strings.Split(file.Filename, ".")
	c := len(nameParts)

	suffix = nameParts[c-1]
	alias = strings.Join(nameParts[:c-1], ".")

	now := time.Now().Local()

	nowString := now.Format("20060102150405")

	uniqueFileName = fmt.Sprintf("%s%s", name, nowString)

	fileFullPath := dir + "/" + uniqueFileName + "." + suffix
	dst, err := os.Create(fileFullPath)
	if err != nil {
		merr.Set(true, err, mypes.HTTPInternalServerError)
		return
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		err = os.Remove(fileFullPath)
		merr.Set(true, err, mypes.HTTPInternalServerError)
		return
	}

	return
}

// Reply ...
func Reply(ctx echo.Context, content mypes.MSON) error {

	template := &templates.ResponseTemplate{}

	template = templates.Ok(
		content,
		nil,
	)
	return ctx.JSON(http.StatusOK, template)

}

// ReplyWithMeta ...
func ReplyWithMeta(ctx echo.Context, content mypes.MSON, meta interface{}) error {

	template := &templates.ResponseTemplate{}

	template = templates.Ok(
		content,
		meta,
	)
	return ctx.JSON(http.StatusOK, template)

}

// ReplyError ...
func ReplyError(ctx echo.Context, merror mypes.Merror) error {
	reqURL := GetCurrentURL(ctx)
	log.Println("error: ", merror, "url: ", reqURL)
	template := templates.GetWithCode(merror.HTTPStatus, merror.Data)
	return ctx.JSON(merror.HTTPStatus, template)
}
