// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/secretmanagerregional/RegionalSecretVersion.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package secretmanagerregional

import (
	"encoding/base64"
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const SecretManagerRegionalRegionalSecretVersionAssetType string = "secretmanager.googleapis.com/RegionalSecretVersion"

func ResourceConverterSecretManagerRegionalRegionalSecretVersion() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: SecretManagerRegionalRegionalSecretVersionAssetType,
		Convert:   GetSecretManagerRegionalRegionalSecretVersionCaiObject,
	}
}

func GetSecretManagerRegionalRegionalSecretVersionCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//secretmanager.googleapis.com/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetSecretManagerRegionalRegionalSecretVersionApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: SecretManagerRegionalRegionalSecretVersionAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/secretmanager/v1/rest",
				DiscoveryName:        "RegionalSecretVersion",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetSecretManagerRegionalRegionalSecretVersionApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	stateProp, err := expandSecretManagerRegionalRegionalSecretVersionEnabled(d.Get("enabled"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("enabled"); !tpgresource.IsEmptyValue(reflect.ValueOf(stateProp)) && (ok || !reflect.DeepEqual(v, stateProp)) {
		obj["state"] = stateProp
	}
	payloadProp, err := expandSecretManagerRegionalRegionalSecretVersionPayload(nil, d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("payload"); !tpgresource.IsEmptyValue(reflect.ValueOf(payloadProp)) && (ok || !reflect.DeepEqual(v, payloadProp)) {
		obj["payload"] = payloadProp
	}

	return obj, nil
}

func expandSecretManagerRegionalRegionalSecretVersionEnabled(_ interface{}, _ tpgresource.TerraformResourceData, _ *transport_tpg.Config) (interface{}, error) {
	return nil, nil
}

func expandSecretManagerRegionalRegionalSecretVersionPayload(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	transformed := make(map[string]interface{})
	transformedSecretData, err := expandSecretManagerRegionalRegionalSecretVersionPayloadSecretData(d.Get("secret_data"), d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSecretData); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["data"] = transformedSecretData
	}

	return transformed, nil
}

func expandSecretManagerRegionalRegionalSecretVersionPayloadSecretData(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	if v == nil {
		return nil, nil
	}

	if d.Get("is_secret_data_base64").(bool) {
		return v, nil
	}
	return base64.StdEncoding.EncodeToString([]byte(v.(string))), nil
}
