// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/iamworkforcepool/WorkforcePoolProviderKey.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package iamworkforcepool

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const workforcePoolProviderKeyIdRegexp = `^[a-z0-9-]{4,32}$`

func ValidateWorkforcePoolProviderKeyId(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)

	if strings.HasPrefix(value, "gcp-") {
		errors = append(errors, fmt.Errorf(
			"%q (%q) can not start with \"gcp-\". "+
				"The prefix `gcp-` is reserved for use by Google, and may not be specified.", k, value))
	}

	if !regexp.MustCompile(workforcePoolProviderKeyIdRegexp).MatchString(value) {
		errors = append(errors, fmt.Errorf(
			"%q (%q) must be 4-32 characters, and may contain the characters [a-z0-9-].", k, value))
	}

	return
}

const IAMWorkforcePoolWorkforcePoolProviderKeyAssetType string = "iam.googleapis.com/WorkforcePoolProviderKey"

func ResourceConverterIAMWorkforcePoolWorkforcePoolProviderKey() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: IAMWorkforcePoolWorkforcePoolProviderKeyAssetType,
		Convert:   GetIAMWorkforcePoolWorkforcePoolProviderKeyCaiObject,
	}
}

func GetIAMWorkforcePoolWorkforcePoolProviderKeyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//iam.googleapis.com/locations/{{location}}/workforcePools/{{workforce_pool_id}}/providers/{{provider_id}}/keys/{{key_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetIAMWorkforcePoolWorkforcePoolProviderKeyApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: IAMWorkforcePoolWorkforcePoolProviderKeyAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/iam/v1/rest",
				DiscoveryName:        "WorkforcePoolProviderKey",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetIAMWorkforcePoolWorkforcePoolProviderKeyApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	keyDataProp, err := expandIAMWorkforcePoolWorkforcePoolProviderKeyKeyData(d.Get("key_data"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("key_data"); !tpgresource.IsEmptyValue(reflect.ValueOf(keyDataProp)) && (ok || !reflect.DeepEqual(v, keyDataProp)) {
		obj["keyData"] = keyDataProp
	}
	useProp, err := expandIAMWorkforcePoolWorkforcePoolProviderKeyUse(d.Get("use"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("use"); !tpgresource.IsEmptyValue(reflect.ValueOf(useProp)) && (ok || !reflect.DeepEqual(v, useProp)) {
		obj["use"] = useProp
	}

	return obj, nil
}

func expandIAMWorkforcePoolWorkforcePoolProviderKeyKeyData(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedFormat, err := expandIAMWorkforcePoolWorkforcePoolProviderKeyKeyDataFormat(original["format"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFormat); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["format"] = transformedFormat
	}

	transformedNotBeforeTime, err := expandIAMWorkforcePoolWorkforcePoolProviderKeyKeyDataNotBeforeTime(original["not_before_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNotBeforeTime); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["notBeforeTime"] = transformedNotBeforeTime
	}

	transformedNotAfterTime, err := expandIAMWorkforcePoolWorkforcePoolProviderKeyKeyDataNotAfterTime(original["not_after_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNotAfterTime); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["notAfterTime"] = transformedNotAfterTime
	}

	transformedKey, err := expandIAMWorkforcePoolWorkforcePoolProviderKeyKeyDataKey(original["key"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKey); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["key"] = transformedKey
	}

	transformedKeySpec, err := expandIAMWorkforcePoolWorkforcePoolProviderKeyKeyDataKeySpec(original["key_spec"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKeySpec); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["keySpec"] = transformedKeySpec
	}

	return transformed, nil
}

func expandIAMWorkforcePoolWorkforcePoolProviderKeyKeyDataFormat(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAMWorkforcePoolWorkforcePoolProviderKeyKeyDataNotBeforeTime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAMWorkforcePoolWorkforcePoolProviderKeyKeyDataNotAfterTime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAMWorkforcePoolWorkforcePoolProviderKeyKeyDataKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAMWorkforcePoolWorkforcePoolProviderKeyKeyDataKeySpec(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAMWorkforcePoolWorkforcePoolProviderKeyUse(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
