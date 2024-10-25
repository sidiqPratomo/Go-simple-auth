package appconstant

const (
	VerificationEmailSubject = "Verify Your Email Address to Activate Your MaxHealth Account"

	VerificationEmailTemplate = `
		<!DOCTYPE html>

		<html>

		<head>
			<title>VERIFICATION EMAIL</title>
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
                <p>Thank you for signing up for our service! To complete the registration process and ensure the security of your account, we kindly ask you to verify your email address.</p>
                <p>Please click the following link to verify your email and proceed with the registration process:</p>
                <a class="verify-button" href="{{.Url}}">Verify Your Email</a>
                <p>Upon clicking the link, you will be directed to our service's page where you will be prompted to enter the verification code provided below:</p>
                <p>Verification Code: <strong class="verification-code">{{.Code}}</strong></p>
                <p>Once you've entered the code, you have to set up your password and access your account.</p>
                <p>If you did not request this verification, please disregard this email.</p>
                <p>Thank you for choosing our service. We're excited to have you on board!</p>
                <p>Best regards,<br>MaxHealth Team</p>
            </div>
		</body>

		</html>
    `
)
