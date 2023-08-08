package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/haji-sudo/ShabehRoshan/db"
	"github.com/haji-sudo/ShabehRoshan/repository"
	"github.com/haji-sudo/ShabehRoshan/router/url"
	"github.com/haji-sudo/ShabehRoshan/util"
)

// Check if User signin can't access to login/signup Page and redirect to Home
func CheckAuthNotValid(c *fiber.Ctx) bool {
	session_id := c.Cookies("session_id")
	if session_id == "" {
		return true
	}
	sess, err := db.Store.Get(c)
	if err != nil {
		return true
	}
	var tokenString string
	if token, ok := sess.Get("token").(string); ok {
		tokenString = token
	}

	if tokenString == "" {
		sess.Destroy()
		return true
	}
	userid, err := util.ValidateToken(tokenString)
	if err != nil {
		if err.Error() == "expired" {
			newToken, err := refreshToken(userid)
			sess.Destroy()
			if err != nil {
				return true
			}
			sess.Regenerate()
			sess.Set("token", newToken)
			sess.Save()
			return false
		}
		sess.Destroy()
		return true
	}
	return false
}

// Check if User signin can't access to login/signup Page and redirect to Home
func IsAuth(c *fiber.Ctx) bool {
	session_id := c.Cookies("session_id")
	if session_id == "" {
		return false
	}
	sess, err := db.Store.Get(c)
	if err != nil {
		return false
	}
	var tokenString string
	if token, ok := sess.Get("token").(string); ok {
		tokenString = token
	}

	if tokenString == "" {
		sess.Destroy()
		return false
	}
	userid, err := util.ValidateToken(tokenString)
	if err != nil {
		if err.Error() == "expired" {
			newToken, err := refreshToken(userid)
			sess.Destroy()
			if err != nil {
				return false
			}
			sess.Regenerate()
			sess.Set("token", newToken)
			sess.Save()
			return true
		}
		sess.Destroy()
		return false
	}
	c.Locals("userid", userid)
	return true
}

// Middleware Check user Auth
func Auth(c *fiber.Ctx) error {
	session_id := c.Cookies("session_id")
	if session_id == "" {
		return c.Redirect(url.Login)
	}
	sess, err := db.Store.Get(c)
	if err != nil {
		return c.Redirect(url.Login)
	}
	var tokenString string
	if token, ok := sess.Get("token").(string); ok {
		tokenString = token
	}
	if tokenString == "" {
		sess.Destroy()
		return c.Redirect(url.Login)
	}
	userid, err := util.ValidateToken(tokenString)
	if err != nil {
		if err.Error() == "expired" {
			newToken, err := refreshToken(userid)
			sess.Destroy()
			if err != nil {
				return c.Redirect(url.Login)
			}
			sess.Regenerate()
			sess.Set("token", newToken)
			sess.Save()
			return c.Next()
		}
		sess.Destroy()
		return c.Redirect(url.Login)
	}
	c.Locals("userid", userid)
	return c.Next()
}

// Check if refresh token in db It had not expired generate new token and return
func refreshToken(UserID string) (string, error) {
	userRepo := repository.NewUserRepository()
	user, err := userRepo.GetByID(uuid.MustParse(UserID))
	userRepo.GetToken(user)
	if err != nil {
		return "", err
	}
	userid, err := util.ValidateToken(user.Token.RefreshToken)
	if err != nil {
		return "", err
	}
	if userid == UserID {
		tokenString, err := util.GenerateToken(*user)
		if err != nil {
			return "", nil
		}

		return tokenString, nil
	}

	return "", fmt.Errorf("userid not equals to refresh token userid")
}
