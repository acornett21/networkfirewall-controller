// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package firewall_policy

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"

	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	ackcondition "github.com/aws-controllers-k8s/runtime/pkg/condition"
	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
	ackrequeue "github.com/aws-controllers-k8s/runtime/pkg/requeue"
	ackrtlog "github.com/aws-controllers-k8s/runtime/pkg/runtime/log"
	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/networkfirewall"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws-controllers-k8s/networkfirewall-controller/apis/v1alpha1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = strings.ToLower("")
	_ = &aws.JSONValue{}
	_ = &svcsdk.NetworkFirewall{}
	_ = &svcapitypes.FirewallPolicy{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
	_ = &ackcondition.NotManagedMessage
	_ = &reflect.Value{}
	_ = fmt.Sprintf("")
	_ = &ackrequeue.NoRequeue{}
)

// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkFind")
	defer func() {
		exit(err)
	}()
	// If any required fields in the input shape are missing, AWS resource is
	// not created yet. Return NotFound here to indicate to callers that the
	// resource isn't yet created.
	if rm.requiredFieldsMissingFromReadOneInput(r) {
		return nil, ackerr.NotFound
	}

	input, err := rm.newDescribeRequestPayload(r)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.DescribeFirewallPolicyOutput
	resp, err = rm.sdkapi.DescribeFirewallPolicyWithContext(ctx, input)
	rm.metrics.RecordAPICall("READ_ONE", "DescribeFirewallPolicy", err)
	if err != nil {
		if reqErr, ok := ackerr.AWSRequestFailure(err); ok && reqErr.StatusCode() == 404 {
			return nil, ackerr.NotFound
		}
		if awsErr, ok := ackerr.AWSError(err); ok && awsErr.Code() == "ResourceNotFoundException" {
			return nil, ackerr.NotFound
		}
		return nil, err
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if resp.FirewallPolicy != nil {
		f0 := &svcapitypes.FirewallPolicy_SDK{}
		if resp.FirewallPolicy.StatefulDefaultActions != nil {
			f0f0 := []*string{}
			for _, f0f0iter := range resp.FirewallPolicy.StatefulDefaultActions {
				var f0f0elem string
				f0f0elem = *f0f0iter
				f0f0 = append(f0f0, &f0f0elem)
			}
			f0.StatefulDefaultActions = f0f0
		}
		if resp.FirewallPolicy.StatefulEngineOptions != nil {
			f0f1 := &svcapitypes.StatefulEngineOptions{}
			if resp.FirewallPolicy.StatefulEngineOptions.RuleOrder != nil {
				f0f1.RuleOrder = resp.FirewallPolicy.StatefulEngineOptions.RuleOrder
			}
			f0.StatefulEngineOptions = f0f1
		}
		if resp.FirewallPolicy.StatefulRuleGroupReferences != nil {
			f0f2 := []*svcapitypes.StatefulRuleGroupReference{}
			for _, f0f2iter := range resp.FirewallPolicy.StatefulRuleGroupReferences {
				f0f2elem := &svcapitypes.StatefulRuleGroupReference{}
				if f0f2iter.Override != nil {
					f0f2elemf0 := &svcapitypes.StatefulRuleGroupOverride{}
					if f0f2iter.Override.Action != nil {
						f0f2elemf0.Action = f0f2iter.Override.Action
					}
					f0f2elem.Override = f0f2elemf0
				}
				if f0f2iter.Priority != nil {
					f0f2elem.Priority = f0f2iter.Priority
				}
				if f0f2iter.ResourceArn != nil {
					f0f2elem.ResourceARN = f0f2iter.ResourceArn
				}
				f0f2 = append(f0f2, f0f2elem)
			}
			f0.StatefulRuleGroupReferences = f0f2
		}
		if resp.FirewallPolicy.StatelessCustomActions != nil {
			f0f3 := []*svcapitypes.CustomAction{}
			for _, f0f3iter := range resp.FirewallPolicy.StatelessCustomActions {
				f0f3elem := &svcapitypes.CustomAction{}
				if f0f3iter.ActionDefinition != nil {
					f0f3elemf0 := &svcapitypes.ActionDefinition{}
					if f0f3iter.ActionDefinition.PublishMetricAction != nil {
						f0f3elemf0f0 := &svcapitypes.PublishMetricAction{}
						if f0f3iter.ActionDefinition.PublishMetricAction.Dimensions != nil {
							f0f3elemf0f0f0 := []*svcapitypes.Dimension{}
							for _, f0f3elemf0f0f0iter := range f0f3iter.ActionDefinition.PublishMetricAction.Dimensions {
								f0f3elemf0f0f0elem := &svcapitypes.Dimension{}
								if f0f3elemf0f0f0iter.Value != nil {
									f0f3elemf0f0f0elem.Value = f0f3elemf0f0f0iter.Value
								}
								f0f3elemf0f0f0 = append(f0f3elemf0f0f0, f0f3elemf0f0f0elem)
							}
							f0f3elemf0f0.Dimensions = f0f3elemf0f0f0
						}
						f0f3elemf0.PublishMetricAction = f0f3elemf0f0
					}
					f0f3elem.ActionDefinition = f0f3elemf0
				}
				if f0f3iter.ActionName != nil {
					f0f3elem.ActionName = f0f3iter.ActionName
				}
				f0f3 = append(f0f3, f0f3elem)
			}
			f0.StatelessCustomActions = f0f3
		}
		if resp.FirewallPolicy.StatelessDefaultActions != nil {
			f0f4 := []*string{}
			for _, f0f4iter := range resp.FirewallPolicy.StatelessDefaultActions {
				var f0f4elem string
				f0f4elem = *f0f4iter
				f0f4 = append(f0f4, &f0f4elem)
			}
			f0.StatelessDefaultActions = f0f4
		}
		if resp.FirewallPolicy.StatelessFragmentDefaultActions != nil {
			f0f5 := []*string{}
			for _, f0f5iter := range resp.FirewallPolicy.StatelessFragmentDefaultActions {
				var f0f5elem string
				f0f5elem = *f0f5iter
				f0f5 = append(f0f5, &f0f5elem)
			}
			f0.StatelessFragmentDefaultActions = f0f5
		}
		if resp.FirewallPolicy.StatelessRuleGroupReferences != nil {
			f0f6 := []*svcapitypes.StatelessRuleGroupReference{}
			for _, f0f6iter := range resp.FirewallPolicy.StatelessRuleGroupReferences {
				f0f6elem := &svcapitypes.StatelessRuleGroupReference{}
				if f0f6iter.Priority != nil {
					f0f6elem.Priority = f0f6iter.Priority
				}
				if f0f6iter.ResourceArn != nil {
					f0f6elem.ResourceARN = f0f6iter.ResourceArn
				}
				f0f6 = append(f0f6, f0f6elem)
			}
			f0.StatelessRuleGroupReferences = f0f6
		}
		ko.Spec.FirewallPolicy = f0
	} else {
		ko.Spec.FirewallPolicy = nil
	}
	if resp.FirewallPolicyResponse != nil {
		f1 := &svcapitypes.FirewallPolicyResponse{}
		if resp.FirewallPolicyResponse.ConsumedStatefulRuleCapacity != nil {
			f1.ConsumedStatefulRuleCapacity = resp.FirewallPolicyResponse.ConsumedStatefulRuleCapacity
		}
		if resp.FirewallPolicyResponse.ConsumedStatelessRuleCapacity != nil {
			f1.ConsumedStatelessRuleCapacity = resp.FirewallPolicyResponse.ConsumedStatelessRuleCapacity
		}
		if resp.FirewallPolicyResponse.Description != nil {
			f1.Description = resp.FirewallPolicyResponse.Description
		}
		if resp.FirewallPolicyResponse.EncryptionConfiguration != nil {
			f1f3 := &svcapitypes.EncryptionConfiguration{}
			if resp.FirewallPolicyResponse.EncryptionConfiguration.KeyId != nil {
				f1f3.KeyID = resp.FirewallPolicyResponse.EncryptionConfiguration.KeyId
			}
			if resp.FirewallPolicyResponse.EncryptionConfiguration.Type != nil {
				f1f3.Type = resp.FirewallPolicyResponse.EncryptionConfiguration.Type
			}
			f1.EncryptionConfiguration = f1f3
		}
		if resp.FirewallPolicyResponse.FirewallPolicyArn != nil {
			f1.FirewallPolicyARN = resp.FirewallPolicyResponse.FirewallPolicyArn
		}
		if resp.FirewallPolicyResponse.FirewallPolicyId != nil {
			f1.FirewallPolicyID = resp.FirewallPolicyResponse.FirewallPolicyId
		}
		if resp.FirewallPolicyResponse.FirewallPolicyName != nil {
			f1.FirewallPolicyName = resp.FirewallPolicyResponse.FirewallPolicyName
		}
		if resp.FirewallPolicyResponse.FirewallPolicyStatus != nil {
			f1.FirewallPolicyStatus = resp.FirewallPolicyResponse.FirewallPolicyStatus
		}
		if resp.FirewallPolicyResponse.LastModifiedTime != nil {
			f1.LastModifiedTime = &metav1.Time{*resp.FirewallPolicyResponse.LastModifiedTime}
		}
		if resp.FirewallPolicyResponse.NumberOfAssociations != nil {
			f1.NumberOfAssociations = resp.FirewallPolicyResponse.NumberOfAssociations
		}
		if resp.FirewallPolicyResponse.Tags != nil {
			f1f10 := []*svcapitypes.Tag{}
			for _, f1f10iter := range resp.FirewallPolicyResponse.Tags {
				f1f10elem := &svcapitypes.Tag{}
				if f1f10iter.Key != nil {
					f1f10elem.Key = f1f10iter.Key
				}
				if f1f10iter.Value != nil {
					f1f10elem.Value = f1f10iter.Value
				}
				f1f10 = append(f1f10, f1f10elem)
			}
			f1.Tags = f1f10
		}
		ko.Status.FirewallPolicyResponse = f1
	} else {
		ko.Status.FirewallPolicyResponse = nil
	}
	if resp.UpdateToken != nil {
		ko.Status.UpdateToken = resp.UpdateToken
	} else {
		ko.Status.UpdateToken = nil
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// requiredFieldsMissingFromReadOneInput returns true if there are any fields
// for the ReadOne Input shape that are required but not present in the
// resource's Spec or Status
func (rm *resourceManager) requiredFieldsMissingFromReadOneInput(
	r *resource,
) bool {
	return false
}

// newDescribeRequestPayload returns SDK-specific struct for the HTTP request
// payload of the Describe API call for the resource
func (rm *resourceManager) newDescribeRequestPayload(
	r *resource,
) (*svcsdk.DescribeFirewallPolicyInput, error) {
	res := &svcsdk.DescribeFirewallPolicyInput{}

	if r.ko.Status.ACKResourceMetadata != nil && r.ko.Status.ACKResourceMetadata.ARN != nil {
		res.SetFirewallPolicyArn(string(*r.ko.Status.ACKResourceMetadata.ARN))
	}
	if r.ko.Spec.FirewallPolicyName != nil {
		res.SetFirewallPolicyName(*r.ko.Spec.FirewallPolicyName)
	}

	return res, nil
}

// sdkCreate creates the supplied resource in the backend AWS service API and
// returns a copy of the resource with resource fields (in both Spec and
// Status) filled in with values from the CREATE API operation's Output shape.
func (rm *resourceManager) sdkCreate(
	ctx context.Context,
	desired *resource,
) (created *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkCreate")
	defer func() {
		exit(err)
	}()
	input, err := rm.newCreateRequestPayload(ctx, desired)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.CreateFirewallPolicyOutput
	_ = resp
	resp, err = rm.sdkapi.CreateFirewallPolicyWithContext(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "CreateFirewallPolicy", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if resp.FirewallPolicyResponse != nil {
		f0 := &svcapitypes.FirewallPolicyResponse{}
		if resp.FirewallPolicyResponse.ConsumedStatefulRuleCapacity != nil {
			f0.ConsumedStatefulRuleCapacity = resp.FirewallPolicyResponse.ConsumedStatefulRuleCapacity
		}
		if resp.FirewallPolicyResponse.ConsumedStatelessRuleCapacity != nil {
			f0.ConsumedStatelessRuleCapacity = resp.FirewallPolicyResponse.ConsumedStatelessRuleCapacity
		}
		if resp.FirewallPolicyResponse.Description != nil {
			f0.Description = resp.FirewallPolicyResponse.Description
		}
		if resp.FirewallPolicyResponse.EncryptionConfiguration != nil {
			f0f3 := &svcapitypes.EncryptionConfiguration{}
			if resp.FirewallPolicyResponse.EncryptionConfiguration.KeyId != nil {
				f0f3.KeyID = resp.FirewallPolicyResponse.EncryptionConfiguration.KeyId
			}
			if resp.FirewallPolicyResponse.EncryptionConfiguration.Type != nil {
				f0f3.Type = resp.FirewallPolicyResponse.EncryptionConfiguration.Type
			}
			f0.EncryptionConfiguration = f0f3
		}
		if resp.FirewallPolicyResponse.FirewallPolicyArn != nil {
			f0.FirewallPolicyARN = resp.FirewallPolicyResponse.FirewallPolicyArn
		}
		if resp.FirewallPolicyResponse.FirewallPolicyId != nil {
			f0.FirewallPolicyID = resp.FirewallPolicyResponse.FirewallPolicyId
		}
		if resp.FirewallPolicyResponse.FirewallPolicyName != nil {
			f0.FirewallPolicyName = resp.FirewallPolicyResponse.FirewallPolicyName
		}
		if resp.FirewallPolicyResponse.FirewallPolicyStatus != nil {
			f0.FirewallPolicyStatus = resp.FirewallPolicyResponse.FirewallPolicyStatus
		}
		if resp.FirewallPolicyResponse.LastModifiedTime != nil {
			f0.LastModifiedTime = &metav1.Time{*resp.FirewallPolicyResponse.LastModifiedTime}
		}
		if resp.FirewallPolicyResponse.NumberOfAssociations != nil {
			f0.NumberOfAssociations = resp.FirewallPolicyResponse.NumberOfAssociations
		}
		if resp.FirewallPolicyResponse.Tags != nil {
			f0f10 := []*svcapitypes.Tag{}
			for _, f0f10iter := range resp.FirewallPolicyResponse.Tags {
				f0f10elem := &svcapitypes.Tag{}
				if f0f10iter.Key != nil {
					f0f10elem.Key = f0f10iter.Key
				}
				if f0f10iter.Value != nil {
					f0f10elem.Value = f0f10iter.Value
				}
				f0f10 = append(f0f10, f0f10elem)
			}
			f0.Tags = f0f10
		}
		ko.Status.FirewallPolicyResponse = f0
	} else {
		ko.Status.FirewallPolicyResponse = nil
	}
	if resp.UpdateToken != nil {
		ko.Status.UpdateToken = resp.UpdateToken
	} else {
		ko.Status.UpdateToken = nil
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.CreateFirewallPolicyInput, error) {
	res := &svcsdk.CreateFirewallPolicyInput{}

	if r.ko.Spec.Description != nil {
		res.SetDescription(*r.ko.Spec.Description)
	}
	if r.ko.Spec.EncryptionConfiguration != nil {
		f1 := &svcsdk.EncryptionConfiguration{}
		if r.ko.Spec.EncryptionConfiguration.KeyID != nil {
			f1.SetKeyId(*r.ko.Spec.EncryptionConfiguration.KeyID)
		}
		if r.ko.Spec.EncryptionConfiguration.Type != nil {
			f1.SetType(*r.ko.Spec.EncryptionConfiguration.Type)
		}
		res.SetEncryptionConfiguration(f1)
	}
	if r.ko.Spec.FirewallPolicy != nil {
		f2 := &svcsdk.FirewallPolicy{}
		if r.ko.Spec.FirewallPolicy.StatefulDefaultActions != nil {
			f2f0 := []*string{}
			for _, f2f0iter := range r.ko.Spec.FirewallPolicy.StatefulDefaultActions {
				var f2f0elem string
				f2f0elem = *f2f0iter
				f2f0 = append(f2f0, &f2f0elem)
			}
			f2.SetStatefulDefaultActions(f2f0)
		}
		if r.ko.Spec.FirewallPolicy.StatefulEngineOptions != nil {
			f2f1 := &svcsdk.StatefulEngineOptions{}
			if r.ko.Spec.FirewallPolicy.StatefulEngineOptions.RuleOrder != nil {
				f2f1.SetRuleOrder(*r.ko.Spec.FirewallPolicy.StatefulEngineOptions.RuleOrder)
			}
			f2.SetStatefulEngineOptions(f2f1)
		}
		if r.ko.Spec.FirewallPolicy.StatefulRuleGroupReferences != nil {
			f2f2 := []*svcsdk.StatefulRuleGroupReference{}
			for _, f2f2iter := range r.ko.Spec.FirewallPolicy.StatefulRuleGroupReferences {
				f2f2elem := &svcsdk.StatefulRuleGroupReference{}
				if f2f2iter.Override != nil {
					f2f2elemf0 := &svcsdk.StatefulRuleGroupOverride{}
					if f2f2iter.Override.Action != nil {
						f2f2elemf0.SetAction(*f2f2iter.Override.Action)
					}
					f2f2elem.SetOverride(f2f2elemf0)
				}
				if f2f2iter.Priority != nil {
					f2f2elem.SetPriority(*f2f2iter.Priority)
				}
				if f2f2iter.ResourceARN != nil {
					f2f2elem.SetResourceArn(*f2f2iter.ResourceARN)
				}
				f2f2 = append(f2f2, f2f2elem)
			}
			f2.SetStatefulRuleGroupReferences(f2f2)
		}
		if r.ko.Spec.FirewallPolicy.StatelessCustomActions != nil {
			f2f3 := []*svcsdk.CustomAction{}
			for _, f2f3iter := range r.ko.Spec.FirewallPolicy.StatelessCustomActions {
				f2f3elem := &svcsdk.CustomAction{}
				if f2f3iter.ActionDefinition != nil {
					f2f3elemf0 := &svcsdk.ActionDefinition{}
					if f2f3iter.ActionDefinition.PublishMetricAction != nil {
						f2f3elemf0f0 := &svcsdk.PublishMetricAction{}
						if f2f3iter.ActionDefinition.PublishMetricAction.Dimensions != nil {
							f2f3elemf0f0f0 := []*svcsdk.Dimension{}
							for _, f2f3elemf0f0f0iter := range f2f3iter.ActionDefinition.PublishMetricAction.Dimensions {
								f2f3elemf0f0f0elem := &svcsdk.Dimension{}
								if f2f3elemf0f0f0iter.Value != nil {
									f2f3elemf0f0f0elem.SetValue(*f2f3elemf0f0f0iter.Value)
								}
								f2f3elemf0f0f0 = append(f2f3elemf0f0f0, f2f3elemf0f0f0elem)
							}
							f2f3elemf0f0.SetDimensions(f2f3elemf0f0f0)
						}
						f2f3elemf0.SetPublishMetricAction(f2f3elemf0f0)
					}
					f2f3elem.SetActionDefinition(f2f3elemf0)
				}
				if f2f3iter.ActionName != nil {
					f2f3elem.SetActionName(*f2f3iter.ActionName)
				}
				f2f3 = append(f2f3, f2f3elem)
			}
			f2.SetStatelessCustomActions(f2f3)
		}
		if r.ko.Spec.FirewallPolicy.StatelessDefaultActions != nil {
			f2f4 := []*string{}
			for _, f2f4iter := range r.ko.Spec.FirewallPolicy.StatelessDefaultActions {
				var f2f4elem string
				f2f4elem = *f2f4iter
				f2f4 = append(f2f4, &f2f4elem)
			}
			f2.SetStatelessDefaultActions(f2f4)
		}
		if r.ko.Spec.FirewallPolicy.StatelessFragmentDefaultActions != nil {
			f2f5 := []*string{}
			for _, f2f5iter := range r.ko.Spec.FirewallPolicy.StatelessFragmentDefaultActions {
				var f2f5elem string
				f2f5elem = *f2f5iter
				f2f5 = append(f2f5, &f2f5elem)
			}
			f2.SetStatelessFragmentDefaultActions(f2f5)
		}
		if r.ko.Spec.FirewallPolicy.StatelessRuleGroupReferences != nil {
			f2f6 := []*svcsdk.StatelessRuleGroupReference{}
			for _, f2f6iter := range r.ko.Spec.FirewallPolicy.StatelessRuleGroupReferences {
				f2f6elem := &svcsdk.StatelessRuleGroupReference{}
				if f2f6iter.Priority != nil {
					f2f6elem.SetPriority(*f2f6iter.Priority)
				}
				if f2f6iter.ResourceARN != nil {
					f2f6elem.SetResourceArn(*f2f6iter.ResourceARN)
				}
				f2f6 = append(f2f6, f2f6elem)
			}
			f2.SetStatelessRuleGroupReferences(f2f6)
		}
		res.SetFirewallPolicy(f2)
	}
	if r.ko.Spec.FirewallPolicyName != nil {
		res.SetFirewallPolicyName(*r.ko.Spec.FirewallPolicyName)
	}
	if r.ko.Spec.Tags != nil {
		f4 := []*svcsdk.Tag{}
		for _, f4iter := range r.ko.Spec.Tags {
			f4elem := &svcsdk.Tag{}
			if f4iter.Key != nil {
				f4elem.SetKey(*f4iter.Key)
			}
			if f4iter.Value != nil {
				f4elem.SetValue(*f4iter.Value)
			}
			f4 = append(f4, f4elem)
		}
		res.SetTags(f4)
	}

	return res, nil
}

// sdkUpdate patches the supplied resource in the backend AWS service API and
// returns a new resource with updated fields.
func (rm *resourceManager) sdkUpdate(
	ctx context.Context,
	desired *resource,
	latest *resource,
	delta *ackcompare.Delta,
) (updated *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkUpdate")
	defer func() {
		exit(err)
	}()
	input, err := rm.newUpdateRequestPayload(ctx, desired, delta)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.UpdateFirewallPolicyOutput
	_ = resp
	resp, err = rm.sdkapi.UpdateFirewallPolicyWithContext(ctx, input)
	rm.metrics.RecordAPICall("UPDATE", "UpdateFirewallPolicy", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if resp.FirewallPolicyResponse != nil {
		f0 := &svcapitypes.FirewallPolicyResponse{}
		if resp.FirewallPolicyResponse.ConsumedStatefulRuleCapacity != nil {
			f0.ConsumedStatefulRuleCapacity = resp.FirewallPolicyResponse.ConsumedStatefulRuleCapacity
		}
		if resp.FirewallPolicyResponse.ConsumedStatelessRuleCapacity != nil {
			f0.ConsumedStatelessRuleCapacity = resp.FirewallPolicyResponse.ConsumedStatelessRuleCapacity
		}
		if resp.FirewallPolicyResponse.Description != nil {
			f0.Description = resp.FirewallPolicyResponse.Description
		}
		if resp.FirewallPolicyResponse.EncryptionConfiguration != nil {
			f0f3 := &svcapitypes.EncryptionConfiguration{}
			if resp.FirewallPolicyResponse.EncryptionConfiguration.KeyId != nil {
				f0f3.KeyID = resp.FirewallPolicyResponse.EncryptionConfiguration.KeyId
			}
			if resp.FirewallPolicyResponse.EncryptionConfiguration.Type != nil {
				f0f3.Type = resp.FirewallPolicyResponse.EncryptionConfiguration.Type
			}
			f0.EncryptionConfiguration = f0f3
		}
		if resp.FirewallPolicyResponse.FirewallPolicyArn != nil {
			f0.FirewallPolicyARN = resp.FirewallPolicyResponse.FirewallPolicyArn
		}
		if resp.FirewallPolicyResponse.FirewallPolicyId != nil {
			f0.FirewallPolicyID = resp.FirewallPolicyResponse.FirewallPolicyId
		}
		if resp.FirewallPolicyResponse.FirewallPolicyName != nil {
			f0.FirewallPolicyName = resp.FirewallPolicyResponse.FirewallPolicyName
		}
		if resp.FirewallPolicyResponse.FirewallPolicyStatus != nil {
			f0.FirewallPolicyStatus = resp.FirewallPolicyResponse.FirewallPolicyStatus
		}
		if resp.FirewallPolicyResponse.LastModifiedTime != nil {
			f0.LastModifiedTime = &metav1.Time{*resp.FirewallPolicyResponse.LastModifiedTime}
		}
		if resp.FirewallPolicyResponse.NumberOfAssociations != nil {
			f0.NumberOfAssociations = resp.FirewallPolicyResponse.NumberOfAssociations
		}
		if resp.FirewallPolicyResponse.Tags != nil {
			f0f10 := []*svcapitypes.Tag{}
			for _, f0f10iter := range resp.FirewallPolicyResponse.Tags {
				f0f10elem := &svcapitypes.Tag{}
				if f0f10iter.Key != nil {
					f0f10elem.Key = f0f10iter.Key
				}
				if f0f10iter.Value != nil {
					f0f10elem.Value = f0f10iter.Value
				}
				f0f10 = append(f0f10, f0f10elem)
			}
			f0.Tags = f0f10
		}
		ko.Status.FirewallPolicyResponse = f0
	} else {
		ko.Status.FirewallPolicyResponse = nil
	}
	if resp.UpdateToken != nil {
		ko.Status.UpdateToken = resp.UpdateToken
	} else {
		ko.Status.UpdateToken = nil
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// newUpdateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Update API call for the resource
func (rm *resourceManager) newUpdateRequestPayload(
	ctx context.Context,
	r *resource,
	delta *ackcompare.Delta,
) (*svcsdk.UpdateFirewallPolicyInput, error) {
	res := &svcsdk.UpdateFirewallPolicyInput{}

	if r.ko.Spec.Description != nil {
		res.SetDescription(*r.ko.Spec.Description)
	}
	if r.ko.Spec.EncryptionConfiguration != nil {
		f2 := &svcsdk.EncryptionConfiguration{}
		if r.ko.Spec.EncryptionConfiguration.KeyID != nil {
			f2.SetKeyId(*r.ko.Spec.EncryptionConfiguration.KeyID)
		}
		if r.ko.Spec.EncryptionConfiguration.Type != nil {
			f2.SetType(*r.ko.Spec.EncryptionConfiguration.Type)
		}
		res.SetEncryptionConfiguration(f2)
	}
	if r.ko.Spec.FirewallPolicy != nil {
		f3 := &svcsdk.FirewallPolicy{}
		if r.ko.Spec.FirewallPolicy.StatefulDefaultActions != nil {
			f3f0 := []*string{}
			for _, f3f0iter := range r.ko.Spec.FirewallPolicy.StatefulDefaultActions {
				var f3f0elem string
				f3f0elem = *f3f0iter
				f3f0 = append(f3f0, &f3f0elem)
			}
			f3.SetStatefulDefaultActions(f3f0)
		}
		if r.ko.Spec.FirewallPolicy.StatefulEngineOptions != nil {
			f3f1 := &svcsdk.StatefulEngineOptions{}
			if r.ko.Spec.FirewallPolicy.StatefulEngineOptions.RuleOrder != nil {
				f3f1.SetRuleOrder(*r.ko.Spec.FirewallPolicy.StatefulEngineOptions.RuleOrder)
			}
			f3.SetStatefulEngineOptions(f3f1)
		}
		if r.ko.Spec.FirewallPolicy.StatefulRuleGroupReferences != nil {
			f3f2 := []*svcsdk.StatefulRuleGroupReference{}
			for _, f3f2iter := range r.ko.Spec.FirewallPolicy.StatefulRuleGroupReferences {
				f3f2elem := &svcsdk.StatefulRuleGroupReference{}
				if f3f2iter.Override != nil {
					f3f2elemf0 := &svcsdk.StatefulRuleGroupOverride{}
					if f3f2iter.Override.Action != nil {
						f3f2elemf0.SetAction(*f3f2iter.Override.Action)
					}
					f3f2elem.SetOverride(f3f2elemf0)
				}
				if f3f2iter.Priority != nil {
					f3f2elem.SetPriority(*f3f2iter.Priority)
				}
				if f3f2iter.ResourceARN != nil {
					f3f2elem.SetResourceArn(*f3f2iter.ResourceARN)
				}
				f3f2 = append(f3f2, f3f2elem)
			}
			f3.SetStatefulRuleGroupReferences(f3f2)
		}
		if r.ko.Spec.FirewallPolicy.StatelessCustomActions != nil {
			f3f3 := []*svcsdk.CustomAction{}
			for _, f3f3iter := range r.ko.Spec.FirewallPolicy.StatelessCustomActions {
				f3f3elem := &svcsdk.CustomAction{}
				if f3f3iter.ActionDefinition != nil {
					f3f3elemf0 := &svcsdk.ActionDefinition{}
					if f3f3iter.ActionDefinition.PublishMetricAction != nil {
						f3f3elemf0f0 := &svcsdk.PublishMetricAction{}
						if f3f3iter.ActionDefinition.PublishMetricAction.Dimensions != nil {
							f3f3elemf0f0f0 := []*svcsdk.Dimension{}
							for _, f3f3elemf0f0f0iter := range f3f3iter.ActionDefinition.PublishMetricAction.Dimensions {
								f3f3elemf0f0f0elem := &svcsdk.Dimension{}
								if f3f3elemf0f0f0iter.Value != nil {
									f3f3elemf0f0f0elem.SetValue(*f3f3elemf0f0f0iter.Value)
								}
								f3f3elemf0f0f0 = append(f3f3elemf0f0f0, f3f3elemf0f0f0elem)
							}
							f3f3elemf0f0.SetDimensions(f3f3elemf0f0f0)
						}
						f3f3elemf0.SetPublishMetricAction(f3f3elemf0f0)
					}
					f3f3elem.SetActionDefinition(f3f3elemf0)
				}
				if f3f3iter.ActionName != nil {
					f3f3elem.SetActionName(*f3f3iter.ActionName)
				}
				f3f3 = append(f3f3, f3f3elem)
			}
			f3.SetStatelessCustomActions(f3f3)
		}
		if r.ko.Spec.FirewallPolicy.StatelessDefaultActions != nil {
			f3f4 := []*string{}
			for _, f3f4iter := range r.ko.Spec.FirewallPolicy.StatelessDefaultActions {
				var f3f4elem string
				f3f4elem = *f3f4iter
				f3f4 = append(f3f4, &f3f4elem)
			}
			f3.SetStatelessDefaultActions(f3f4)
		}
		if r.ko.Spec.FirewallPolicy.StatelessFragmentDefaultActions != nil {
			f3f5 := []*string{}
			for _, f3f5iter := range r.ko.Spec.FirewallPolicy.StatelessFragmentDefaultActions {
				var f3f5elem string
				f3f5elem = *f3f5iter
				f3f5 = append(f3f5, &f3f5elem)
			}
			f3.SetStatelessFragmentDefaultActions(f3f5)
		}
		if r.ko.Spec.FirewallPolicy.StatelessRuleGroupReferences != nil {
			f3f6 := []*svcsdk.StatelessRuleGroupReference{}
			for _, f3f6iter := range r.ko.Spec.FirewallPolicy.StatelessRuleGroupReferences {
				f3f6elem := &svcsdk.StatelessRuleGroupReference{}
				if f3f6iter.Priority != nil {
					f3f6elem.SetPriority(*f3f6iter.Priority)
				}
				if f3f6iter.ResourceARN != nil {
					f3f6elem.SetResourceArn(*f3f6iter.ResourceARN)
				}
				f3f6 = append(f3f6, f3f6elem)
			}
			f3.SetStatelessRuleGroupReferences(f3f6)
		}
		res.SetFirewallPolicy(f3)
	}
	if r.ko.Status.ACKResourceMetadata != nil && r.ko.Status.ACKResourceMetadata.ARN != nil {
		res.SetFirewallPolicyArn(string(*r.ko.Status.ACKResourceMetadata.ARN))
	}
	if r.ko.Spec.FirewallPolicyName != nil {
		res.SetFirewallPolicyName(*r.ko.Spec.FirewallPolicyName)
	}
	if r.ko.Status.UpdateToken != nil {
		res.SetUpdateToken(*r.ko.Status.UpdateToken)
	}

	return res, nil
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkDelete")
	defer func() {
		exit(err)
	}()
	input, err := rm.newDeleteRequestPayload(r)
	if err != nil {
		return nil, err
	}
	var resp *svcsdk.DeleteFirewallPolicyOutput
	_ = resp
	resp, err = rm.sdkapi.DeleteFirewallPolicyWithContext(ctx, input)
	rm.metrics.RecordAPICall("DELETE", "DeleteFirewallPolicy", err)
	return nil, err
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteFirewallPolicyInput, error) {
	res := &svcsdk.DeleteFirewallPolicyInput{}

	if r.ko.Status.ACKResourceMetadata != nil && r.ko.Status.ACKResourceMetadata.ARN != nil {
		res.SetFirewallPolicyArn(string(*r.ko.Status.ACKResourceMetadata.ARN))
	}
	if r.ko.Spec.FirewallPolicyName != nil {
		res.SetFirewallPolicyName(*r.ko.Spec.FirewallPolicyName)
	}

	return res, nil
}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.FirewallPolicy,
) {
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if ko.Status.ACKResourceMetadata.Region == nil {
		ko.Status.ACKResourceMetadata.Region = &rm.awsRegion
	}
	if ko.Status.ACKResourceMetadata.OwnerAccountID == nil {
		ko.Status.ACKResourceMetadata.OwnerAccountID = &rm.awsAccountID
	}
	if ko.Status.Conditions == nil {
		ko.Status.Conditions = []*ackv1alpha1.Condition{}
	}
}

// updateConditions returns updated resource, true; if conditions were updated
// else it returns nil, false
func (rm *resourceManager) updateConditions(
	r *resource,
	onSuccess bool,
	err error,
) (*resource, bool) {
	ko := r.ko.DeepCopy()
	rm.setStatusDefaults(ko)

	// Terminal condition
	var terminalCondition *ackv1alpha1.Condition = nil
	var recoverableCondition *ackv1alpha1.Condition = nil
	var syncCondition *ackv1alpha1.Condition = nil
	for _, condition := range ko.Status.Conditions {
		if condition.Type == ackv1alpha1.ConditionTypeTerminal {
			terminalCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeRecoverable {
			recoverableCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeResourceSynced {
			syncCondition = condition
		}
	}
	var termError *ackerr.TerminalError
	if rm.terminalAWSError(err) || err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound || errors.As(err, &termError) {
		if terminalCondition == nil {
			terminalCondition = &ackv1alpha1.Condition{
				Type: ackv1alpha1.ConditionTypeTerminal,
			}
			ko.Status.Conditions = append(ko.Status.Conditions, terminalCondition)
		}
		var errorMessage = ""
		if err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound || errors.As(err, &termError) {
			errorMessage = err.Error()
		} else {
			awsErr, _ := ackerr.AWSError(err)
			errorMessage = awsErr.Error()
		}
		terminalCondition.Status = corev1.ConditionTrue
		terminalCondition.Message = &errorMessage
	} else {
		// Clear the terminal condition if no longer present
		if terminalCondition != nil {
			terminalCondition.Status = corev1.ConditionFalse
			terminalCondition.Message = nil
		}
		// Handling Recoverable Conditions
		if err != nil {
			if recoverableCondition == nil {
				// Add a new Condition containing a non-terminal error
				recoverableCondition = &ackv1alpha1.Condition{
					Type: ackv1alpha1.ConditionTypeRecoverable,
				}
				ko.Status.Conditions = append(ko.Status.Conditions, recoverableCondition)
			}
			recoverableCondition.Status = corev1.ConditionTrue
			awsErr, _ := ackerr.AWSError(err)
			errorMessage := err.Error()
			if awsErr != nil {
				errorMessage = awsErr.Error()
			}
			recoverableCondition.Message = &errorMessage
		} else if recoverableCondition != nil {
			recoverableCondition.Status = corev1.ConditionFalse
			recoverableCondition.Message = nil
		}
	}
	// Required to avoid the "declared but not used" error in the default case
	_ = syncCondition
	if terminalCondition != nil || recoverableCondition != nil || syncCondition != nil {
		return &resource{ko}, true // updated
	}
	return nil, false // not updated
}

// terminalAWSError returns awserr, true; if the supplied error is an aws Error type
// and if the exception indicates that it is a Terminal exception
// 'Terminal' exception are specified in generator configuration
func (rm *resourceManager) terminalAWSError(err error) bool {
	if err == nil {
		return false
	}
	awsErr, ok := ackerr.AWSError(err)
	if !ok {
		return false
	}
	switch awsErr.Code() {
	case "InvalidRequestException":
		return true
	default:
		return false
	}
}
