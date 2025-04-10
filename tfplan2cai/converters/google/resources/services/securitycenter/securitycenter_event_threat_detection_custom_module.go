// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/securitycenter/EventThreatDetectionCustomModule.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package securitycenter

import (
	"encoding/json"
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const SecurityCenterEventThreatDetectionCustomModuleAssetType string = "securitycenter.googleapis.com/EventThreatDetectionCustomModule"

func ResourceConverterSecurityCenterEventThreatDetectionCustomModule() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: SecurityCenterEventThreatDetectionCustomModuleAssetType,
		Convert:   GetSecurityCenterEventThreatDetectionCustomModuleCaiObject,
	}
}

func GetSecurityCenterEventThreatDetectionCustomModuleCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//securitycenter.googleapis.com/organizations/{{organization}}/eventThreatDetectionSettings/customModules/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetSecurityCenterEventThreatDetectionCustomModuleApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: SecurityCenterEventThreatDetectionCustomModuleAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/securitycenter/v1/rest",
				DiscoveryName:        "EventThreatDetectionCustomModule",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetSecurityCenterEventThreatDetectionCustomModuleApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	configProp, err := expandSecurityCenterEventThreatDetectionCustomModuleConfig(d.Get("config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("config"); !tpgresource.IsEmptyValue(reflect.ValueOf(configProp)) && (ok || !reflect.DeepEqual(v, configProp)) {
		obj["config"] = configProp
	}
	enablementStateProp, err := expandSecurityCenterEventThreatDetectionCustomModuleEnablementState(d.Get("enablement_state"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("enablement_state"); !tpgresource.IsEmptyValue(reflect.ValueOf(enablementStateProp)) && (ok || !reflect.DeepEqual(v, enablementStateProp)) {
		obj["enablementState"] = enablementStateProp
	}
	typeProp, err := expandSecurityCenterEventThreatDetectionCustomModuleType(d.Get("type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("type"); !tpgresource.IsEmptyValue(reflect.ValueOf(typeProp)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	displayNameProp, err := expandSecurityCenterEventThreatDetectionCustomModuleDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}

	return obj, nil
}

func expandSecurityCenterEventThreatDetectionCustomModuleConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	b := []byte(v.(string))
	if len(b) == 0 {
		return nil, nil
	}
	m := make(map[string]interface{})
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return m, nil
}

func expandSecurityCenterEventThreatDetectionCustomModuleEnablementState(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterEventThreatDetectionCustomModuleType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterEventThreatDetectionCustomModuleDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
