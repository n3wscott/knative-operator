package main

import (
	"fmt"
	duckapis "github.com/knative/pkg/apis"
	"github.com/n3wscott/knative-operator/pkg/yaml"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func main() {
	cf := yaml.ConfigFile{
		Path: "./yaml/istio-v0.3.0.yaml",
	}
	if err := cf.Read(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Printf("%#v\n", cf)

	// apiGroup
	// resource
	// scope

	things := make(map[string]string)

	for _, v := range cf.Resources {

		kind := v.GetKind()

		vers := v.GetAPIVersion()

		gv, err := schema.ParseGroupVersion(vers)
		if err != nil {
			return
		}

		gvk := gv.WithKind(kind)

		gvr := duckapis.KindToResource(gvk)

		thing := fmt.Sprintf("%s - %s - %s",
			v.GetNamespace(),
			gvr.Group,
			gvr.Resource,
		)

		things[thing] = thing

	}

	for k, _ := range things {
		fmt.Println(k)
	}

}
