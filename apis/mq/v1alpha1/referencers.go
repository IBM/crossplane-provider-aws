/*
Copyright 2021 The Crossplane Authors.

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

package v1alpha1

import (
	"context"

	"github.com/crossplane/crossplane-runtime/pkg/reference"

	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"

	ec2 "github.com/crossplane/provider-aws/apis/ec2/v1beta1"
)

// ResolveReferences of this Broker
func (mg *Broker) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	// Resolve spec.forProvider.SubnetIds
	mrsp, err := r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SubnetIDs),
		References:    mg.Spec.ForProvider.SubnetIDRefs,
		Selector:      mg.Spec.ForProvider.SubnetIDSelector,
		To:            reference.To{Managed: &ec2.Subnet{}, List: &ec2.SubnetList{}},
		Extract:       reference.ExternalName(),
	})
	if err != nil {
		return errors.Wrap(err, "spec.forProvider.SubnetIDs")
	}
	mg.Spec.ForProvider.SubnetIDs = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SubnetIDRefs = mrsp.ResolvedReferences

	// Resolve spec.forProvider.SecurityGroups
	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SecurityGroups),
		References:    mg.Spec.ForProvider.SecurityGroupIDRefs,
		Selector:      mg.Spec.ForProvider.SecurityGroupIDSelector,
		To:            reference.To{Managed: &ec2.SecurityGroup{}, List: &ec2.SecurityGroupList{}},
		Extract:       reference.ExternalName(),
	})
	if err != nil {
		return errors.Wrap(err, "spec.forProvider.SecurityGroups")
	}
	mg.Spec.ForProvider.SecurityGroups = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SecurityGroupIDRefs = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this User
func (mg *User) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	// Resolve spec.forProvider.brokerID
	rsp, err := r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.BrokerID),
		Reference:    mg.Spec.ForProvider.BrokerIDRef,
		Selector:     mg.Spec.ForProvider.BrokerIDSelector,
		To:           reference.To{Managed: &Broker{}, List: &BrokerList{}},
		Extract:      reference.ExternalName(),
	})
	if err != nil {
		return errors.Wrap(err, "spec.forProvider.brokerID")
	}
	mg.Spec.ForProvider.BrokerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BrokerIDRef = rsp.ResolvedReference

	return nil
}
