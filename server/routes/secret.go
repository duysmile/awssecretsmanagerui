package routes

import (
	"net/http"

	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/labstack/echo"
	"github.com/ledongthuc/awssecretsmanagerui/server/actions"
)

func setupSecretRoutes(g *echo.Group) {
	g.GET("", func(c echo.Context) error {
		secrets, err := actions.GetListSecrets()
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, secrets)
	})

	g.GET("/detail", func(c echo.Context) error {
		arn := c.QueryParam("arn")
		if len(arn) == 0 {
			return echo.NewHTTPError(http.StatusInternalServerError, "Missing selected ARN")
		}
		secret, err := actions.GetSecretByARN(arn)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, secret)
	})

	g.GET("/value", func(c echo.Context) error {
		arn := c.QueryParam("arn")
		if len(arn) == 0 {
			return echo.NewHTTPError(http.StatusInternalServerError, "Missing selected ARN")
		}
		secret, err := actions.GetSecretValueByARN(arn)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, secret)
	})

	g.POST("/value/update", func(c echo.Context) error {
		var request secretsmanager.PutSecretValueInput
		err := c.Bind(&request)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		secret, err := actions.UpdateSecretValue(request)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, secret)
	})
}
