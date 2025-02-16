package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

// GEO will fetch IP GEO data from ip-api.com
func GEO() fiber.Handler {
	// Create fasthttp client
	client := fasthttp.Client{}

	// Return handler
	return func(c *fiber.Ctx) error {
		// Get domain from query else default to remote IP
		ip := c.Params("ip", c.IP())

		// Get request/response from pool
		req := fasthttp.AcquireRequest()
		res := fasthttp.AcquireResponse()
		defer fasthttp.ReleaseRequest(req)
		defer fasthttp.ReleaseResponse(res)

		// Set request URL
		req.SetRequestURI("http://ip-api.com/json/" + ip)

		// Make request
		if err := client.DoTimeout(req, res, 30*time.Second); err != nil {
			return err
		}

		// Check statuscode
		if res.StatusCode() != fiber.StatusOK {
			return fiber.ErrRequestTimeout
		}

		// Set correct content-type
		c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

		// Send body
		return c.Send(res.Body())
	}
}
