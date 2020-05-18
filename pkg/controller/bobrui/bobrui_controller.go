package bobrui

import (
	"context"

	devopsv1alpha1 "github.com/dxas90/bobr-operator/pkg/apis/devops/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_bobrui")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new BobRUI Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileBobRUI{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("bobrui-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource BobRUI
	err = c.Watch(&source.Kind{Type: &devopsv1alpha1.BobRUI{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner BobRUI
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &devopsv1alpha1.BobRUI{},
	})
	if err != nil {
		return err
	}

	return nil
}

// blank assignment to verify that ReconcileBobRUI implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileBobRUI{}

// ReconcileBobRUI reconciles a BobRUI object
type ReconcileBobRUI struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a BobRUI object and makes changes based on the state read
// and what is in the BobRUI.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileBobRUI) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling BobRUI")

	// Fetch the BobRUI instance
	instance := &devopsv1alpha1.BobRUI{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	// Define a new Deployment object
	deploy := newDeploymentForCR(instance)

	// Set BobRDB instance as the owner and controller
	if err := controllerutil.SetControllerReference(instance, deploy, r.scheme); err != nil {
		return reconcile.Result{}, err
	}

	// Check if this Deployment already exists
	found := &appsv1.Deployment{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: deploy.Name, Namespace: deploy.Namespace}, found)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("Creating a new Deployment", "Deployment.Namespace", deploy.Namespace, "Deployment.Name", deploy.Name, "Request.Replicas", deploy.Spec.Replicas)
		err = r.client.Create(context.TODO(), deploy)
		if err != nil {
			return reconcile.Result{}, err
		}

		// Deployment created successfully - don't requeue
		return reconcile.Result{}, nil
	} else if err != nil {
		return reconcile.Result{}, err
	}

	// Ensure the deployment size is the same as the spec
	size := instance.Spec.Replicas
	if *found.Spec.Replicas != size {
		found.Spec.Replicas = &size
		err = r.client.Update(context.TODO(), found)
		if err != nil {
			reqLogger.Error(err, "Failed to update Deployment.", "Deployment.Namespace", found.Namespace, "Deployment.Name", found.Name)
			return reconcile.Result{}, err
		}
		reqLogger.Info("Updating  Deployment replicas", "Deployment.Namespace", found.Namespace, "Deployment.Name", found.Name, "Request.Replicas", found.Spec.Replicas)
		// Spec updated - return and requeue
		return reconcile.Result{Requeue: true}, nil
	}
	// Deployment already exists - don't requeue
	reqLogger.Info("Skip reconcile: Deployment already exists", "Deployment.Namespace", found.Namespace, "Deployment.Name", found.Name, "Request.Replicas", found.Spec.Replicas)
	return reconcile.Result{}, nil
}

// newDeploymentForCR returns a deployment name/namespace as the cr
func newDeploymentForCR(cr *devopsv1alpha1.BobRUI) *appsv1.Deployment {
	labels := map[string]string{
		"app":    cr.Name,
		"devops": cr.Name,
	}
	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name + "-deploy",
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &cr.Spec.Replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  cr.Name,
							Image: cr.Spec.Image,
						},
					},
				},
			},
		},
	}
}
