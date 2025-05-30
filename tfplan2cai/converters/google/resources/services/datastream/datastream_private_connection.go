// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/datastream/PrivateConnection.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package datastream

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const DatastreamPrivateConnectionAssetType string = "datastream.googleapis.com/PrivateConnection"

func ResourceConverterDatastreamPrivateConnection() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: DatastreamPrivateConnectionAssetType,
		Convert:   GetDatastreamPrivateConnectionCaiObject,
	}
}

func GetDatastreamPrivateConnectionCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//datastream.googleapis.com/projects/{{project}}/locations/{{location}}/privateConnections/{{private_connection_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetDatastreamPrivateConnectionApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: DatastreamPrivateConnectionAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/datastream/v1/rest",
				DiscoveryName:        "PrivateConnection",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetDatastreamPrivateConnectionApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandDatastreamPrivateConnectionDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	vpcPeeringConfigProp, err := expandDatastreamPrivateConnectionVpcPeeringConfig(d.Get("vpc_peering_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("vpc_peering_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(vpcPeeringConfigProp)) && (ok || !reflect.DeepEqual(v, vpcPeeringConfigProp)) {
		obj["vpcPeeringConfig"] = vpcPeeringConfigProp
	}
	pscInterfaceConfigProp, err := expandDatastreamPrivateConnectionPscInterfaceConfig(d.Get("psc_interface_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("psc_interface_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(pscInterfaceConfigProp)) && (ok || !reflect.DeepEqual(v, pscInterfaceConfigProp)) {
		obj["pscInterfaceConfig"] = pscInterfaceConfigProp
	}
	labelsProp, err := expandDatastreamPrivateConnectionEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandDatastreamPrivateConnectionDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDatastreamPrivateConnectionVpcPeeringConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedVpc, err := expandDatastreamPrivateConnectionVpcPeeringConfigVpc(original["vpc"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedVpc); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["vpc"] = transformedVpc
	}

	transformedSubnet, err := expandDatastreamPrivateConnectionVpcPeeringConfigSubnet(original["subnet"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSubnet); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["subnet"] = transformedSubnet
	}

	return transformed, nil
}

func expandDatastreamPrivateConnectionVpcPeeringConfigVpc(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDatastreamPrivateConnectionVpcPeeringConfigSubnet(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDatastreamPrivateConnectionPscInterfaceConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedNetworkAttachment, err := expandDatastreamPrivateConnectionPscInterfaceConfigNetworkAttachment(original["network_attachment"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNetworkAttachment); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["networkAttachment"] = transformedNetworkAttachment
	}

	return transformed, nil
}

func expandDatastreamPrivateConnectionPscInterfaceConfigNetworkAttachment(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDatastreamPrivateConnectionEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
