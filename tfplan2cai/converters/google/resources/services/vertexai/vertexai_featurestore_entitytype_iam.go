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

package vertexai

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const VertexAIFeaturestoreEntitytypeIAMAssetType string = "aiplatform.googleapis.com/FeaturestoreEntitytype"

func ResourceConverterVertexAIFeaturestoreEntitytypeIamPolicy() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         VertexAIFeaturestoreEntitytypeIAMAssetType,
		Convert:           GetVertexAIFeaturestoreEntitytypeIamPolicyCaiObject,
		MergeCreateUpdate: MergeVertexAIFeaturestoreEntitytypeIamPolicy,
	}
}

func ResourceConverterVertexAIFeaturestoreEntitytypeIamBinding() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         VertexAIFeaturestoreEntitytypeIAMAssetType,
		Convert:           GetVertexAIFeaturestoreEntitytypeIamBindingCaiObject,
		FetchFullResource: FetchVertexAIFeaturestoreEntitytypeIamPolicy,
		MergeCreateUpdate: MergeVertexAIFeaturestoreEntitytypeIamBinding,
		MergeDelete:       MergeVertexAIFeaturestoreEntitytypeIamBindingDelete,
	}
}

func ResourceConverterVertexAIFeaturestoreEntitytypeIamMember() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         VertexAIFeaturestoreEntitytypeIAMAssetType,
		Convert:           GetVertexAIFeaturestoreEntitytypeIamMemberCaiObject,
		FetchFullResource: FetchVertexAIFeaturestoreEntitytypeIamPolicy,
		MergeCreateUpdate: MergeVertexAIFeaturestoreEntitytypeIamMember,
		MergeDelete:       MergeVertexAIFeaturestoreEntitytypeIamMemberDelete,
	}
}

func GetVertexAIFeaturestoreEntitytypeIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newVertexAIFeaturestoreEntitytypeIamAsset(d, config, cai.ExpandIamPolicyBindings)
}

func GetVertexAIFeaturestoreEntitytypeIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newVertexAIFeaturestoreEntitytypeIamAsset(d, config, cai.ExpandIamRoleBindings)
}

func GetVertexAIFeaturestoreEntitytypeIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newVertexAIFeaturestoreEntitytypeIamAsset(d, config, cai.ExpandIamMemberBindings)
}

func MergeVertexAIFeaturestoreEntitytypeIamPolicy(existing, incoming cai.Asset) cai.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeVertexAIFeaturestoreEntitytypeIamBinding(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAuthoritativeBindings)
}

func MergeVertexAIFeaturestoreEntitytypeIamBindingDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAuthoritativeBindings)
}

func MergeVertexAIFeaturestoreEntitytypeIamMember(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAdditiveBindings)
}

func MergeVertexAIFeaturestoreEntitytypeIamMemberDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAdditiveBindings)
}

func newVertexAIFeaturestoreEntitytypeIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]cai.IAMBinding, error),
) ([]cai.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []cai.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := cai.AssetName(d, config, "//aiplatform.googleapis.com/{{featurestore}}/entityTypes/{{entitytype}}")
	if err != nil {
		return []cai.Asset{}, err
	}

	return []cai.Asset{{
		Name: name,
		Type: VertexAIFeaturestoreEntitytypeIAMAssetType,
		IAMPolicy: &cai.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchVertexAIFeaturestoreEntitytypeIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (cai.Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("featurestore"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("entitytype"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}

	return cai.FetchIamPolicy(
		VertexAIFeaturestoreEntitytypeIamUpdaterProducer,
		d,
		config,
		"//aiplatform.googleapis.com/{{featurestore}}/entityTypes/{{entitytype}}",
		VertexAIFeaturestoreEntitytypeIAMAssetType,
	)
}
