package handlers

import (
	"fmt"
	"my_go_project/internal/services"

	"github.com/labstack/echo/v4"
)

func GetIPInfo(c echo.Context) error {
	ip := c.RealIP() // Client IP adresini alır

	fmt.Println("ip: ", ip)

	ipInfo, err := services.FetchIPInfo(ip)
	if err != nil {
		return c.JSON(400, err)
	}

	return c.JSON(200, ipInfo)
}
