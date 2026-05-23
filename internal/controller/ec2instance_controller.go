/*
Copyright 2026.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	logf "sigs.k8s.io/controller-runtime/pkg/log"

	ec2v1alpha1 "github.com/inboxamitraj/ec2operator/api/v1alpha1"
)

// EC2InstanceReconciler reconciles a EC2Instance object
type EC2InstanceReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=ec2.mycompany.io,resources=ec2instances,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=ec2.mycompany.io,resources=ec2instances/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=ec2.mycompany.io,resources=ec2instances/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the EC2Instance object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.23.3/pkg/reconcile
func (r *EC2InstanceReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = logf.FromContext(ctx)

	// TODO(user): your logic here

	l := log.FromContext(ctx)

	ec2instance := &ec2v1alpha1.EC2Instance{}

	if err := r.Get(ctx, req.NamespacedName, ec2instance); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	l.Info("Reconciling EC2Instance", "Name", ec2instance.Name)

	fmt.Println("I got a request for an ec2 instance in the namespace", req.NamespacedName)
	fmt.Println("The ec2 instance name is ", ec2instance.Name)
	fmt.Println("Instance type is ", ec2instance.Spec.Type)
	fmt.Println("Ami ID is ", ec2instance.Spec.AmiID)
	fmt.Println("Subnet is ", ec2instance.Spec.Subnet)
	fmt.Println("Tags are ", ec2instance.Spec.Tags)
	fmt.Println("Storage is", ec2instance.Spec.Storage.Size, "GB", "and type is", ec2instance.Spec.Storage.Type)

	l.Info("Reconciled Ec2 Instance", "Name", ec2instance.Name)

	return ctrl.Result{}, nil
	
	}


// SetupWithManager sets up the controller with the Manager.
func (r *EC2InstanceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&ec2v1alpha1.EC2Instance{}).
		Named("ec2instance").
		Complete(r)
}
