package controller

import (
	"github.com/n3wscott/knative-operator/pkg/controller/knative"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, knative.Add)
}
