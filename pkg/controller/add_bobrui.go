package controller

import (
	"github.com/dxas90/bobr-operator/pkg/controller/bobrui"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, bobrui.Add)
}
