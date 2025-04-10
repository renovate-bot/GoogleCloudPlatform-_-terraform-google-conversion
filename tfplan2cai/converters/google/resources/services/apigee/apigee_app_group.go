// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/apigee/AppGroup.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package apigee

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const ApigeeAppGroupAssetType string = "apigee.googleapis.com/AppGroup"

func ResourceConverterApigeeAppGroup() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ApigeeAppGroupAssetType,
		Convert:   GetApigeeAppGroupCaiObject,
	}
}

func GetApigeeAppGroupCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//apigee.googleapis.com/{{org_id}}/appgroups/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetApigeeAppGroupApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ApigeeAppGroupAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/apigee/v1/rest",
				DiscoveryName:        "AppGroup",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetApigeeAppGroupApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandApigeeAppGroupName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	channelUriProp, err := expandApigeeAppGroupChannelUri(d.Get("channel_uri"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("channel_uri"); !tpgresource.IsEmptyValue(reflect.ValueOf(channelUriProp)) && (ok || !reflect.DeepEqual(v, channelUriProp)) {
		obj["channelUri"] = channelUriProp
	}
	channelIdProp, err := expandApigeeAppGroupChannelId(d.Get("channel_id"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("channel_id"); !tpgresource.IsEmptyValue(reflect.ValueOf(channelIdProp)) && (ok || !reflect.DeepEqual(v, channelIdProp)) {
		obj["channelId"] = channelIdProp
	}
	displayNameProp, err := expandApigeeAppGroupDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	statusProp, err := expandApigeeAppGroupStatus(d.Get("status"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("status"); !tpgresource.IsEmptyValue(reflect.ValueOf(statusProp)) && (ok || !reflect.DeepEqual(v, statusProp)) {
		obj["status"] = statusProp
	}
	attributesProp, err := expandApigeeAppGroupAttributes(d.Get("attributes"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("attributes"); !tpgresource.IsEmptyValue(reflect.ValueOf(attributesProp)) && (ok || !reflect.DeepEqual(v, attributesProp)) {
		obj["attributes"] = attributesProp
	}

	return obj, nil
}

func expandApigeeAppGroupName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeAppGroupChannelUri(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeAppGroupChannelId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeAppGroupDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeAppGroupStatus(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeAppGroupAttributes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandApigeeAppGroupAttributesName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["name"] = transformedName
		}

		transformedValue, err := expandApigeeAppGroupAttributesValue(original["value"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedValue); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["value"] = transformedValue
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandApigeeAppGroupAttributesName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeAppGroupAttributesValue(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
