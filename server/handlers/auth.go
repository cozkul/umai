package handlers

import (
	"strings"
	"time"

	"github.com/cozkul/umai/server/config"
	"github.com/cozkul/umai/server/database"
	"github.com/cozkul/umai/server/helpers"
	"github.com/cozkul/umai/server/models"
	"github.com/cozkul/umai/server/schemas"
	"github.com/gofiber/fiber/v2"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func SignUpUser(c *fiber.Ctx) error {
	payload := new(schemas.SignUpInput)

	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password+config.Config.Salt), bcrypt.DefaultCost)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	newUser := models.User{
		Username: payload.UserName,
		Email:    strings.ToLower(payload.Email),
		Password: string(hashedPassword),
	}

	if err := helpers.InitializeUser(&newUser); err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "Unable to initialize new account",
			"error":   err.Error(),
		})
	}

	if result := database.DB.Create(&newUser); result.Error != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "Unable to initialize new account",
			"error":   result.Error.Error(),
		})
	}

	c.Locals("user", models.FilteredUser{
		ID:       newUser.ID,
		UserName: newUser.Username,
		Email:    newUser.Email,
	})

	// Success status code
	return c.Status(fiber.StatusCreated).JSON(models.FilteredUser{
		ID:       newUser.ID,
		UserName: newUser.Username,
		Email:    newUser.Email,
	})
}

func SignInUser(c *fiber.Ctx) error {
	payload := new(schemas.SignInInput)

	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var user models.User
	if result := database.DB.First(&user, "email = ?", strings.ToLower(payload.Email)); result.Error != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Invalid email or Password",
		})
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password+config.Config.Salt)); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Invalid email or Password",
		})
	}

	tokenByte := jwt.New(jwt.SigningMethodHS256)
	now := time.Now().UTC()
	claims := tokenByte.Claims.(jwt.MapClaims)
	claims["sub"] = user.ID
	claims["exp"] = now.Add(time.Hour * 24 * config.Config.JwtExpiresIn).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	tokenString, err := tokenByte.SignedString([]byte(config.Config.JwtSecret))

	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:        "token",
		Value:       tokenString,
		Path:        "/",
		SessionOnly: false,
		MaxAge:      config.Config.JwtMaxAge * 60 * 24,
		Secure:      false,
		HTTPOnly:    true,
		Domain:      "localhost",
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "token": tokenString})
}

func LogoutUser(c *fiber.Ctx) error {
	expired := time.Now().Add(-time.Hour * 24)
	c.Cookie(&fiber.Cookie{
		Name:    "token",
		Value:   "",
		Expires: expired,
	})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success"})
}
