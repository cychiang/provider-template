package certificate

import (
	"crypto"
	"github.com/go-acme/lego/v4/registration"
)

type LegoUser struct {
	Email        string
	Registration *registration.Resource
	key          crypto.PrivateKey
}

func (lu *LegoUser) GetEmail() string {
	return lu.Email
}

func (lu LegoUser) GetRegistration() *registration.Resource {
	return lu.Registration
}

func (lu *LegoUser) GetPrivateKey() crypto.PrivateKey {
	return lu.key
}
