package user

type Request struct {
	Name           string `json:"name"  validate:"required,gte=3,lte=24"`
	Email          string `json:"email" validate:"required,email"`
	MasterPassword string `json:"masterPassword" validate:"required,gte=6,lte=244"`
}
