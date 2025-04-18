// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/securityposture/PostureDeployment.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package securityposture

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const SecurityposturePostureDeploymentAssetType string = "securityposture.googleapis.com/PostureDeployment"

func ResourceConverterSecurityposturePostureDeployment() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: SecurityposturePostureDeploymentAssetType,
		Convert:   GetSecurityposturePostureDeploymentCaiObject,
	}
}

func GetSecurityposturePostureDeploymentCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//securityposture.googleapis.com/{{parent}}/locations/{{location}}/postureDeployments/{{posture_deployment_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetSecurityposturePostureDeploymentApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: SecurityposturePostureDeploymentAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/securityposture/v1/rest",
				DiscoveryName:        "PostureDeployment",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetSecurityposturePostureDeploymentApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	targetResourceProp, err := expandSecurityposturePostureDeploymentTargetResource(d.Get("target_resource"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("target_resource"); !tpgresource.IsEmptyValue(reflect.ValueOf(targetResourceProp)) && (ok || !reflect.DeepEqual(v, targetResourceProp)) {
		obj["targetResource"] = targetResourceProp
	}
	postureIdProp, err := expandSecurityposturePostureDeploymentPostureId(d.Get("posture_id"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("posture_id"); !tpgresource.IsEmptyValue(reflect.ValueOf(postureIdProp)) && (ok || !reflect.DeepEqual(v, postureIdProp)) {
		obj["postureId"] = postureIdProp
	}
	postureRevisionIdProp, err := expandSecurityposturePostureDeploymentPostureRevisionId(d.Get("posture_revision_id"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("posture_revision_id"); !tpgresource.IsEmptyValue(reflect.ValueOf(postureRevisionIdProp)) && (ok || !reflect.DeepEqual(v, postureRevisionIdProp)) {
		obj["postureRevisionId"] = postureRevisionIdProp
	}
	descriptionProp, err := expandSecurityposturePostureDeploymentDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}

	return obj, nil
}

func expandSecurityposturePostureDeploymentTargetResource(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityposturePostureDeploymentPostureId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityposturePostureDeploymentPostureRevisionId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityposturePostureDeploymentDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
