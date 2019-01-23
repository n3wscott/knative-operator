package yaml

import (
	duckapis "github.com/knative/pkg/apis"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

var log = logf.Log.WithName("helper_yaml")

func ReconcileConfig(cf *ConfigFile, request reconcile.Request, dynamicClient dynamic.Interface) error {
	logger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	logger.Info("Reconciling CRDs")

	for k, v := range cf.Resources {
		logger.Info("inspecting resource", "Reconcile.Key", k)

		kind := v.GetKind()

		if kind == "ClusterRole" {
			continue
		}
		if kind == "ClusterRoleBinding" {
			continue
		}
		if kind == "RoleBinding" {
			continue
		}
		if kind == "Role" {
			continue
		}
		if kind == "Namespace" {
			continue
		}

		vers := v.GetAPIVersion()

		gv, err := schema.ParseGroupVersion(vers)
		if err != nil {
			return err
		}

		gvk := gv.WithKind(kind)

		gvr := duckapis.KindToResource(gvk)

		gvrClient := dynamicClient.Resource(gvr)

		name := v.GetName()
		namespace := v.GetNamespace()

		var gvrC dynamic.ResourceInterface

		if namespace == "" {
			gvrC = gvrClient
		} else {
			gvrC = gvrClient.Namespace(namespace)
		}

		_, err = gvrC.Get(name, metav1.GetOptions{})
		if err != nil {
			if errors.IsNotFound(err) {
				logger.Info("need to create resource", "Reconcile.NeedCreation", k)

				_, err := gvrC.Create(&v, metav1.CreateOptions{})
				if err != nil {
					logger.Error(err, "failed to create")
				}

			} else {
				logger.Error(err, "get failed")
				//return err
			}
			continue
		}
		logger.Info("resource valid", "Reconcile.Resource.Exists", k)
	}
	return nil
}
