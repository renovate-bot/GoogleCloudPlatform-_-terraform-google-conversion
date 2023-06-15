// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package bigquerydatapolicy

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgiamresource"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const BigqueryDatapolicyDataPolicyIAMAssetType string = "bigquerydatapolicy.googleapis.com/DataPolicy"

func ResourceConverterBigqueryDatapolicyDataPolicyIamPolicy() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType:         BigqueryDatapolicyDataPolicyIAMAssetType,
		Convert:           GetBigqueryDatapolicyDataPolicyIamPolicyCaiObject,
		MergeCreateUpdate: MergeBigqueryDatapolicyDataPolicyIamPolicy,
	}
}

func ResourceConverterBigqueryDatapolicyDataPolicyIamBinding() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType:         BigqueryDatapolicyDataPolicyIAMAssetType,
		Convert:           GetBigqueryDatapolicyDataPolicyIamBindingCaiObject,
		FetchFullResource: FetchBigqueryDatapolicyDataPolicyIamPolicy,
		MergeCreateUpdate: MergeBigqueryDatapolicyDataPolicyIamBinding,
		MergeDelete:       MergeBigqueryDatapolicyDataPolicyIamBindingDelete,
	}
}

func ResourceConverterBigqueryDatapolicyDataPolicyIamMember() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType:         BigqueryDatapolicyDataPolicyIAMAssetType,
		Convert:           GetBigqueryDatapolicyDataPolicyIamMemberCaiObject,
		FetchFullResource: FetchBigqueryDatapolicyDataPolicyIamPolicy,
		MergeCreateUpdate: MergeBigqueryDatapolicyDataPolicyIamMember,
		MergeDelete:       MergeBigqueryDatapolicyDataPolicyIamMemberDelete,
	}
}

func GetBigqueryDatapolicyDataPolicyIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	return newBigqueryDatapolicyDataPolicyIamAsset(d, config, tpgiamresource.ExpandIamPolicyBindings)
}

func GetBigqueryDatapolicyDataPolicyIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	return newBigqueryDatapolicyDataPolicyIamAsset(d, config, tpgiamresource.ExpandIamRoleBindings)
}

func GetBigqueryDatapolicyDataPolicyIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	return newBigqueryDatapolicyDataPolicyIamAsset(d, config, tpgiamresource.ExpandIamMemberBindings)
}

func MergeBigqueryDatapolicyDataPolicyIamPolicy(existing, incoming tpgresource.Asset) tpgresource.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeBigqueryDatapolicyDataPolicyIamBinding(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeIamAssets(existing, incoming, tpgiamresource.MergeAuthoritativeBindings)
}

func MergeBigqueryDatapolicyDataPolicyIamBindingDelete(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeDeleteIamAssets(existing, incoming, tpgiamresource.MergeDeleteAuthoritativeBindings)
}

func MergeBigqueryDatapolicyDataPolicyIamMember(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeIamAssets(existing, incoming, tpgiamresource.MergeAdditiveBindings)
}

func MergeBigqueryDatapolicyDataPolicyIamMemberDelete(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeDeleteIamAssets(existing, incoming, tpgiamresource.MergeDeleteAdditiveBindings)
}

func newBigqueryDatapolicyDataPolicyIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]tpgresource.IAMBinding, error),
) ([]tpgresource.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []tpgresource.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := tpgresource.AssetName(d, config, "//bigquerydatapolicy.googleapis.com/projects/{{project}}/locations/{{location}}/dataPolicies/{{data_policy_id}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}

	return []tpgresource.Asset{{
		Name: name,
		Type: BigqueryDatapolicyDataPolicyIAMAssetType,
		IAMPolicy: &tpgresource.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchBigqueryDatapolicyDataPolicyIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (tpgresource.Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("location"); !ok {
		return tpgresource.Asset{}, tpgresource.ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("data_policy_id"); !ok {
		return tpgresource.Asset{}, tpgresource.ErrEmptyIdentityField
	}

	return tpgiamresource.FetchIamPolicy(
		BigqueryDatapolicyDataPolicyIamUpdaterProducer,
		d,
		config,
		"//bigquerydatapolicy.googleapis.com/projects/{{project}}/locations/{{location}}/dataPolicies/{{data_policy_id}}",
		BigqueryDatapolicyDataPolicyIAMAssetType,
	)
}