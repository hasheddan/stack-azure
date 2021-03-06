/*
Copyright 2019 The Crossplane Authors.

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

package mysqlservervirtualnetworkrule

import (
	"context"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplaneio/stack-azure/apis/database/v1alpha3"
	azurev1alpha3 "github.com/crossplaneio/stack-azure/apis/v1alpha3"
	azure "github.com/crossplaneio/stack-azure/pkg/clients"

	runtimev1alpha1 "github.com/crossplaneio/crossplane-runtime/apis/core/v1alpha1"
	"github.com/crossplaneio/crossplane-runtime/pkg/meta"
	"github.com/crossplaneio/crossplane-runtime/pkg/resource"
)

// Error strings.
const (
	errNewClient                           = "cannot create new MySQLServerVirtualNetworkRule"
	errNotMySQLServerVirtualNetworkRule    = "managed resource is not an MySQLServerVirtualNetworkRule"
	errCreateMySQLServerVirtualNetworkRule = "cannot create MySQLServerVirtualNetworkRule"
	errUpdateMySQLServerVirtualNetworkRule = "cannot update MySQLServerVirtualNetworkRule"
	errGetMySQLServerVirtualNetworkRule    = "cannot get MySQLServerVirtualNetworkRule"
	errDeleteMySQLServerVirtualNetworkRule = "cannot delete MySQLServerVirtualNetworkRule"
)

// Controller is responsible for adding the MySQLServerVirtualNetworkRule
// Controller and its corresponding reconciler to the manager with any runtime configuration.
type Controller struct{}

// SetupWithManager creates a new MySQLServerVirtualNetworkRule Controller and adds it to the
// Manager with default RBAC. The Manager will set fields on the Controller and
// start it when the Manager is Started.
func (c *Controller) SetupWithManager(mgr ctrl.Manager) error {
	r := resource.NewManagedReconciler(mgr,
		resource.ManagedKind(v1alpha3.MySQLServerVirtualNetworkRuleGroupVersionKind),
		resource.WithManagedConnectionPublishers(),
		resource.WithExternalConnecter(&connecter{client: mgr.GetClient()}))

	name := strings.ToLower(fmt.Sprintf("%s.%s", v1alpha3.MySQLServerVirtualNetworkRuleKind, v1alpha3.Group))

	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		For(&v1alpha3.MySQLServerVirtualNetworkRule{}).
		Complete(r)
}

type connecter struct {
	client      client.Client
	newClientFn func(ctx context.Context, credentials []byte) (azure.MySQLVirtualNetworkRulesClient, error)
}

func (c *connecter) Connect(ctx context.Context, mg resource.Managed) (resource.ExternalClient, error) {
	v, ok := mg.(*v1alpha3.MySQLServerVirtualNetworkRule)
	if !ok {
		return nil, errors.New(errNotMySQLServerVirtualNetworkRule)
	}

	p := &azurev1alpha3.Provider{}
	n := meta.NamespacedNameOf(v.Spec.ProviderReference)
	if err := c.client.Get(ctx, n, p); err != nil {
		return nil, errors.Wrapf(err, "cannot get provider %s", n)
	}

	s := &corev1.Secret{}
	n = types.NamespacedName{Namespace: p.Spec.Secret.Namespace, Name: p.Spec.Secret.Name}
	if err := c.client.Get(ctx, n, s); err != nil {
		return nil, errors.Wrapf(err, "cannot get provider secret %s", n)
	}
	newClientFn := azure.NewMySQLVirtualNetworkRulesClient
	if c.newClientFn != nil {
		newClientFn = c.newClientFn
	}
	client, err := newClientFn(ctx, s.Data[p.Spec.Secret.Key])
	return &external{client: client}, errors.Wrap(err, errNewClient)
}

type external struct {
	client azure.MySQLVirtualNetworkRulesClient
}

func (e *external) Observe(ctx context.Context, mg resource.Managed) (resource.ExternalObservation, error) {
	v, ok := mg.(*v1alpha3.MySQLServerVirtualNetworkRule)
	if !ok {
		return resource.ExternalObservation{}, errors.New(errNotMySQLServerVirtualNetworkRule)
	}

	az, err := e.client.Get(ctx, v.Spec.ResourceGroupName, v.Spec.ServerName, v.Spec.Name)
	if azure.IsNotFound(err) {
		return resource.ExternalObservation{ResourceExists: false}, nil
	}
	if err != nil {
		return resource.ExternalObservation{}, errors.Wrap(err, errGetMySQLServerVirtualNetworkRule)
	}

	azure.UpdateMySQLVirtualNetworkRuleStatusFromAzure(v, az)
	v.SetConditions(runtimev1alpha1.Available())

	o := resource.ExternalObservation{
		ResourceExists:    true,
		ConnectionDetails: resource.ConnectionDetails{},
	}

	return o, nil
}

func (e *external) Create(ctx context.Context, mg resource.Managed) (resource.ExternalCreation, error) {
	v, ok := mg.(*v1alpha3.MySQLServerVirtualNetworkRule)
	if !ok {
		return resource.ExternalCreation{}, errors.New(errNotMySQLServerVirtualNetworkRule)
	}

	v.SetConditions(runtimev1alpha1.Creating())

	vnet := azure.NewMySQLVirtualNetworkRuleParameters(v)
	if _, err := e.client.CreateOrUpdate(ctx, v.Spec.ResourceGroupName, v.Spec.ServerName, v.Spec.Name, vnet); err != nil {
		return resource.ExternalCreation{}, errors.Wrap(err, errCreateMySQLServerVirtualNetworkRule)
	}

	return resource.ExternalCreation{}, nil
}

func (e *external) Update(ctx context.Context, mg resource.Managed) (resource.ExternalUpdate, error) {
	v, ok := mg.(*v1alpha3.MySQLServerVirtualNetworkRule)
	if !ok {
		return resource.ExternalUpdate{}, errors.New(errNotMySQLServerVirtualNetworkRule)
	}

	az, err := e.client.Get(ctx, v.Spec.ResourceGroupName, v.Spec.ServerName, v.Spec.Name)
	if err != nil {
		return resource.ExternalUpdate{}, errors.Wrap(err, errGetMySQLServerVirtualNetworkRule)
	}

	if azure.MySQLServerVirtualNetworkRuleNeedsUpdate(v, az) {
		vnet := azure.NewMySQLVirtualNetworkRuleParameters(v)
		if _, err := e.client.CreateOrUpdate(ctx, v.Spec.ResourceGroupName, v.Spec.ServerName, v.Spec.Name, vnet); err != nil {
			return resource.ExternalUpdate{}, errors.Wrap(err, errUpdateMySQLServerVirtualNetworkRule)
		}
	}
	return resource.ExternalUpdate{}, nil
}

func (e *external) Delete(ctx context.Context, mg resource.Managed) error {
	v, ok := mg.(*v1alpha3.MySQLServerVirtualNetworkRule)
	if !ok {
		return errors.New(errNotMySQLServerVirtualNetworkRule)
	}

	v.SetConditions(runtimev1alpha1.Deleting())

	_, err := e.client.Delete(ctx, v.Spec.ResourceGroupName, v.Spec.ServerName, v.Spec.Name)

	return errors.Wrap(resource.Ignore(azure.IsNotFound, err), errDeleteMySQLServerVirtualNetworkRule)
}
