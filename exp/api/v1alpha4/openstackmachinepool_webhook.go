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
	"k8s.io/apimachinery/pkg/runtime"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
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
}

// +kubebuilder:webhook:verbs=create;update,path=/validate-exp-cluster-x-k8s-io-v1alpha4-openstackmachinepool,mutating=false,failurePolicy=fail,matchPolicy=Equivalent,groups=exp.cluster.x-k8s.io,resources=openstackmachinepools,versions=v1alpha4,name=validation.openstackmachinepool.exp.infrastructure.cluster.x-k8s.io

var _ webhook.Validator = &OpenStackMachinePool{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (omp *OpenStackMachinePool) ValidateCreate() error {
	openstackmachinepoollog.Info("validate create", "name", omp.Name)
	return omp.Validate()
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (omp *OpenStackMachinePool) ValidateUpdate(old runtime.Object) error {
	openstackmachinepoollog.Info("validate update", "name", omp.Name)
	return omp.Validate()
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (omp *OpenStackMachinePool) ValidateDelete() error {
	openstackmachinepoollog.Info("validate delete", "name", omp.Name)
	return nil
}

// Validate the Azure Machine Pool and return an aggregate error
func (omp *OpenStackMachinePool) Validate() error {
	validators := []func() error{
		omp.ValidateImage,
	}

	var errs []error
	for _, validator := range validators {
		if err := validator(); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return kerrors.NewAggregate(errs)
	}

	return nil
}

// ValidateImage of an OpenStackMachinePool
func (omp *OpenStackMachinePool) ValidateImage() error {
	return nil
}
