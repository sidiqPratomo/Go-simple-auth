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
	VerifyUserRegister(ctx context.Context, verifyDTO dto.VerifyOTPRequest) error
	LoginUser(ctx context.Context, loginDTO dto.LoginRequest) error
	VerifyUserLogin(ctx context.Context, verifyOtpLogin dto.VerifyUserLoginRequest) (*dto.VerifyLoginUserResponse, error)
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

func (u *authenticationUsecaseImpl) VerifyUserLogin(ctx context.Context, verifyOtpLogin dto.VerifyUserLoginRequest) (*dto.VerifyLoginUserResponse, error){
	fmt.Println("masuk:::::::::::::::::")
	accountUsername, err :=u.userRepository.FindAccountByUsername(ctx, verifyOtpLogin.Username)
	if err != nil {
		return nil, apperror.InternalServerError(err)
	}
	if accountUsername.EmailVerifiedAt == nil{
		return nil, apperror.NewAppError(400, err, "User Not verified")
	}
	fmt.Println("LOLOs1:::::::::::::::::")
	otpDetails, err := u.userRepository.GetOTPByCode(ctx, verifyOtpLogin.OTP, int(accountUsername.Id))
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, apperror.BadRequestError(errors.New("invalid or expired OTP"))
		}
		return nil, apperror.InternalServerError(err)
	}
	fmt.Println("LOLOs2:::::::::::::::::")
	user, err := u.userRepository.FindAccountByUserId(ctx, int(otpDetails.User_id))
	if err != nil {
		return nil, apperror.InternalServerError(err)
	}
	fmt.Println("LOLOs2:::::::::::::::::")
	customClaims := util.JwtCustomClaims{UserId: user.Id, Email: user.Email, Role: user.RoleName, TokenDuration: 15}
	token, expired, err := u.JwtHelper.CreateAndSign(customClaims, u.JwtHelper.Config.AccessSecret)
	if err != nil {
		return nil, apperror.InternalServerError(err)
	}
	fmt.Println("LOLOs3:::::::::::::::::")
	roles, privileges, err := u.userRepository.GetUserRoles(ctx, user.Id)
	if err != nil {
		return nil, apperror.InternalServerError(err)
	}
	fmt.Println("LOLOS4::::::::::::::::::::::")
	roleDTOs := dto.MapRolesToDTOs(roles)
	privilegeDTOs := dto.MapPrivilegesToDTOs(privileges)

	userDetails := dto.User{
		Id:              user.Id,
		Nik:			 user.Nik,
		Photo:           &user.Photo,
		FirstName:       user.FirstName,
		LastName:        user.LastName,
		Username:        user.Username,
		Email:           user.Email,
		Gender:          user.Gender,
		Address:         user.Address,
		PhoneNumber:     user.PhoneNumber,
		EmailVerifiedAt: *user.EmailVerifiedAt,
		Status:          user.Status,
		Role:            []string{user.RoleName}, // Set roles accordingly
		Roles:			 roleDTOs,
	}

	roleResponse := dto.Role{
		Privileges: privilegeDTOs,
		Role:       roleDTOs,
	}

	response := &dto.VerifyLoginUserResponse{
		User:        userDetails,
		Role:        roleResponse,
		AccessToken: *token,
		ExpiresAt:   *expired,
	}

	return response, nil
}

func (u *authenticationUsecaseImpl) LoginUser(ctx context.Context, loginDTO dto.LoginRequest) error {
	user, err := u.userRepository.FindAccountByUsername(ctx, loginDTO.Username)
	if err != nil {
		if err == repository.ErrNotFound {
			return apperror.BadRequestError(errors.New("username not found"))
		}
		return apperror.InternalServerError(err)
	}

	isPasswordValid,err := u.hashHelper.CheckPassword(loginDTO.Password,[]byte(user.Password))
	if !isPasswordValid {
		return apperror.WrongPasswordError(err)
	}
	accountId64 := int(user.Id)

	if err := u.createAndSendOTP(ctx, &accountId64, user.Email);err !=nil{
		return apperror.InternalServerError(err)
	}

	return nil
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
		return apperror.InternalServerError(err)
	}

	// Send the email
	err = u.emailHelper.SendEmail()
	if err != nil {
		return apperror.InternalServerError(err)
	}

	return nil
}

func (u *authenticationUsecaseImpl) VerifyUserRegister(ctx context.Context, verifyDTO dto.VerifyOTPRequest) error {
	// Step 1: Find user by Username
	//FIX Find By OTP and username and expired_at >= now()
	account, err := u.userRepository.FindAccountByUsername(ctx, verifyDTO.Username)
	if err != nil{
		return apperror.InternalServerError(err)
	}

	// Step 2: Verify the OTP
	isValid, err := u.userRepository.VerifyOTP(ctx, int(account.Id), verifyDTO.OTP)
	if err != nil {
		return apperror.InternalServerError(err)
	}
	if !isValid {
		return apperror.BadRequestError(errors.New("invalid or expired OTP"))
	}

	tx, err :=u.transaction.BeginTx()
	if err != nil{
		return apperror.InternalServerError(err)
	}
	defer func(){
		if err != nil {
			tx.Rollback()
		}
	}()
	updateUserTx := tx.UserRepository()

	// Step 3: Update the user's email verified status
	err = updateUserTx.UpdateUserVerificationStatus(ctx, int(account.Id))
	if err != nil {
		tx.Rollback()
		return apperror.InternalServerError(err)
	}

	if err := tx.Commit(); err != nil {
		return apperror.InternalServerError(err)
	}

	return nil
}
