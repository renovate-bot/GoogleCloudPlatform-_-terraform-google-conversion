// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/gemini/RepositoryGroup.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package gemini

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const GeminiRepositoryGroupAssetType string = "cloudaicompanion.googleapis.com/RepositoryGroup"

func ResourceConverterGeminiRepositoryGroup() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: GeminiRepositoryGroupAssetType,
		Convert:   GetGeminiRepositoryGroupCaiObject,
	}
}

func GetGeminiRepositoryGroupCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//cloudaicompanion.googleapis.com/projects/{{project}}/locations/{{location}}/codeRepositoryIndexes/{{code_repository_index}}/repositoryGroups/{{repository_group_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetGeminiRepositoryGroupApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: GeminiRepositoryGroupAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/cloudaicompanion/v1/rest",
				DiscoveryName:        "RepositoryGroup",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetGeminiRepositoryGroupApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	repositoriesProp, err := expandGeminiRepositoryGroupRepositories(d.Get("repositories"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("repositories"); !tpgresource.IsEmptyValue(reflect.ValueOf(repositoriesProp)) && (ok || !reflect.DeepEqual(v, repositoriesProp)) {
		obj["repositories"] = repositoriesProp
	}
	labelsProp, err := expandGeminiRepositoryGroupEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandGeminiRepositoryGroupRepositories(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedResource, err := expandGeminiRepositoryGroupRepositoriesResource(original["resource"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedResource); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["resource"] = transformedResource
		}

		transformedBranchPattern, err := expandGeminiRepositoryGroupRepositoriesBranchPattern(original["branch_pattern"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedBranchPattern); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["branchPattern"] = transformedBranchPattern
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandGeminiRepositoryGroupRepositoriesResource(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGeminiRepositoryGroupRepositoriesBranchPattern(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGeminiRepositoryGroupEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
