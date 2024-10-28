package usecase

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/sidiqPratomo/DJKI-Pengaduan/appconstant"
	"github.com/sidiqPratomo/DJKI-Pengaduan/apperror"
	"github.com/sidiqPratomo/DJKI-Pengaduan/dto"
	"github.com/sidiqPratomo/DJKI-Pengaduan/repository"
	"github.com/sidiqPratomo/DJKI-Pengaduan/util"
)

type AuthenticationUsecase interface {
	RegisterUser(ctx context.Context, registerDTO dto.RegisterRequest) error
}

type authenticationUsecaseImpl struct {
	userRepository repository.UserRepository
	emailHelper    util.EmailHelper
	transaction    repository.Transaction
	hashHelper     util.HashHelperIntf
	JwtHelper      util.JwtAuthentication
}

type AuthenticationUsecaseImplOpts struct {
	UserRepository repository.UserRepository
	Transaction    repository.Transaction
	HashHelper     util.HashHelperIntf
	JwtHelper      util.JwtAuthentication
	EmailHelper    util.EmailHelper
}

func NewAuthenticationUsecaseImpl(opts AuthenticationUsecaseImplOpts) authenticationUsecaseImpl {
	return authenticationUsecaseImpl{
		userRepository: opts.UserRepository,
		transaction:    opts.Transaction,
		emailHelper:    opts.EmailHelper,
		JwtHelper:      opts.JwtHelper,
		hashHelper:     opts.HashHelper,
	}
}

func (u *authenticationUsecaseImpl) RegisterUser(ctx context.Context, registerDTO dto.RegisterRequest) error {
	account := dto.RegisterRequestToAccount(registerDTO)

	isNameValid := util.RegexValidate(account.Username, appconstant.NameRegexPattern)
	if !isNameValid {
		return apperror.InvalidNameError(errors.New("invalid name"))
	}

	existingAccountByEmail, err := u.userRepository.FindAccountByEmail(ctx, account.Email)
	if err != nil {
		return apperror.InternalServerError(err)
	}

	if existingAccountByEmail != nil {
		if err := u.updateOTPAndSendEmail(ctx, int(existingAccountByEmail.Id), existingAccountByEmail.Email); err != nil {
			return err
		}
		return nil
	}

	existingAccountByUsername, err := u.userRepository.FindAccountByUsername(ctx, account.Username)
	if err != nil {
		return apperror.InternalServerError(err)
	}

	if existingAccountByUsername != nil {
		return errors.New("username has been taken")
	}

	hashedPassword, err := u.hashHelper.HashPassword(account.Password)
	if err != nil {
		return apperror.InternalServerError(err)
	}
	account.Password = hashedPassword

	tx, err := u.transaction.BeginTx()
	if err != nil {
		return apperror.InternalServerError(err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	accountRepo := tx.UserRepository()
	accountId, err := accountRepo.PostOneUser(ctx, account)
	if err != nil {
		tx.Rollback()
		return apperror.InternalServerError(err)
	}
	accountId64 := int(*accountId)

	if err := u.createAndSendOTP(ctx, &accountId64, account.Email); err != nil {
		tx.Rollback()
		return apperror.InternalServerError(err)
	}

	if err := tx.Commit(); err != nil {
		return apperror.InternalServerError(err)
	}

	return nil
}

func (u *authenticationUsecaseImpl) updateOTPAndSendEmail(ctx context.Context, userId int, email string) error {
	// Generate a new OTP
	otp, err := u.userRepository.CreateOTP(ctx, strconv.Itoa(userId))
	if err != nil {
		return err
	}

	// Step 5: Send the OTP via email
	err = u.sendOTPEmail(email, *otp)
	if err != nil {
		return err
	}

	return nil
}

func (u *authenticationUsecaseImpl) createAndSendOTP(ctx context.Context, userId *int, email string) error {
	// Generate a new OTP
	otp, err := u.userRepository.CreateOTP(ctx, strconv.Itoa(*userId))
	if err != nil {
		return err
	}

	// Step 5: Send the OTP via email
	err = u.sendOTPEmail(email, *otp)
	if err != nil {
		return err
	}

	return nil
}

func (u *authenticationUsecaseImpl) sendOTPEmail(email string, otp string) error {
	// Define the subject and the email template
	subject := "Your OTP Code"
	emailTemplate := `<p>Your OTP code is <strong>{{.OTP}}</strong>. It will expire in 10 minutes.</p>`

	// Set the recipient(s) and subject
	u.emailHelper.AddRequest([]string{email}, subject)

	// Generate the email body with the OTP code
	data := map[string]interface{}{
		"OTP": otp,
	}
	err := u.emailHelper.CreateBody(emailTemplate, data)
	if err != nil {
		return fmt.Errorf("failed to create email body: %w", err)
	}

	// Send the email
	err = u.emailHelper.SendEmail()
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
