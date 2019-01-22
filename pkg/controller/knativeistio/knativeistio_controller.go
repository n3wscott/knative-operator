package knativeistio

import (
	"context"
	duckapis "github.com/knative/pkg/apis"
	knativev1alpha1 "github.com/n3wscott/knative-operator/pkg/apis/knative/v1alpha1"
	"github.com/n3wscott/knative-operator/pkg/yaml"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_knativeistio")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new KnativeIstio Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileKnativeIstio{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("knativeistio-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource KnativeIstio
	err = c.Watch(&source.Kind{Type: &knativev1alpha1.KnativeIstio{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner KnativeIstio
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &knativev1alpha1.KnativeIstio{},
	})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileKnativeIstio{}

// ReconcileKnativeIstio reconciles a KnativeIstio object
type ReconcileKnativeIstio struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client        client.Client
	scheme        *runtime.Scheme
	dynamicClient dynamic.Interface

	crds *yaml.ConfigFile
	core *yaml.ConfigFile
}

const (
	CRDS_CONFIG_FILE_PATH = "/etc/config/istio-crds-v0.3.0.yaml"
	CORE_CONFIG_FILE_PATH = "/etc/config/istio-v0.3.0.yaml"
)

func (r *ReconcileKnativeIstio) InjectConfig(c *rest.Config) error {
	var err error
	r.dynamicClient, err = dynamic.NewForConfig(c)
	return err
}

func (r *ReconcileKnativeIstio) UpdateConfig() {
	if r.crds == nil {
		r.crds = &yaml.ConfigFile{
			Path: CRDS_CONFIG_FILE_PATH,
		}
		if err := r.crds.Read(); err != nil {
			log.Error(err, "error reading config file %q", r.crds.Path)
		}
	}
	logger := log.WithValues("Config", r.crds.Path)
	logger.Info("Updated CRD Config", "ConfigMap.Istio.Resources", len(r.crds.Resources))
}

func (r *ReconcileKnativeIstio) UpdateCore() {
	if r.core == nil {
		r.core = &yaml.ConfigFile{
			Path: CORE_CONFIG_FILE_PATH,
		}
		if err := r.core.Read(); err != nil {
			log.Error(err, "error reading config file %q", r.core.Path)
		}
	}
	logger := log.WithValues("Config", r.core.Path)
	logger.Info("Updated Core Config", "ConfigMap.Istio.Resources", len(r.core.Resources))
}

func (r *ReconcileKnativeIstio) ReconcileIstioCRDs(request reconcile.Request) error {
	logger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	logger.Info("Reconciling Knative Istio CRDs")

	for k, v := range r.crds.Resources {
		logger.Info("inspecting resource", "Reconcile.Istio.ResourceKey", k)

		kind := v.GetKind()

		vers := v.GetAPIVersion()

		gv, err := schema.ParseGroupVersion(vers)
		if err != nil {
			return err
		}

		gvk := gv.WithKind(kind)

		gvr := duckapis.KindToResource(gvk)

		gvrClient := r.dynamicClient.Resource(gvr)

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
				logger.Info("need to create resource", "Reconcile.Istio.NeedCreation", k)

				_, err := gvrC.Create(&v, metav1.CreateOptions{})
				if err != nil {
					logger.Error(err, "failed to create")
				}

			} else {
				return err
			}
			continue
		}
		logger.Info("resource valid", "Reconcile.Istio.Resource.Exists", k)
	}
	return nil
}

func (r *ReconcileKnativeIstio) ReconcileIstioCore(request reconcile.Request) error {
	logger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	logger.Info("Reconciling Knative Istio Core")

	for k, v := range r.core.Resources {
		logger.Info("inspecting resource", "Reconcile.Istio.ResourceKey", k)

		kind := v.GetKind()

		vers := v.GetAPIVersion()

		gv, err := schema.ParseGroupVersion(vers)
		if err != nil {
			return err
		}

		gvk := gv.WithKind(kind)

		gvr := duckapis.KindToResource(gvk)

		gvrClient := r.dynamicClient.Resource(gvr)

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
				logger.Info("need to create resource", "Reconcile.Istio.NeedCreation", k)

				_, err := gvrC.Create(&v, metav1.CreateOptions{})
				if err != nil {
					logger.Error(err, "failed to create")
				}
				continue
			} else {
				logger.Error(err, "failed to get")
				continue
			}
		}
		logger.Info("resource valid", "Reconcile.Istio.Resource.Exists", k)
	}
	return nil
}

// Reconcile reads that state of the cluster for a KnativeIstio object and makes changes based on the state read
// and what is in the KnativeIstio.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileKnativeIstio) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling KnativeIstio")

	r.UpdateConfig()

	if err := r.ReconcileIstioCRDs(request); err != nil {
		return reconcile.Result{}, err
	}

	r.UpdateCore()
	if err := r.ReconcileIstioCore(request); err != nil {
		return reconcile.Result{}, err
	}

	// Fetch the KnativeIstio instance
	instance := &knativev1alpha1.KnativeIstio{}
	if err := r.client.Get(context.TODO(), request.NamespacedName, instance); err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	// Define a new Pod object
	pod := newPodForCR(instance)

	// Set KnativeIstio instance as the owner and controller
	if err := controllerutil.SetControllerReference(instance, pod, r.scheme); err != nil {
		return reconcile.Result{}, err
	}

	// Check if this Pod already exists
	found := &corev1.Pod{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Name: pod.Name, Namespace: pod.Namespace}, found)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("Creating a new Pod", "Pod.Namespace", pod.Namespace, "Pod.Name", pod.Name)
		err = r.client.Create(context.TODO(), pod)
		if err != nil {
			return reconcile.Result{}, err
		}

		// Pod created successfully - don't requeue
		return reconcile.Result{}, nil
	} else if err != nil {
		return reconcile.Result{}, err
	}

	// Pod already exists - don't requeue
	reqLogger.Info("Skip reconcile: Pod already exists", "Pod.Namespace", found.Namespace, "Pod.Name", found.Name)
	return reconcile.Result{}, nil
}

// newPodForCR returns a busybox pod with the same name/namespace as the cr
func newPodForCR(cr *knativev1alpha1.KnativeIstio) *corev1.Pod {
	labels := map[string]string{
		"app": cr.Name,
	}
	return &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name + "-pod",
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:    "busybox",
					Image:   "busybox",
					Command: []string{"sleep", "3600"},
				},
			},
		},
	}
}
