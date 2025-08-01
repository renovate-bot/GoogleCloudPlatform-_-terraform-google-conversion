// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/bigqueryanalyticshub/DataExchange.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package bigqueryanalyticshub

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const BigqueryAnalyticsHubDataExchangeAssetType string = "analyticshub.googleapis.com/DataExchange"

func ResourceConverterBigqueryAnalyticsHubDataExchange() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: BigqueryAnalyticsHubDataExchangeAssetType,
		Convert:   GetBigqueryAnalyticsHubDataExchangeCaiObject,
	}
}

func GetBigqueryAnalyticsHubDataExchangeCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//analyticshub.googleapis.com/projects/{{project}}/locations/{{location}}/dataExchanges/{{data_exchange_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetBigqueryAnalyticsHubDataExchangeApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: BigqueryAnalyticsHubDataExchangeAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/analyticshub/v1/rest",
				DiscoveryName:        "DataExchange",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetBigqueryAnalyticsHubDataExchangeApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandBigqueryAnalyticsHubDataExchangeDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandBigqueryAnalyticsHubDataExchangeDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	primaryContactProp, err := expandBigqueryAnalyticsHubDataExchangePrimaryContact(d.Get("primary_contact"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("primary_contact"); !tpgresource.IsEmptyValue(reflect.ValueOf(primaryContactProp)) && (ok || !reflect.DeepEqual(v, primaryContactProp)) {
		obj["primaryContact"] = primaryContactProp
	}
	documentationProp, err := expandBigqueryAnalyticsHubDataExchangeDocumentation(d.Get("documentation"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("documentation"); !tpgresource.IsEmptyValue(reflect.ValueOf(documentationProp)) && (ok || !reflect.DeepEqual(v, documentationProp)) {
		obj["documentation"] = documentationProp
	}
	iconProp, err := expandBigqueryAnalyticsHubDataExchangeIcon(d.Get("icon"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("icon"); !tpgresource.IsEmptyValue(reflect.ValueOf(iconProp)) && (ok || !reflect.DeepEqual(v, iconProp)) {
		obj["icon"] = iconProp
	}
	sharingEnvironmentConfigProp, err := expandBigqueryAnalyticsHubDataExchangeSharingEnvironmentConfig(d.Get("sharing_environment_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("sharing_environment_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(sharingEnvironmentConfigProp)) && (ok || !reflect.DeepEqual(v, sharingEnvironmentConfigProp)) {
		obj["sharingEnvironmentConfig"] = sharingEnvironmentConfigProp
	}
	discoveryTypeProp, err := expandBigqueryAnalyticsHubDataExchangeDiscoveryType(d.Get("discovery_type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("discovery_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(discoveryTypeProp)) && (ok || !reflect.DeepEqual(v, discoveryTypeProp)) {
		obj["discoveryType"] = discoveryTypeProp
	}
	logLinkedDatasetQueryUserEmailProp, err := expandBigqueryAnalyticsHubDataExchangeLogLinkedDatasetQueryUserEmail(d.Get("log_linked_dataset_query_user_email"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("log_linked_dataset_query_user_email"); !tpgresource.IsEmptyValue(reflect.ValueOf(logLinkedDatasetQueryUserEmailProp)) && (ok || !reflect.DeepEqual(v, logLinkedDatasetQueryUserEmailProp)) {
		obj["logLinkedDatasetQueryUserEmail"] = logLinkedDatasetQueryUserEmailProp
	}

	return obj, nil
}

func expandBigqueryAnalyticsHubDataExchangeDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryAnalyticsHubDataExchangeDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryAnalyticsHubDataExchangePrimaryContact(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryAnalyticsHubDataExchangeDocumentation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryAnalyticsHubDataExchangeIcon(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryAnalyticsHubDataExchangeSharingEnvironmentConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDefaultExchangeConfig, err := expandBigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigDefaultExchangeConfig(original["default_exchange_config"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["defaultExchangeConfig"] = transformedDefaultExchangeConfig
	}

	transformedDcrExchangeConfig, err := expandBigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigDcrExchangeConfig(original["dcr_exchange_config"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["dcrExchangeConfig"] = transformedDcrExchangeConfig
	}

	return transformed, nil
}

func expandBigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigDefaultExchangeConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 {
		return nil, nil
	}

	if l[0] == nil {
		transformed := make(map[string]interface{})
		return transformed, nil
	}
	transformed := make(map[string]interface{})

	return transformed, nil
}

func expandBigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigDcrExchangeConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 {
		return nil, nil
	}

	if l[0] == nil {
		transformed := make(map[string]interface{})
		return transformed, nil
	}
	transformed := make(map[string]interface{})

	return transformed, nil
}

func expandBigqueryAnalyticsHubDataExchangeDiscoveryType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryAnalyticsHubDataExchangeLogLinkedDatasetQueryUserEmail(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
