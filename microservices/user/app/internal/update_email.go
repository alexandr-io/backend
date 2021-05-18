package internal

import (
	"context"
	"encoding/json"
	"os"

	"github.com/alexandr-io/backend/common/generate"
	"github.com/alexandr-io/backend/user/data"
	grpcclient "github.com/alexandr-io/backend/user/grpc/client"
	"github.com/alexandr-io/backend/user/redis"

	"github.com/gofiber/fiber/v2"
)

// updateEmail is the internal logic function used to update and verify an email.
func updateEmail(ctx context.Context, user *data.User, newEmail string) (string, error) {
	verifyEmailToken := generate.RandomStringNoSpecialChar(12)

	if err := redis.VerifyEmail.Create(ctx, verifyEmailToken, data.EmailVerification{
		OldEmail: user.Email,
		NewEmail: newEmail,
	}); err != nil {
		return "", err
	}

	return verifyEmailToken, nil
}

// VerifyEmailCreation store email info in Redis and send an email for verification
func VerifyEmailCreation(ctx context.Context, user *data.User) error {
	verifyEmailToken, err := updateEmail(ctx, user, user.Email)
	if err != nil {
		return err
	}
	verifyEmailURL := os.Getenv("USER_URI") + "/verify?token=" + verifyEmailToken

	if err := grpcclient.SendEmail(ctx, data.Email{
		Email:    user.Email,
		Username: user.Username,
		Type:     data.VerifyEmail,
		Data:     verifyEmailURL,
	}); err != nil {
		return err
	}
	return nil
}

// VerifyEmailUpdate store email info in Redis and send two emails.
// One email to the new email address to verify the email, another one to the previous email address to give the possibility to cancel
func VerifyEmailUpdate(ctx context.Context, user *data.User, newEmail string) error {
	verifyEmailToken, err := updateEmail(ctx, user, newEmail)
	if err != nil {
		return err
	}
	verifyEmailURL := os.Getenv("USER_URI") + "/verify/update?token=" + verifyEmailToken
	cancelEmailUpdateURL := os.Getenv("USER_URI") + "/email/cancel?token=" + verifyEmailToken

	if err := grpcclient.SendEmail(ctx, data.Email{
		Email:    newEmail,
		Username: user.Username,
		Type:     data.UpdateEmailVerify,
		Data:     verifyEmailURL,
	}); err != nil {
		return err
	}

	emailData, err := json.Marshal(&struct {
		NewEmail string `json:"new_email"`
		Link     string `json:"link"`
	}{
		NewEmail: newEmail,
		Link:     cancelEmailUpdateURL,
	})
	if err != nil {
		return data.NewHTTPErrorInfo(fiber.StatusInternalServerError, err.Error())
	}
	if err := grpcclient.SendEmail(ctx, data.Email{
		Email:    user.Email,
		Username: user.Username,
		Type:     data.UpdateEmailOldEmail,
		Data:     string(emailData),
	}); err != nil {
		return err
	}
	return nil
}
