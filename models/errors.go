package models

import "strings"

const (
	// ErrNotFound is return when a resource cannot be found
	// in a database.
	ErrNotFound modelError = "models: resource not found"

	// ErrPasswordIncorrect is returned when an invalid password
	// is used when attempting to authenticate a user.
	ErrPasswordIncorrect modelError = "models: incorrect password provided"

	// ErrEmailRequired is return when an empty email
	// is passed mainly during registration.
	ErrEmailRequired modelError = "models: email address is required"

	// ErrEmailInvalid is returned when the email format
	// passed in does not match emailRegex.
	ErrEmailInvalid modelError = "models: email address is not valid"

	// ErrEmailTaken is returned when an update or create
	// is attempted with an email address that is already in use.
	ErrEmailTaken modelError = "models: email address is already taken"

	// ErrPasswordTooShort is returned when an update or create is
	// attempted with a user password that is less than 8 characters
	ErrPasswordTooShort modelError = "models: password must be at least 8 characters long"

	// ErrPasswordRequired is returned when a create is attempted
	// without a user password provided
	ErrPasswordRequired modelError = "models: password is required"

	ErrTitleRequired modelError = "models: title is required"

	// ErrRememberTooShort is returned when a remember token is not at least
	// 32 bytes
	ErrRememberTooShort privateError = "models: remember token must be at least 32 bytes"

	// ErrPasswordRequired is returned when a create or update is attempted
	// without a valid user remember token hash
	ErrRememberRequired privateError = "models: remember token is required"

	ErrUserIDRequired privateError = "models: user ID is required"

	// ErrInvalidID is returned when an invalid ID is provided
	// to a method like Delete.
	ErrInvalidID privateError = "models: ID provided was invalid"

)

type modelError string

func (e modelError) Error() string {
	return string(e)
}

func (e modelError) Public() string {
	s := strings.Replace(string(e), "models: ", "", 1)
	split := strings.Split(s, " ")
	split[0] = strings.Title(split[0])
	return strings.Join(split, " ")
}

type privateError string

func (p privateError) Error() string {
	return string(p)
}