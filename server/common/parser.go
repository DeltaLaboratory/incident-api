package common

import (
	"errors"
	"fmt"

	"github.com/fxamacker/cbor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

type Validatable interface {
	Validate() error
}

func ParseBody(ctx *fiber.Ctx, request any) error {
	contentType := utils.ToLower(ctx.Get("Content-Type"))
	contentType = utils.ParseVendorSpecificContentType(contentType)

	switch contentType {
	case "application/cbor":
		if err := cbor.Unmarshal(ctx.Body(), request); err != nil {
			return fmt.Errorf("failed to parse request: %w", err)
		}
	default:
		if err := ctx.BodyParser(request); err != nil {
			if errors.Is(err, fiber.ErrUnprocessableEntity) {
				return fmt.Errorf("content type not supported: %w", err)
			}
			return fmt.Errorf("failed to parse request: %w", err)
		}
	}

	if validatable, ok := request.(Validatable); ok {
		if err := validatable.Validate(); err != nil {
			return fmt.Errorf("invalid request: %w", err)
		}
	}

	return nil
}
