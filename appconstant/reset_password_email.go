package appconstant

const (
	ResetPasswordEmailSubject = "Reset Password"

	ResetPasswordEmailTemplate = `
		<!DOCTYPE html>

		<html>

		<head>
			<title>RESET PASSWORD</title>
			<style>
                .email-container {
                    border: 1px solid #ccc;
                    border-radius: 5px;
                    padding: 20px;
                }

				.verify-button {
                    background-color: #4CAF50;
                    color: white !important;
                    padding: 10px 20px;
                    text-decoration: none;
                    border-radius: 5px;
                }

                .verification-code{
                    font-weight: 600;
                    font-size: 16px;
                    letter-spacing: 10px;
                    margin: 15px 0px;
                }
			</style>
		</head>

		<body>
            <div class="email-container">
                <h2>Welcome to MaxHealth!</h2>
                <p>Hi {{.Name}},</p>
                <p>You recently request to reset your account password.</p>
                <p>Please click the following link to proceed with the reset password process:</p>
                <a class="verify-button" href="{{.Url}}">Reset Password</a>
                <p>Upon clicking the link, you will be directed to our service's page where you will be prompted to enter the verification code provided below:</p>
                <p>Verification Code: <strong class="verification-code">{{.Code}}</strong></p>
                <p>Once you've entered the code, you have to set up your new password and access your account.</p>
                <p>If you did not request this reset password, please disregard this email.</p>
                <p>Thank you for choosing our service. We're excited to have you on board!</p>
                <p>Best regards,<br>MaxHealth Team</p>
            </div>
		</body>

		</html>
    `
)
