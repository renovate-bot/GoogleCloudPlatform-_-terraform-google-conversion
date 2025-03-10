// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/dns/ResponsePolicy.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package dns

import (
	"reflect"
	"strings"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const DNSResponsePolicyAssetType string = "dns.googleapis.com/ResponsePolicy"

func ResourceConverterDNSResponsePolicy() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: DNSResponsePolicyAssetType,
		Convert:   GetDNSResponsePolicyCaiObject,
	}
}

func GetDNSResponsePolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//dns.googleapis.com/projects/{{project}}/responsePolicies/{{response_policy_name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetDNSResponsePolicyApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: DNSResponsePolicyAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta2",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/dns/v1beta2/rest",
				DiscoveryName:        "ResponsePolicy",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetDNSResponsePolicyApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	responsePolicyNameProp, err := expandDNSResponsePolicyResponsePolicyName(d.Get("response_policy_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("response_policy_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(responsePolicyNameProp)) && (ok || !reflect.DeepEqual(v, responsePolicyNameProp)) {
		obj["responsePolicyName"] = responsePolicyNameProp
	}
	descriptionProp, err := expandDNSResponsePolicyDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	networksProp, err := expandDNSResponsePolicyNetworks(d.Get("networks"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("networks"); ok || !reflect.DeepEqual(v, networksProp) {
		obj["networks"] = networksProp
	}
	gkeClustersProp, err := expandDNSResponsePolicyGkeClusters(d.Get("gke_clusters"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("gke_clusters"); !tpgresource.IsEmptyValue(reflect.ValueOf(gkeClustersProp)) && (ok || !reflect.DeepEqual(v, gkeClustersProp)) {
		obj["gkeClusters"] = gkeClustersProp
	}

	return obj, nil
}

func expandDNSResponsePolicyResponsePolicyName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDNSResponsePolicyDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDNSResponsePolicyNetworks(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedNetworkUrl, err := expandDNSResponsePolicyNetworksNetworkUrl(original["network_url"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedNetworkUrl); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["networkUrl"] = transformedNetworkUrl
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDNSResponsePolicyNetworksNetworkUrl(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	if v == nil || v.(string) == "" {
		return "", nil
	} else if strings.HasPrefix(v.(string), "https://") {
		return v, nil
	}
	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}"+v.(string))
	if err != nil {
		return "", err
	}
	return tpgresource.ConvertSelfLinkToV1(url), nil
}

func expandDNSResponsePolicyGkeClusters(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedGkeClusterName, err := expandDNSResponsePolicyGkeClustersGkeClusterName(original["gke_cluster_name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedGkeClusterName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["gkeClusterName"] = transformedGkeClusterName
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDNSResponsePolicyGkeClustersGkeClusterName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
