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

package apigateway

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgiamresource"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const ApiGatewayApiIAMAssetType string = "apigateway.googleapis.com/Api"

func ResourceConverterApiGatewayApiIamPolicy() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType:         ApiGatewayApiIAMAssetType,
		Convert:           GetApiGatewayApiIamPolicyCaiObject,
		MergeCreateUpdate: MergeApiGatewayApiIamPolicy,
	}
}

func ResourceConverterApiGatewayApiIamBinding() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType:         ApiGatewayApiIAMAssetType,
		Convert:           GetApiGatewayApiIamBindingCaiObject,
		FetchFullResource: FetchApiGatewayApiIamPolicy,
		MergeCreateUpdate: MergeApiGatewayApiIamBinding,
		MergeDelete:       MergeApiGatewayApiIamBindingDelete,
	}
}

func ResourceConverterApiGatewayApiIamMember() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType:         ApiGatewayApiIAMAssetType,
		Convert:           GetApiGatewayApiIamMemberCaiObject,
		FetchFullResource: FetchApiGatewayApiIamPolicy,
		MergeCreateUpdate: MergeApiGatewayApiIamMember,
		MergeDelete:       MergeApiGatewayApiIamMemberDelete,
	}
}

func GetApiGatewayApiIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	return newApiGatewayApiIamAsset(d, config, tpgiamresource.ExpandIamPolicyBindings)
}

func GetApiGatewayApiIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	return newApiGatewayApiIamAsset(d, config, tpgiamresource.ExpandIamRoleBindings)
}

func GetApiGatewayApiIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	return newApiGatewayApiIamAsset(d, config, tpgiamresource.ExpandIamMemberBindings)
}

func MergeApiGatewayApiIamPolicy(existing, incoming tpgresource.Asset) tpgresource.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeApiGatewayApiIamBinding(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeIamAssets(existing, incoming, tpgiamresource.MergeAuthoritativeBindings)
}

func MergeApiGatewayApiIamBindingDelete(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeDeleteIamAssets(existing, incoming, tpgiamresource.MergeDeleteAuthoritativeBindings)
}

func MergeApiGatewayApiIamMember(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeIamAssets(existing, incoming, tpgiamresource.MergeAdditiveBindings)
}

func MergeApiGatewayApiIamMemberDelete(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeDeleteIamAssets(existing, incoming, tpgiamresource.MergeDeleteAdditiveBindings)
}

func newApiGatewayApiIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]tpgresource.IAMBinding, error),
) ([]tpgresource.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []tpgresource.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := tpgresource.AssetName(d, config, "//apigateway.googleapis.com/projects/{{project}}/locations/global/apis/{{api}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}

	return []tpgresource.Asset{{
		Name: name,
		Type: ApiGatewayApiIAMAssetType,
		IAMPolicy: &tpgresource.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchApiGatewayApiIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (tpgresource.Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("api"); !ok {
		return tpgresource.Asset{}, tpgresource.ErrEmptyIdentityField
	}

	return tpgiamresource.FetchIamPolicy(
		ApiGatewayApiIamUpdaterProducer,
		d,
		config,
		"//apigateway.googleapis.com/projects/{{project}}/locations/global/apis/{{api}}",
		ApiGatewayApiIAMAssetType,
	)
}