// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/compute/BackendBucket.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter_iam.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package compute

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const ComputeBackendBucketIAMAssetType string = "compute.googleapis.com/BackendBucket"

func ResourceConverterComputeBackendBucketIamPolicy() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         ComputeBackendBucketIAMAssetType,
		Convert:           GetComputeBackendBucketIamPolicyCaiObject,
		MergeCreateUpdate: MergeComputeBackendBucketIamPolicy,
	}
}

func ResourceConverterComputeBackendBucketIamBinding() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         ComputeBackendBucketIAMAssetType,
		Convert:           GetComputeBackendBucketIamBindingCaiObject,
		FetchFullResource: FetchComputeBackendBucketIamPolicy,
		MergeCreateUpdate: MergeComputeBackendBucketIamBinding,
		MergeDelete:       MergeComputeBackendBucketIamBindingDelete,
	}
}

func ResourceConverterComputeBackendBucketIamMember() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         ComputeBackendBucketIAMAssetType,
		Convert:           GetComputeBackendBucketIamMemberCaiObject,
		FetchFullResource: FetchComputeBackendBucketIamPolicy,
		MergeCreateUpdate: MergeComputeBackendBucketIamMember,
		MergeDelete:       MergeComputeBackendBucketIamMemberDelete,
	}
}

func GetComputeBackendBucketIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newComputeBackendBucketIamAsset(d, config, cai.ExpandIamPolicyBindings)
}

func GetComputeBackendBucketIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newComputeBackendBucketIamAsset(d, config, cai.ExpandIamRoleBindings)
}

func GetComputeBackendBucketIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newComputeBackendBucketIamAsset(d, config, cai.ExpandIamMemberBindings)
}

func MergeComputeBackendBucketIamPolicy(existing, incoming cai.Asset) cai.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeComputeBackendBucketIamBinding(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAuthoritativeBindings)
}

func MergeComputeBackendBucketIamBindingDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAuthoritativeBindings)
}

func MergeComputeBackendBucketIamMember(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAdditiveBindings)
}

func MergeComputeBackendBucketIamMemberDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAdditiveBindings)
}

func newComputeBackendBucketIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]cai.IAMBinding, error),
) ([]cai.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []cai.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := cai.AssetName(d, config, "//compute.googleapis.com/projects/{{project}}/global/backendBuckets/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}

	return []cai.Asset{{
		Name: name,
		Type: ComputeBackendBucketIAMAssetType,
		IAMPolicy: &cai.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchComputeBackendBucketIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (cai.Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("name"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}

	return cai.FetchIamPolicy(
		ComputeBackendBucketIamUpdaterProducer,
		d,
		config,
		"//compute.googleapis.com/projects/{{project}}/global/backendBuckets/{{name}}",
		ComputeBackendBucketIAMAssetType,
	)
}
