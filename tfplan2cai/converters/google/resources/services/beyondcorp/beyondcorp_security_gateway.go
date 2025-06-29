// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/beyondcorp/SecurityGateway.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package beyondcorp

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func beyondcorpSecurityGatewayHubsHash(v interface{}) int {
	if v == nil {
		return 0
	}

	var buf bytes.Buffer
	m := v.(map[string]interface{})

	buf.WriteString(fmt.Sprintf("%s-", m["region"].(string)))

	return tpgresource.Hashcode(buf.String())
}

const BeyondcorpSecurityGatewayAssetType string = "beyondcorp.googleapis.com/SecurityGateway"

func ResourceConverterBeyondcorpSecurityGateway() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: BeyondcorpSecurityGatewayAssetType,
		Convert:   GetBeyondcorpSecurityGatewayCaiObject,
	}
}

func GetBeyondcorpSecurityGatewayCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//beyondcorp.googleapis.com/projects/{{project}}/locations/{{location}}/securityGateways/{{security_gateway_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetBeyondcorpSecurityGatewayApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: BeyondcorpSecurityGatewayAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/beyondcorp/v1/rest",
				DiscoveryName:        "SecurityGateway",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetBeyondcorpSecurityGatewayApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	hubsProp, err := expandBeyondcorpSecurityGatewayHubs(d.Get("hubs"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("hubs"); !tpgresource.IsEmptyValue(reflect.ValueOf(hubsProp)) && (ok || !reflect.DeepEqual(v, hubsProp)) {
		obj["hubs"] = hubsProp
	}
	displayNameProp, err := expandBeyondcorpSecurityGatewayDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}

	return obj, nil
}

func expandBeyondcorpSecurityGatewayHubs(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	if v == nil {
		return map[string]interface{}{}, nil
	}
	m := make(map[string]interface{})
	for _, raw := range v.(*schema.Set).List() {
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedInternetGateway, err := expandBeyondcorpSecurityGatewayHubsInternetGateway(original["internet_gateway"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedInternetGateway); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["internetGateway"] = transformedInternetGateway
		}

		transformedRegion, err := tpgresource.ExpandString(original["region"], d, config)
		if err != nil {
			return nil, err
		}
		m[transformedRegion] = transformed
	}
	return m, nil
}

func expandBeyondcorpSecurityGatewayHubsInternetGateway(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAssignedIps, err := expandBeyondcorpSecurityGatewayHubsInternetGatewayAssignedIps(original["assigned_ips"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAssignedIps); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["assignedIps"] = transformedAssignedIps
	}

	return transformed, nil
}

func expandBeyondcorpSecurityGatewayHubsInternetGatewayAssignedIps(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBeyondcorpSecurityGatewayDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
