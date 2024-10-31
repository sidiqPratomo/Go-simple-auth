package usecase

import (
	"context"
	"errors"
	"fmt"
	"strconv"

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
	existingAccountByEmail, err := u.userRepository.FindAccountByEmail(ctx, registerDTO.Email)
	if err != nil && err != repository.ErrNotFound {
		return apperror.InternalServerError(err)
	}
	existingAccountByUsername, err := u.userRepository.FindAccountByUsername(ctx, registerDTO.Username)
	if err != nil && err != repository.ErrNotFound {
		return apperror.InternalServerError(err)
	}

	if existingAccountByEmail != nil && existingAccountByUsername != nil {
		// Check if the account has not verified the email yet
		if existingAccountByEmail.EmailVerifiedAt == nil { // Now using nil check for *time.Time
			if err := u.updateOTPAndSendEmail(ctx, int(existingAccountByEmail.Id), existingAccountByEmail.Email); err != nil {
				return err
			}
			return nil
		}
		return apperror.BadRequestError(errors.New("user already exists and verified"))
	}

	if existingAccountByUsername != nil {
		return errors.New("username has been taken")
	}

	account, err := dto.RegisterRequestToAccount(registerDTO)
	if err != nil{
		return err
	}

	hashedPassword, err := u.hashHelper.HashPassword(registerDTO.Password)
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
	err = accountRepo.CreateRoleUser(ctx, accountId64)
	if err != nil {
		tx.Rollback()
		return apperror.InternalServerError(err)
	}

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
	otp, err := u.userRepository.CreateOTP(ctx, strconv.Itoa(userId))
	if err != nil {
		return err
	}

	err = u.sendOTPEmail(email, *otp)
	if err != nil {
		return err
	}

	return nil
}

func (u *authenticationUsecaseImpl) createAndSendOTP(ctx context.Context, userId *int, email string) error {
	otp, err := u.userRepository.CreateOTP(ctx, strconv.Itoa(*userId))
	if err != nil {
		return err
	}

	err = u.sendOTPEmail(email, *otp)
	if err != nil {
		return err
	}

	return nil
}

func (u *authenticationUsecaseImpl) sendOTPEmail(email string, otp string) error {
	// Define the subject and the email template
	subject := "Pengaduan DJKI-OTP"
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

func (u *authenticationUsecaseImpl) VerifyUserRegister(ctx context.Context, verifyDTO dto.VerifyOTPRequest) error {
	// Step 1: Find user by ID
	//FIX Find By OTP and expired_at >= now()
	otp := verifyDTO.OTP
	existingAccount, err := u.userRepository.FindAccountByEmail(ctx, otp)
	if err != nil {
		return apperror.InternalServerError(err)
	}
	if existingAccount == nil {
		return apperror.NotFoundError()
	}

	// Step 2: Verify the OTP
	isValid, err := u.userRepository.VerifyOTP(ctx, otp, verifyDTO.OTP)
	if err != nil {
		return apperror.InternalServerError(err)
	}
	if !isValid {
		return apperror.BadRequestError(errors.New("invalid OTP"))
	}

	// Step 3: Update the user's email verified status
	err = u.userRepository.UpdateUserVerificationStatus(ctx, otp)
	if err != nil {
		return apperror.InternalServerError(err)
	}

	return nil
}
