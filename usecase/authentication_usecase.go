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

type AuthenticationUsecase interface{
	RegisterUser(ctx context.Context) error 
}

type authenticationUsecaseImpl struct{
	userRepository repository.UserRepository
	emailHelper util.EmailHelper
}

// type AuthenticationUsecaseImplOpts struct{
// 	userRepository repository.UserRepository
// }

func NewAuthenticationUsecaseImpl(UserRepository repository.UserRepository) authenticationUsecaseImpl{
	return authenticationUsecaseImpl{
		userRepository: UserRepository,
	}
}

func (u *authenticationUsecaseImpl)RegisterUser(ctx context.Context, registerDTO dto.RegisterRequest) error {
	// Convert the DTO to an Account entity
	account := dto.RegisterRequestToAccount(registerDTO)

	isNameValid := util.RegexValidate(account.Username, appconstant.NameRegexPattern)
	if !isNameValid {
		return apperror.InvalidNameError(errors.New("invalid name"))
	}

	// Step 1: Check if the email is already registered
	existingAccountByEmail, err := u.userRepository.FindAccountByEmail(ctx, account.Email)
	if err != nil {
		return err
	}

	if existingAccountByEmail != nil {
		// If email is found, update OTP and send it again
		err := u.updateOTPAndSendEmail(ctx, int(existingAccountByEmail.Id), existingAccountByEmail.Email)
		if err != nil {
			return err
		}
		return nil
	}

	// Step 2: Check if the username is already taken
	existingAccountByUsername, err := u.userRepository.FindAccountByUsername(ctx, account.Username)
	if err != nil {
		return err
	}

	if existingAccountByUsername != nil {
		return errors.New("username has been taken")
	}

	// Step 3: Create a new user if both email and username are not taken
	accountId, err := u.userRepository.PostOneUser(ctx, account)
	if err != nil {
		return err
	}

	// Step 4: Generate OTP and associate it with the new user
	err = u.createAndSendOTP(ctx, accountId, account.Email)
	if err != nil {
		return err
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