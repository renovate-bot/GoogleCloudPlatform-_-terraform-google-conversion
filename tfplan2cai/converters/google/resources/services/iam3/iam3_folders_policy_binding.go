// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/iam3/FoldersPolicyBinding.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package iam3

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const IAM3FoldersPolicyBindingAssetType string = "iam.googleapis.com/FoldersPolicyBinding"

func ResourceConverterIAM3FoldersPolicyBinding() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: IAM3FoldersPolicyBindingAssetType,
		Convert:   GetIAM3FoldersPolicyBindingCaiObject,
	}
}

func GetIAM3FoldersPolicyBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//iam.googleapis.com/folders/{{folder}}/locations/{{location}}/policyBindings/{{policy_binding_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetIAM3FoldersPolicyBindingApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: IAM3FoldersPolicyBindingAssetType,
			Resource: &cai.AssetResource{
				Version:              "v3beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/iam/v3beta/rest",
				DiscoveryName:        "FoldersPolicyBinding",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetIAM3FoldersPolicyBindingApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandIAM3FoldersPolicyBindingDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	targetProp, err := expandIAM3FoldersPolicyBindingTarget(d.Get("target"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("target"); !tpgresource.IsEmptyValue(reflect.ValueOf(targetProp)) && (ok || !reflect.DeepEqual(v, targetProp)) {
		obj["target"] = targetProp
	}
	policyKindProp, err := expandIAM3FoldersPolicyBindingPolicyKind(d.Get("policy_kind"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("policy_kind"); !tpgresource.IsEmptyValue(reflect.ValueOf(policyKindProp)) && (ok || !reflect.DeepEqual(v, policyKindProp)) {
		obj["policyKind"] = policyKindProp
	}
	policyProp, err := expandIAM3FoldersPolicyBindingPolicy(d.Get("policy"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("policy"); !tpgresource.IsEmptyValue(reflect.ValueOf(policyProp)) && (ok || !reflect.DeepEqual(v, policyProp)) {
		obj["policy"] = policyProp
	}
	conditionProp, err := expandIAM3FoldersPolicyBindingCondition(d.Get("condition"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("condition"); !tpgresource.IsEmptyValue(reflect.ValueOf(conditionProp)) && (ok || !reflect.DeepEqual(v, conditionProp)) {
		obj["condition"] = conditionProp
	}
	annotationsProp, err := expandIAM3FoldersPolicyBindingEffectiveAnnotations(d.Get("effective_annotations"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_annotations"); !tpgresource.IsEmptyValue(reflect.ValueOf(annotationsProp)) && (ok || !reflect.DeepEqual(v, annotationsProp)) {
		obj["annotations"] = annotationsProp
	}

	return obj, nil
}

func expandIAM3FoldersPolicyBindingDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAM3FoldersPolicyBindingTarget(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPrincipalSet, err := expandIAM3FoldersPolicyBindingTargetPrincipalSet(original["principal_set"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPrincipalSet); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["principalSet"] = transformedPrincipalSet
	}

	return transformed, nil
}

func expandIAM3FoldersPolicyBindingTargetPrincipalSet(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAM3FoldersPolicyBindingPolicyKind(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAM3FoldersPolicyBindingPolicy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAM3FoldersPolicyBindingCondition(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedExpression, err := expandIAM3FoldersPolicyBindingConditionExpression(original["expression"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExpression); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["expression"] = transformedExpression
	}

	transformedTitle, err := expandIAM3FoldersPolicyBindingConditionTitle(original["title"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTitle); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["title"] = transformedTitle
	}

	transformedDescription, err := expandIAM3FoldersPolicyBindingConditionDescription(original["description"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["description"] = transformedDescription
	}

	transformedLocation, err := expandIAM3FoldersPolicyBindingConditionLocation(original["location"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLocation); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["location"] = transformedLocation
	}

	return transformed, nil
}

func expandIAM3FoldersPolicyBindingConditionExpression(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAM3FoldersPolicyBindingConditionTitle(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAM3FoldersPolicyBindingConditionDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAM3FoldersPolicyBindingConditionLocation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAM3FoldersPolicyBindingEffectiveAnnotations(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
