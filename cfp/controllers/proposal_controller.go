/*
Copyright 2022.

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

package controllers

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	talksv1 "github.com/scottrigby/how-to-write-a-reconciler-using-k8s-controller-runtime/cfp/api/v1"
)

// ProposalReconciler reconciles a Proposal object
type ProposalReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=talks.kubecon.na,resources=proposals,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=talks.kubecon.na,resources=proposals/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=talks.kubecon.na,resources=proposals/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Proposal object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.12.2/pkg/reconcile
func (r *ProposalReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// Get the proposal Object

	// Set reconciling condition on the proposal object

	// Set a finalizer on the proposal object if not set

	// Check if this a deletion, if yes api call to delete the Speaker
	// clean the finalizer
	// Clean the condition

	// Get the speaker referenced object
	// If does not, exist
	// Set a conditon: SpeakerNotFoundCondition
	// return and requeue

	// If we have a Submision on the ProposalStatus sub resource
	// Make an api call with the submission string to get the Proposal
	// If it is not found
	// Set a condition like CNCFProposalrErrorConditon
	// error and requeue
	// else compare the propsal with the spec
	// Update if needed
	// Set the  ready condition (or any error condition)
	// return and requeue

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ProposalReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&talksv1.Proposal{}).
		Complete(r)
}