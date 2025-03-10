// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/apigateway/ApiConfig.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter_iam.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package apigateway

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const ApiGatewayApiConfigIAMAssetType string = "apigateway.googleapis.com/ApiConfig"

func ResourceConverterApiGatewayApiConfigIamPolicy() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         ApiGatewayApiConfigIAMAssetType,
		Convert:           GetApiGatewayApiConfigIamPolicyCaiObject,
		MergeCreateUpdate: MergeApiGatewayApiConfigIamPolicy,
	}
}

func ResourceConverterApiGatewayApiConfigIamBinding() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         ApiGatewayApiConfigIAMAssetType,
		Convert:           GetApiGatewayApiConfigIamBindingCaiObject,
		FetchFullResource: FetchApiGatewayApiConfigIamPolicy,
		MergeCreateUpdate: MergeApiGatewayApiConfigIamBinding,
		MergeDelete:       MergeApiGatewayApiConfigIamBindingDelete,
	}
}

func ResourceConverterApiGatewayApiConfigIamMember() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         ApiGatewayApiConfigIAMAssetType,
		Convert:           GetApiGatewayApiConfigIamMemberCaiObject,
		FetchFullResource: FetchApiGatewayApiConfigIamPolicy,
		MergeCreateUpdate: MergeApiGatewayApiConfigIamMember,
		MergeDelete:       MergeApiGatewayApiConfigIamMemberDelete,
	}
}

func GetApiGatewayApiConfigIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newApiGatewayApiConfigIamAsset(d, config, cai.ExpandIamPolicyBindings)
}

func GetApiGatewayApiConfigIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newApiGatewayApiConfigIamAsset(d, config, cai.ExpandIamRoleBindings)
}

func GetApiGatewayApiConfigIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newApiGatewayApiConfigIamAsset(d, config, cai.ExpandIamMemberBindings)
}

func MergeApiGatewayApiConfigIamPolicy(existing, incoming cai.Asset) cai.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeApiGatewayApiConfigIamBinding(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAuthoritativeBindings)
}

func MergeApiGatewayApiConfigIamBindingDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAuthoritativeBindings)
}

func MergeApiGatewayApiConfigIamMember(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAdditiveBindings)
}

func MergeApiGatewayApiConfigIamMemberDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAdditiveBindings)
}

func newApiGatewayApiConfigIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]cai.IAMBinding, error),
) ([]cai.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []cai.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := cai.AssetName(d, config, "//apigateway.googleapis.com/projects/{{project}}/locations/global/apis/{{api}}/configs/{{api_config}}")
	if err != nil {
		return []cai.Asset{}, err
	}

	return []cai.Asset{{
		Name: name,
		Type: ApiGatewayApiConfigIAMAssetType,
		IAMPolicy: &cai.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchApiGatewayApiConfigIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (cai.Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("api"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("api_config"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}

	return cai.FetchIamPolicy(
		ApiGatewayApiConfigIamUpdaterProducer,
		d,
		config,
		"//apigateway.googleapis.com/projects/{{project}}/locations/global/apis/{{api}}/configs/{{api_config}}",
		ApiGatewayApiConfigIAMAssetType,
	)
}
