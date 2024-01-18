package services

const (

	// Response Status constant
	SUCCESS             string = "Success"
	USER_SIGNIN_SUCCESS string = "User Sign In Sccuess"
	USER_SIGNUP_SUCCESS string = "User Sign Up Sccuess"

	ERROR               string = "Error"
	EMAIL_PASSWORD_NULL string = "Email or Password is null"
	INPUT_INVALID       string = "Input is invalid"
	MONGO_ERROR         string = "Something wrong on mongoDB"
	INVALID_TOKEN       string = "Invalid ID token"
	MISSING_AUTH_HEADER string = "Missing Authorization header"

	Authorization string = "Authorization"

	// validation constant
	TOKEN string = "token"
)
