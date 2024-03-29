package credentials

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
	"github.com/techave-dev/init-go-be/internal/repo/psql"
	"github.com/techave-dev/init-go-be/tools"
)

type Handlers struct {
	config   *tools.Config
	services Services
}

func NewHandlers(config *tools.Config, services Services) Handlers {
	return Handlers{config, services}
}

type LoginHandlerRequest = LoginParams

func (l *LoginHandlerRequest) Validate() error {
	return validation.ValidateStruct(l,
		validation.Field(&l.Email, validation.Required, is.Email),
		validation.Field(&l.Password, validation.Required, validation.Length(8, 255)),
	)
}

type RegisterHandlerRequest struct {
	RegisterParams
	RetypePassword string `json:"retypePassword"`
}

func (r *RegisterHandlerRequest) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Email, validation.Required, is.Email),
		validation.Field(&r.Password, validation.Required, validation.Length(8, 255), validation.By(func(value interface{}) error {
			if value.(string) != r.RetypePassword {
				return fmt.Errorf("kombinasi kata sandi salah")
			}
			return nil
		})),
		validation.Field(&r.RetypePassword, validation.Required, validation.Length(8, 255)),
		validation.Field(&r.Role, validation.Required, validation.By(func(value interface{}) error {
			var err error = fmt.Errorf("invalid role: %s ", value)
			switch enum := value.(type) {
			case psql.RoleEnum:
				if valid := enum.Valid(); valid {
					err = nil
				}
			}
			return err
		})),
	)
}

func (h *Handlers) Register() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req RegisterHandlerRequest
		if err := c.BodyParser(&req); err != nil {
			return tools.Fail(c, fiber.StatusConflict, tools.String(err.Error()))

		}

		if err := req.Validate(); err != nil {
			return tools.Fail(c, fiber.StatusConflict, tools.String(err.Error()))
		}

		credential, err := h.services.Register(c.Context(), req.RegisterParams)
		if err != nil {
			return tools.Fail(c, fiber.StatusConflict, tools.String("Kredensial gagal didaftarkan"))
		}

		return tools.Success(c, 201, tools.Null(), tools.R{"credential": credential})
	}
}

func (h *Handlers) Me() fiber.Handler {
	return func(c *fiber.Ctx) error {
		credentialID := c.Locals("CredentialID").(string)
		id, err := uuid.FromString(credentialID)
		if err != nil {
			return tools.Fail(c, fiber.StatusBadRequest, tools.String("Format id tidak valid"))
		}

		credential, err := h.services.CredentialById(c.Context(), id)
		if err != nil {
			return tools.Fail(c, fiber.StatusNotFound, tools.String("Kredensial tidak ditemukan"))
		}

		return tools.Success(c, fiber.StatusOK, tools.Null(), tools.R{"credential": credential})
	}
}

func (h *Handlers) Login() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req LoginHandlerRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(map[string]any{"error": err.Error()})
		}

		if err := req.Validate(); err != nil {
			return c.Status(400).JSON(map[string]any{"error": err.Error()})
		}

		data, err := h.services.Login(c.Context(), req)
		if err != nil {
			return c.Status(400).JSON(map[string]any{"error": err.Error()})
		}

		return tools.Success(c, fiber.StatusOK, tools.Null(), tools.ToMap(data))
	}
}
