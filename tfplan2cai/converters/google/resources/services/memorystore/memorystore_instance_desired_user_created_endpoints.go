// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/memorystore/InstanceDesiredUserCreatedEndpoints.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package memorystore

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const MemorystoreInstanceDesiredUserCreatedEndpointsAssetType string = "memorystore.googleapis.com/InstanceDesiredUserCreatedEndpoints"

func ResourceConverterMemorystoreInstanceDesiredUserCreatedEndpoints() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: MemorystoreInstanceDesiredUserCreatedEndpointsAssetType,
		Convert:   GetMemorystoreInstanceDesiredUserCreatedEndpointsCaiObject,
	}
}

func GetMemorystoreInstanceDesiredUserCreatedEndpointsCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//memorystore.googleapis.com/projects/{{project}}/locations/{{region}}/instances/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetMemorystoreInstanceDesiredUserCreatedEndpointsApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: MemorystoreInstanceDesiredUserCreatedEndpointsAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/memorystore/v1beta/rest",
				DiscoveryName:        "InstanceDesiredUserCreatedEndpoints",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetMemorystoreInstanceDesiredUserCreatedEndpointsApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	endpointsProp, err := expandMemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpoints(d.Get("desired_user_created_endpoints"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("desired_user_created_endpoints"); !tpgresource.IsEmptyValue(reflect.ValueOf(endpointsProp)) && (ok || !reflect.DeepEqual(v, endpointsProp)) {
		obj["endpoints"] = endpointsProp
	}

	return obj, nil
}

func expandMemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpoints(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedConnections, err := expandMemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnections(original["connections"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedConnections); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["connections"] = transformedConnections
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandMemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnections(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedPscConnection, err := expandMemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsPscConnection(original["psc_connection"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPscConnection); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["pscConnection"] = transformedPscConnection
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandMemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsPscConnection(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPscConnectionId, err := expandMemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsPscConnectionPscConnectionId(original["psc_connection_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPscConnectionId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["pscConnectionId"] = transformedPscConnectionId
	}

	transformedIpAddress, err := expandMemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsPscConnectionIpAddress(original["ip_address"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIpAddress); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["ipAddress"] = transformedIpAddress
	}

	transformedForwardingRule, err := expandMemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsPscConnectionForwardingRule(original["forwarding_rule"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedForwardingRule); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["forwardingRule"] = transformedForwardingRule
	}

	transformedProjectId, err := expandMemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsPscConnectionProjectId(original["project_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProjectId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["projectId"] = transformedProjectId
	}

	transformedNetwork, err := expandMemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsPscConnectionNetwork(original["network"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNetwork); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["network"] = transformedNetwork
	}

	transformedServiceAttachment, err := expandMemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsPscConnectionServiceAttachment(original["service_attachment"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedServiceAttachment); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["serviceAttachment"] = transformedServiceAttachment
	}

	transformedPscConnectionStatus, err := expandMemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsPscConnectionPscConnectionStatus(original["psc_connection_status"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPscConnectionStatus); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["pscConnectionStatus"] = transformedPscConnectionStatus
	}

	transformedConnectionType, err := expandMemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsPscConnectionConnectionType(original["connection_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedConnectionType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["connectionType"] = transformedConnectionType
	}

	return transformed, nil
}

func expandMemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsPscConnectionPscConnectionId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsPscConnectionIpAddress(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsPscConnectionForwardingRule(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsPscConnectionProjectId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsPscConnectionNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsPscConnectionServiceAttachment(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsPscConnectionPscConnectionStatus(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsPscConnectionConnectionType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
