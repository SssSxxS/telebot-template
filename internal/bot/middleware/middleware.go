package middleware

import (
	"time"

	"telebot-template/internal/database/models"
	r "telebot-template/internal/database/repositories"

	tele "gopkg.in/telebot.v4"
	"gorm.io/gorm"
)

// Ignores messages older than the given duration
func IgnoreOld(dur time.Duration) tele.MiddlewareFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			if time.Since(c.Message().Time()) > dur {
				return nil
			}
			return next(c)
		}
	}
}

// Tracks user in the database
func UserTracker() tele.MiddlewareFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			sender := c.Sender()
			newUser := &models.User{
				Username: &sender.Username,
				Status:   1, // active
			}

			existingUser, err := r.UserRepo.GetByTelegramID(sender.ID)

			switch err {
			case nil: // user exists
				if existingUser.Status < 0 { // user is banned
					return nil
				}
				if err = r.UserRepo.Update(existingUser.ID, newUser); err != nil {
					return err
				}

			case gorm.ErrRecordNotFound: // user does not exist
				newUser.TelegramID = sender.ID
				if err = r.UserRepo.Create(newUser); err != nil {
					return err
				}

			default: // other error
				return err
			}

			return next(c)
		}
	}
}

// Checks if user is admin
func IsAdmin() tele.MiddlewareFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			sender := c.Sender()
			if user, err := r.UserRepo.GetByTelegramID(sender.ID); err != nil {
				return err
			} else if !user.IsAdmin {
				return nil
			}
			return next(c)
		}
	}
}
