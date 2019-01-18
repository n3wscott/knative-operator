package controller

import (
	"github.com/n3wscott/knative-operator/pkg/controller/knativeeventingsources"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, knativeeventingsources.Add)
}
