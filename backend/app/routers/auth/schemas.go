package auth

type RegistrationRequestSchema struct {
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=6"`
	Firstname   string `json:"firstname" validate:"required,min=1"`
	Lastname    string `json:"lastname" validate:"required,min=1"`
	PhoneNumber string `json:"phone_number" validate:"required,e164"` // international telephone numbering standard
}
