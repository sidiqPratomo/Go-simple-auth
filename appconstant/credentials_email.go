package appconstant

const (
	CredentialsEmailSubject = "Account Access Details"

	CredentialsEmailTemplate = `
		<!DOCTYPE html>

		<html>

		<head>
			<title>ACCOUNT ACCESS DETAILS</title>
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
                <p>Thank you for signing up for our service! Here are your account access details.</p>
                <p>Email: <strong>{{.Email}}</strong></p>
                <p>Credentials: <strong>{{.Credentials}}</strong></p>
                <p>Once you've logged in to your account, we kindly ask you to immediately change your password to ensure the security of your account.</p>
                <p>Thank you for choosing our service. We're excited to have you on board!</p>
                <p>Best regards,<br>MaxHealth Team</p>
            </div>
		</body>

		</html>
    `
)
