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

package datacatalog

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgiamresource"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const DataCatalogTagTemplateIAMAssetType string = "datacatalog.googleapis.com/TagTemplate"

func ResourceConverterDataCatalogTagTemplateIamPolicy() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType:         DataCatalogTagTemplateIAMAssetType,
		Convert:           GetDataCatalogTagTemplateIamPolicyCaiObject,
		MergeCreateUpdate: MergeDataCatalogTagTemplateIamPolicy,
	}
}

func ResourceConverterDataCatalogTagTemplateIamBinding() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType:         DataCatalogTagTemplateIAMAssetType,
		Convert:           GetDataCatalogTagTemplateIamBindingCaiObject,
		FetchFullResource: FetchDataCatalogTagTemplateIamPolicy,
		MergeCreateUpdate: MergeDataCatalogTagTemplateIamBinding,
		MergeDelete:       MergeDataCatalogTagTemplateIamBindingDelete,
	}
}

func ResourceConverterDataCatalogTagTemplateIamMember() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType:         DataCatalogTagTemplateIAMAssetType,
		Convert:           GetDataCatalogTagTemplateIamMemberCaiObject,
		FetchFullResource: FetchDataCatalogTagTemplateIamPolicy,
		MergeCreateUpdate: MergeDataCatalogTagTemplateIamMember,
		MergeDelete:       MergeDataCatalogTagTemplateIamMemberDelete,
	}
}

func GetDataCatalogTagTemplateIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	return newDataCatalogTagTemplateIamAsset(d, config, tpgiamresource.ExpandIamPolicyBindings)
}

func GetDataCatalogTagTemplateIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	return newDataCatalogTagTemplateIamAsset(d, config, tpgiamresource.ExpandIamRoleBindings)
}

func GetDataCatalogTagTemplateIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	return newDataCatalogTagTemplateIamAsset(d, config, tpgiamresource.ExpandIamMemberBindings)
}

func MergeDataCatalogTagTemplateIamPolicy(existing, incoming tpgresource.Asset) tpgresource.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeDataCatalogTagTemplateIamBinding(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeIamAssets(existing, incoming, tpgiamresource.MergeAuthoritativeBindings)
}

func MergeDataCatalogTagTemplateIamBindingDelete(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeDeleteIamAssets(existing, incoming, tpgiamresource.MergeDeleteAuthoritativeBindings)
}

func MergeDataCatalogTagTemplateIamMember(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeIamAssets(existing, incoming, tpgiamresource.MergeAdditiveBindings)
}

func MergeDataCatalogTagTemplateIamMemberDelete(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeDeleteIamAssets(existing, incoming, tpgiamresource.MergeDeleteAdditiveBindings)
}

func newDataCatalogTagTemplateIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]tpgresource.IAMBinding, error),
) ([]tpgresource.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []tpgresource.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := tpgresource.AssetName(d, config, "//datacatalog.googleapis.com/projects/{{project}}/locations/{{region}}/tagTemplates/{{tag_template}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}

	return []tpgresource.Asset{{
		Name: name,
		Type: DataCatalogTagTemplateIAMAssetType,
		IAMPolicy: &tpgresource.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchDataCatalogTagTemplateIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (tpgresource.Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("region"); !ok {
		return tpgresource.Asset{}, tpgresource.ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("tag_template"); !ok {
		return tpgresource.Asset{}, tpgresource.ErrEmptyIdentityField
	}

	return tpgiamresource.FetchIamPolicy(
		DataCatalogTagTemplateIamUpdaterProducer,
		d,
		config,
		"//datacatalog.googleapis.com/projects/{{project}}/locations/{{region}}/tagTemplates/{{tag_template}}",
		DataCatalogTagTemplateIAMAssetType,
	)
}