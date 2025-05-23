// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/kms/KeyHandle.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package kms

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const KMSKeyHandleAssetType string = "cloudkms.googleapis.com/KeyHandle"

func ResourceConverterKMSKeyHandle() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: KMSKeyHandleAssetType,
		Convert:   GetKMSKeyHandleCaiObject,
	}
}

func GetKMSKeyHandleCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//cloudkms.googleapis.com/projects/{{project}}/locations/{{location}}/keyHandles/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetKMSKeyHandleApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: KMSKeyHandleAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/cloudkms/v1/rest",
				DiscoveryName:        "KeyHandle",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetKMSKeyHandleApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandKMSKeyHandleName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	resourceTypeSelectorProp, err := expandKMSKeyHandleResourceTypeSelector(d.Get("resource_type_selector"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("resource_type_selector"); !tpgresource.IsEmptyValue(reflect.ValueOf(resourceTypeSelectorProp)) && (ok || !reflect.DeepEqual(v, resourceTypeSelectorProp)) {
		obj["resourceTypeSelector"] = resourceTypeSelectorProp
	}

	return obj, nil
}

func expandKMSKeyHandleName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandKMSKeyHandleResourceTypeSelector(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
