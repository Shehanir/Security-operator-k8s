package controller

import (
	"github.com/security-operator/security-operator/pkg/controller/security"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, security.Add)
}
