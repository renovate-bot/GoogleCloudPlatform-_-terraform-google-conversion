// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/discoveryengine/SearchEngine.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package discoveryengine

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const DiscoveryEngineSearchEngineAssetType string = "{{location}}-discoveryengine.googleapis.com/SearchEngine"

func ResourceConverterDiscoveryEngineSearchEngine() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: DiscoveryEngineSearchEngineAssetType,
		Convert:   GetDiscoveryEngineSearchEngineCaiObject,
	}
}

func GetDiscoveryEngineSearchEngineCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//{{location}}-discoveryengine.googleapis.com/projects/{{project}}/locations/{{location}}/collections/{{collection_id}}/engines/{{engine_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetDiscoveryEngineSearchEngineApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: DiscoveryEngineSearchEngineAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/{{location}}-discoveryengine/v1/rest",
				DiscoveryName:        "SearchEngine",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetDiscoveryEngineSearchEngineApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	industryVerticalProp, err := expandDiscoveryEngineSearchEngineIndustryVertical(d.Get("industry_vertical"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("industry_vertical"); !tpgresource.IsEmptyValue(reflect.ValueOf(industryVerticalProp)) && (ok || !reflect.DeepEqual(v, industryVerticalProp)) {
		obj["industryVertical"] = industryVerticalProp
	}
	displayNameProp, err := expandDiscoveryEngineSearchEngineDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	dataStoreIdsProp, err := expandDiscoveryEngineSearchEngineDataStoreIds(d.Get("data_store_ids"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("data_store_ids"); !tpgresource.IsEmptyValue(reflect.ValueOf(dataStoreIdsProp)) && (ok || !reflect.DeepEqual(v, dataStoreIdsProp)) {
		obj["dataStoreIds"] = dataStoreIdsProp
	}
	searchEngineConfigProp, err := expandDiscoveryEngineSearchEngineSearchEngineConfig(d.Get("search_engine_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("search_engine_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(searchEngineConfigProp)) && (ok || !reflect.DeepEqual(v, searchEngineConfigProp)) {
		obj["searchEngineConfig"] = searchEngineConfigProp
	}
	commonConfigProp, err := expandDiscoveryEngineSearchEngineCommonConfig(d.Get("common_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("common_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(commonConfigProp)) && (ok || !reflect.DeepEqual(v, commonConfigProp)) {
		obj["commonConfig"] = commonConfigProp
	}

	return resourceDiscoveryEngineSearchEngineEncoder(d, config, obj)
}

func resourceDiscoveryEngineSearchEngineEncoder(d tpgresource.TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	// hard code solutionType to "SOLUTION_TYPE_SEARCH" for search engine resource
	obj["solutionType"] = "SOLUTION_TYPE_SEARCH"
	return obj, nil
}

func expandDiscoveryEngineSearchEngineIndustryVertical(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDiscoveryEngineSearchEngineDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDiscoveryEngineSearchEngineDataStoreIds(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDiscoveryEngineSearchEngineSearchEngineConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSearchTier, err := expandDiscoveryEngineSearchEngineSearchEngineConfigSearchTier(original["search_tier"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSearchTier); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["searchTier"] = transformedSearchTier
	}

	transformedSearchAddOns, err := expandDiscoveryEngineSearchEngineSearchEngineConfigSearchAddOns(original["search_add_ons"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSearchAddOns); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["searchAddOns"] = transformedSearchAddOns
	}

	return transformed, nil
}

func expandDiscoveryEngineSearchEngineSearchEngineConfigSearchTier(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDiscoveryEngineSearchEngineSearchEngineConfigSearchAddOns(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDiscoveryEngineSearchEngineCommonConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCompanyName, err := expandDiscoveryEngineSearchEngineCommonConfigCompanyName(original["company_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCompanyName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["companyName"] = transformedCompanyName
	}

	return transformed, nil
}

func expandDiscoveryEngineSearchEngineCommonConfigCompanyName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
