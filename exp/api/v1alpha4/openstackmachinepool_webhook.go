/*
Copyright 2020 The Kubernetes Authors.
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

package v1alpha4

import (
	"reflect"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"

	infrav1 "sigs.k8s.io/cluster-api-provider-openstack/api/v1alpha4"
)

// log is for logging in this package.
var openstackmachinepoollog = logf.Log.WithName("openstackmachinepool-resource")

func (omp *OpenStackMachinePool) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(omp).
		Complete()
}

// +kubebuilder:webhook:path=/mutate-exp-cluster-x-k8s-io-v1alpha4-openstackmachinepool,mutating=true,failurePolicy=fail,matchPolicy=Equivalent,groups=exp.cluster.x-k8s.io,resources=openstackmachinepools,verbs=create;update,versions=v1alpha4,name=default.openstackmachinepool.infrastructure.cluster.x-k8s.io

var _ webhook.Defaulter = &OpenStackMachinePool{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (omp *OpenStackMachinePool) Default() {
	openstackmachinepoollog.Info("default", "name", omp.Name)

	if omp.Spec.Template.Spec.Template.Spec.IdentityRef != nil && omp.Spec.Template.Spec.Template.Spec.IdentityRef.Kind == "" {
		omp.Spec.Template.Spec.Template.Spec.IdentityRef.Kind = defaultIdentityRefKind
	}
}

// +kubebuilder:webhook:verbs=create;update,path=/validate-exp-cluster-x-k8s-io-v1alpha4-openstackmachinepool,mutating=false,failurePolicy=fail,matchPolicy=Equivalent,groups=exp.cluster.x-k8s.io,resources=openstackmachinepools,versions=v1alpha4,name=validation.openstackmachinepool.exp.infrastructure.cluster.x-k8s.io

var _ webhook.Validator = &OpenStackMachinePool{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (omp *OpenStackMachinePool) ValidateCreate() error {
	openstackmachinepoollog.Info("validate create", "name", omp.Name)

	var allErrs field.ErrorList

	if omp.Spec.Template.Spec.Template.Spec.ProviderID != nil {
		allErrs = append(allErrs, field.Forbidden(field.NewPath("spec", "template", "spec", "template", "spec", "providerID"), "cannot be set in templates"))
	}

	return aggregateObjErrors(omp.GroupVersionKind().GroupKind(), omp.Name, allErrs)
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (omp *OpenStackMachinePool) ValidateUpdate(old runtime.Object) error {
	var allErrs field.ErrorList
	openstackmachinepoollog.Info("validate update", "name", omp.Name)

	oldOpenStackMachinePool := old.(*OpenStackMachinePool)

	if !reflect.DeepEqual(omp.Spec.Template.Spec.Template.Spec, oldOpenStackMachinePool.Spec.Template.Spec.Template.Spec) {
		allErrs = append(allErrs,
			field.Invalid(field.NewPath("spec", "template", "spec", "template", "spec"), omp, infrav1.OpenStackMachineTemplateImmutableMsg),
		)
	}

	return aggregateObjErrors(omp.GroupVersionKind().GroupKind(), omp.Name, allErrs)
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (omp *OpenStackMachinePool) ValidateDelete() error {
	openstackmachinepoollog.Info("validate delete", "name", omp.Name)
	return nil
}
