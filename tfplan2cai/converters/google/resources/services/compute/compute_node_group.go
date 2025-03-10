// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/compute/NodeGroup.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package compute

import (
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const ComputeNodeGroupAssetType string = "compute.googleapis.com/NodeGroup"

func ResourceConverterComputeNodeGroup() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ComputeNodeGroupAssetType,
		Convert:   GetComputeNodeGroupCaiObject,
	}
}

func GetComputeNodeGroupCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//compute.googleapis.com/projects/{{project}}/zones/{{zone}}/nodeGroups/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetComputeNodeGroupApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ComputeNodeGroupAssetType,
			Resource: &cai.AssetResource{
				Version:              "beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/beta/rest",
				DiscoveryName:        "NodeGroup",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetComputeNodeGroupApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandComputeNodeGroupDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	nameProp, err := expandComputeNodeGroupName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	nodeTemplateProp, err := expandComputeNodeGroupNodeTemplate(d.Get("node_template"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("node_template"); !tpgresource.IsEmptyValue(reflect.ValueOf(nodeTemplateProp)) && (ok || !reflect.DeepEqual(v, nodeTemplateProp)) {
		obj["nodeTemplate"] = nodeTemplateProp
	}
	maintenancePolicyProp, err := expandComputeNodeGroupMaintenancePolicy(d.Get("maintenance_policy"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("maintenance_policy"); !tpgresource.IsEmptyValue(reflect.ValueOf(maintenancePolicyProp)) && (ok || !reflect.DeepEqual(v, maintenancePolicyProp)) {
		obj["maintenancePolicy"] = maintenancePolicyProp
	}
	maintenanceWindowProp, err := expandComputeNodeGroupMaintenanceWindow(d.Get("maintenance_window"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("maintenance_window"); !tpgresource.IsEmptyValue(reflect.ValueOf(maintenanceWindowProp)) && (ok || !reflect.DeepEqual(v, maintenanceWindowProp)) {
		obj["maintenanceWindow"] = maintenanceWindowProp
	}
	autoscalingPolicyProp, err := expandComputeNodeGroupAutoscalingPolicy(d.Get("autoscaling_policy"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("autoscaling_policy"); !tpgresource.IsEmptyValue(reflect.ValueOf(autoscalingPolicyProp)) && (ok || !reflect.DeepEqual(v, autoscalingPolicyProp)) {
		obj["autoscalingPolicy"] = autoscalingPolicyProp
	}
	shareSettingsProp, err := expandComputeNodeGroupShareSettings(d.Get("share_settings"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("share_settings"); !tpgresource.IsEmptyValue(reflect.ValueOf(shareSettingsProp)) && (ok || !reflect.DeepEqual(v, shareSettingsProp)) {
		obj["shareSettings"] = shareSettingsProp
	}
	maintenanceIntervalProp, err := expandComputeNodeGroupMaintenanceInterval(d.Get("maintenance_interval"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("maintenance_interval"); !tpgresource.IsEmptyValue(reflect.ValueOf(maintenanceIntervalProp)) && (ok || !reflect.DeepEqual(v, maintenanceIntervalProp)) {
		obj["maintenanceInterval"] = maintenanceIntervalProp
	}
	zoneProp, err := expandComputeNodeGroupZone(d.Get("zone"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("zone"); !tpgresource.IsEmptyValue(reflect.ValueOf(zoneProp)) && (ok || !reflect.DeepEqual(v, zoneProp)) {
		obj["zone"] = zoneProp
	}

	return obj, nil
}

func expandComputeNodeGroupDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNodeGroupName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNodeGroupNodeTemplate(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseRegionalFieldValue("nodeTemplates", v.(string), "project", "region", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for node_template: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeNodeGroupMaintenancePolicy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNodeGroupMaintenanceWindow(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedStartTime, err := expandComputeNodeGroupMaintenanceWindowStartTime(original["start_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStartTime); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["startTime"] = transformedStartTime
	}

	return transformed, nil
}

func expandComputeNodeGroupMaintenanceWindowStartTime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNodeGroupAutoscalingPolicy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMode, err := expandComputeNodeGroupAutoscalingPolicyMode(original["mode"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMode); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["mode"] = transformedMode
	}

	transformedMinNodes, err := expandComputeNodeGroupAutoscalingPolicyMinNodes(original["min_nodes"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMinNodes); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["minNodes"] = transformedMinNodes
	}

	transformedMaxNodes, err := expandComputeNodeGroupAutoscalingPolicyMaxNodes(original["max_nodes"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxNodes); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxNodes"] = transformedMaxNodes
	}

	return transformed, nil
}

func expandComputeNodeGroupAutoscalingPolicyMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNodeGroupAutoscalingPolicyMinNodes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNodeGroupAutoscalingPolicyMaxNodes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNodeGroupShareSettings(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedShareType, err := expandComputeNodeGroupShareSettingsShareType(original["share_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedShareType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["shareType"] = transformedShareType
	}

	transformedProjectMap, err := expandComputeNodeGroupShareSettingsProjectMap(original["project_map"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProjectMap); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["projectMap"] = transformedProjectMap
	}

	return transformed, nil
}

func expandComputeNodeGroupShareSettingsShareType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNodeGroupShareSettingsProjectMap(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	if v == nil {
		return map[string]interface{}{}, nil
	}
	m := make(map[string]interface{})
	for _, raw := range v.(*schema.Set).List() {
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedProjectId, err := expandComputeNodeGroupShareSettingsProjectMapProjectId(original["project_id"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedProjectId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["projectId"] = transformedProjectId
		}

		transformedId, err := tpgresource.ExpandString(original["id"], d, config)
		if err != nil {
			return nil, err
		}
		m[transformedId] = transformed
	}
	return m, nil
}

func expandComputeNodeGroupShareSettingsProjectMapProjectId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNodeGroupMaintenanceInterval(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNodeGroupZone(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseGlobalFieldValue("zones", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for zone: %s", err)
	}
	return f.RelativeLink(), nil
}
