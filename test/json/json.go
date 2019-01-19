package main

import (
	"fmt"
	"k8s.io/client-go/kubernetes/scheme"
	//_ "k8s.io/kubernetes/pkg/apis/extensions/install"

	"gopkg.in/yaml.v2"
)

var json = `
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: virtualservices.networking.istio.io
  annotations:
    "helm.sh/hook": crd-install
  labels:
    app: istio-pilot
spec:
  group: networking.istio.io
  names:
    kind: VirtualService
    listKind: VirtualServiceList
    plural: virtualservices
    singular: virtualservice
    categories:
    - istio-io
    - networking-istio-io
  scope: Namespaced
  version: v1alpha3
`

func main() {

	// decode := api.Codecs.UniversalDecoder().Decode
	decode := scheme.Codecs.UniversalDeserializer().Decode

	obj, _, err := decode([]byte(json), nil, nil)
	if err != nil {
		fmt.Printf("%#v", err)
	}

	fmt.Printf("%#v\n", obj)
}
