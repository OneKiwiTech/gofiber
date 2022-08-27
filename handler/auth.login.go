package handler

import (
	"strconv"
	"time"

	"github.com/OneKiwiTech/gofiber/config"
	"github.com/OneKiwiTech/gofiber/model"
	"github.com/OneKiwiTech/gofiber/service"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/google/uuid"

	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req LoginRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Email, validation.Required, is.Email),
		validation.Field(&req.Password, validation.Required, validation.Length(8, 18)),
	)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Login(ctx *fiber.Ctx) error {
	var (
		req  LoginRequest
		user model.User
	)
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(fiber.StatusBadGateway)
	}
	if errors := req.Validate(); errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}
	user, _ = service.GetUserByEmail(req.Email)
	if !CheckPasswordHash(req.Password, user.Password) {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Bad Credentials",
			"data":    nil,
		})
	}

	rawAccessToken := jwt.New(jwt.SigningMethodHS256)
	jwtTTL, _ := strconv.Atoi(config.LoadConfig("JWT_TTL"))
	claims := rawAccessToken.Claims.(jwt.MapClaims)
	claims["uid"] = user.ID
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(jwtTTL)).Unix()
	accessToken, err := rawAccessToken.SignedString([]byte(config.LoadConfig("JWT_SECRET")))
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	jwtRefreshTTL, _ := strconv.Atoi(config.LoadConfig("JWT_REFRESH_TTL"))
	token, _ := service.CreateRefreshToken(model.RefreshToken{
		AuthUUID:  uuid.New().String(),
		UserID:    user.ID,
		CreatedAt: time.Now(),
		ExpireAt:  time.Now().Add(time.Minute * time.Duration(jwtRefreshTTL)),
	})

	rawRefreshToken := jwt.New(jwt.SigningMethodHS256)
	claims = rawRefreshToken.Claims.(jwt.MapClaims)
	claims["uid"] = token.UserID
	claims["auid"] = token.AuthUUID
	claims["exp"] = token.ExpireAt.Unix()
	refreshToken, err := rawRefreshToken.SignedString([]byte(config.LoadConfig("JWT_SECRET")))
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.JSON(fiber.Map{
		"message": "Login Successful",
		"data": fiber.Map{
			"access_token":  accessToken,
			"refresh_token": refreshToken,
		},
	})
}
