package server

import (
	"fmt"
	"net/http"

	apn "github.com/codesee/basic-golang-app/pkg/aliaspackagename"
	"github.com/codesee/basic-golang-app/pkg/logger"
	"github.com/codesee/basic-golang-app/pkg/regularpackagename"
	"github.com/codesee/basic-golang-app/pkg/wrongpackagename"
	"github.com/labstack/echo"
	elog "github.com/labstack/gommon/log"
)

type SomeType string

type SomeStruct struct {}

type SomeInterface interface {}

const SomeConst = "ConstValue"

var SomeVar SomeType = "VarValue"

const (
	ConstGroup1 = 1
	ConstGroup2 = 2
)

var somePrivateVar = 1

func New(port int) (*http.Server, error) {
	e := echo.New()

	e.Logger.SetLevel(elog.OFF)

	e.Use(logger.Middleware())

	regularpackagename.RegisterRoutes(e)
	otherpackagename.RegisterRoutes(e)
	apn.RegisterRoutes(e)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: e,
	}

	return srv, nil
}