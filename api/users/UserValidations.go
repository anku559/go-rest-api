package users

var SignUpValidation struct {
	FirstName   *string
	MiddleName  *string
	LastName    *string
	Username    *string
	CountryCode *string
	Phone       *string

	Email    string
	Password string
}
