// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/compute/NetworkPeeringRoutesConfig.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package compute

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const ComputeNetworkPeeringRoutesConfigAssetType string = "compute.googleapis.com/NetworkPeeringRoutesConfig"

func ResourceConverterComputeNetworkPeeringRoutesConfig() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ComputeNetworkPeeringRoutesConfigAssetType,
		Convert:   GetComputeNetworkPeeringRoutesConfigCaiObject,
	}
}

func GetComputeNetworkPeeringRoutesConfigCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//compute.googleapis.com/projects/{{project}}/global/networks/{{network}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetComputeNetworkPeeringRoutesConfigApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ComputeNetworkPeeringRoutesConfigAssetType,
			Resource: &cai.AssetResource{
				Version:              "beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/beta/rest",
				DiscoveryName:        "NetworkPeeringRoutesConfig",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetComputeNetworkPeeringRoutesConfigApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandComputeNetworkPeeringRoutesConfigPeering(d.Get("peering"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("peering"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	exportCustomRoutesProp, err := expandComputeNetworkPeeringRoutesConfigExportCustomRoutes(d.Get("export_custom_routes"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("export_custom_routes"); ok || !reflect.DeepEqual(v, exportCustomRoutesProp) {
		obj["exportCustomRoutes"] = exportCustomRoutesProp
	}
	importCustomRoutesProp, err := expandComputeNetworkPeeringRoutesConfigImportCustomRoutes(d.Get("import_custom_routes"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("import_custom_routes"); ok || !reflect.DeepEqual(v, importCustomRoutesProp) {
		obj["importCustomRoutes"] = importCustomRoutesProp
	}
	exportSubnetRoutesWithPublicIpProp, err := expandComputeNetworkPeeringRoutesConfigExportSubnetRoutesWithPublicIp(d.Get("export_subnet_routes_with_public_ip"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("export_subnet_routes_with_public_ip"); ok || !reflect.DeepEqual(v, exportSubnetRoutesWithPublicIpProp) {
		obj["exportSubnetRoutesWithPublicIp"] = exportSubnetRoutesWithPublicIpProp
	}
	importSubnetRoutesWithPublicIpProp, err := expandComputeNetworkPeeringRoutesConfigImportSubnetRoutesWithPublicIp(d.Get("import_subnet_routes_with_public_ip"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("import_subnet_routes_with_public_ip"); ok || !reflect.DeepEqual(v, importSubnetRoutesWithPublicIpProp) {
		obj["importSubnetRoutesWithPublicIp"] = importSubnetRoutesWithPublicIpProp
	}

	return resourceComputeNetworkPeeringRoutesConfigEncoder(d, config, obj)
}

func resourceComputeNetworkPeeringRoutesConfigEncoder(d tpgresource.TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	// Stick request in a networkPeering block as in
	// https://cloud.google.com/compute/docs/reference/rest/v1/networks/updatePeering
	newObj := make(map[string]interface{})
	newObj["networkPeering"] = obj
	return newObj, nil
}

func expandComputeNetworkPeeringRoutesConfigPeering(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkPeeringRoutesConfigExportCustomRoutes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkPeeringRoutesConfigImportCustomRoutes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkPeeringRoutesConfigExportSubnetRoutesWithPublicIp(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkPeeringRoutesConfigImportSubnetRoutesWithPublicIp(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
