package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go-wedding/models"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{DB: db}
}

func (h *Handler) CreateInvite(c *fiber.Ctx) error {

	inviteRepo := models.NewInviteRepo(h.DB)

	invite := new(models.Invite)

	err := c.BodyParser(invite)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.GeneralResponse{
			StatusCode: fiber.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	err = invite.Validate()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.GeneralResponse{
			StatusCode: fiber.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	err = inviteRepo.CreateAlias(invite)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.GeneralResponse{
			StatusCode: fiber.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(models.GeneralResponse{
		StatusCode: fiber.StatusCreated,
		Message:    "success",
		Data:       nil,
	})

}

func (h *Handler) ScanInvite(c *fiber.Ctx) error {

	var err error

	inviteRepo := models.NewInviteRepo(h.DB)

	sc := new(models.ScanRequest)

	err = c.BodyParser(sc)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.GeneralResponse{
			StatusCode: fiber.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	i, err := inviteRepo.GetInviteBySecret(sc.Secret)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(models.GeneralResponse{
			StatusCode: fiber.StatusNotFound,
			Message:    "Invite not found with error " + err.Error(),
			Data:       nil,
		})
	}

	if i.InviteStatus == models.Used {

		return c.Status(fiber.StatusInternalServerError).JSON(models.GeneralResponse{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Invite used",
			Data:       nil,
		})
	}

	i.Uses += 1

	if i.Uses == i.MaxUsage {
		i.InviteStatus = models.Used
	}

	err = inviteRepo.UpdateInvite(i)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.GeneralResponse{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "failed to update with err " + err.Error(),
			Data:       nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.GeneralResponse{
		StatusCode: fiber.StatusOK,
		Message:    "success",
		Data:       i,
	})

}
