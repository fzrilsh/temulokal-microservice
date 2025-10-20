package utils

import (
	"temulokal-microservice/shared-service/response"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

func ProxyRequest(c *fiber.Ctx, target string) error {
	req := c.Request()
	res := c.Response()

	client := &fasthttp.Client{}

	req.URI().SetHost(target[7:])
	req.URI().SetScheme("http")
	req.URI().SetPathBytes(req.URI().Path())

	if err := client.Do(req, res); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	return nil
}
