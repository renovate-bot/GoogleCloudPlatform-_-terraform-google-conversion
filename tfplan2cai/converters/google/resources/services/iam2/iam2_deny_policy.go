// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package iam2

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

const IAM2DenyPolicyAssetType string = "iam.googleapis.com/DenyPolicy"

func ResourceConverterIAM2DenyPolicy() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType: IAM2DenyPolicyAssetType,
		Convert:   GetIAM2DenyPolicyCaiObject,
	}
}

func GetIAM2DenyPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	name, err := tpgresource.AssetName(d, config, "//iam.googleapis.com/policies/{{parent}}/denypolicies/{{name}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}
	if obj, err := GetIAM2DenyPolicyApiObject(d, config); err == nil {
		return []tpgresource.Asset{{
			Name: name,
			Type: IAM2DenyPolicyAssetType,
			Resource: &tpgresource.AssetResource{
				Version:              "v2beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/iam/v2beta/rest",
				DiscoveryName:        "DenyPolicy",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []tpgresource.Asset{}, err
	}
}

func GetIAM2DenyPolicyApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandIAM2DenyPolicyDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	etagProp, err := expandIAM2DenyPolicyEtag(d.Get("etag"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("etag"); !tpgresource.IsEmptyValue(reflect.ValueOf(etagProp)) && (ok || !reflect.DeepEqual(v, etagProp)) {
		obj["etag"] = etagProp
	}
	rulesProp, err := expandIAM2DenyPolicyRules(d.Get("rules"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("rules"); !tpgresource.IsEmptyValue(reflect.ValueOf(rulesProp)) && (ok || !reflect.DeepEqual(v, rulesProp)) {
		obj["rules"] = rulesProp
	}

	return obj, nil
}

func expandIAM2DenyPolicyDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAM2DenyPolicyEtag(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAM2DenyPolicyRules(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedDescription, err := expandIAM2DenyPolicyRulesDescription(original["description"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["description"] = transformedDescription
		}

		transformedDenyRule, err := expandIAM2DenyPolicyRulesDenyRule(original["deny_rule"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDenyRule); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["denyRule"] = transformedDenyRule
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandIAM2DenyPolicyRulesDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAM2DenyPolicyRulesDenyRule(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDeniedPrincipals, err := expandIAM2DenyPolicyRulesDenyRuleDeniedPrincipals(original["denied_principals"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDeniedPrincipals); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["deniedPrincipals"] = transformedDeniedPrincipals
	}

	transformedExceptionPrincipals, err := expandIAM2DenyPolicyRulesDenyRuleExceptionPrincipals(original["exception_principals"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExceptionPrincipals); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["exceptionPrincipals"] = transformedExceptionPrincipals
	}

	transformedDeniedPermissions, err := expandIAM2DenyPolicyRulesDenyRuleDeniedPermissions(original["denied_permissions"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDeniedPermissions); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["deniedPermissions"] = transformedDeniedPermissions
	}

	transformedExceptionPermissions, err := expandIAM2DenyPolicyRulesDenyRuleExceptionPermissions(original["exception_permissions"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExceptionPermissions); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["exceptionPermissions"] = transformedExceptionPermissions
	}

	transformedDenialCondition, err := expandIAM2DenyPolicyRulesDenyRuleDenialCondition(original["denial_condition"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDenialCondition); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["denialCondition"] = transformedDenialCondition
	}

	return transformed, nil
}

func expandIAM2DenyPolicyRulesDenyRuleDeniedPrincipals(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAM2DenyPolicyRulesDenyRuleExceptionPrincipals(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAM2DenyPolicyRulesDenyRuleDeniedPermissions(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAM2DenyPolicyRulesDenyRuleExceptionPermissions(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAM2DenyPolicyRulesDenyRuleDenialCondition(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedExpression, err := expandIAM2DenyPolicyRulesDenyRuleDenialConditionExpression(original["expression"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExpression); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["expression"] = transformedExpression
	}

	transformedTitle, err := expandIAM2DenyPolicyRulesDenyRuleDenialConditionTitle(original["title"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTitle); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["title"] = transformedTitle
	}

	transformedDescription, err := expandIAM2DenyPolicyRulesDenyRuleDenialConditionDescription(original["description"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["description"] = transformedDescription
	}

	transformedLocation, err := expandIAM2DenyPolicyRulesDenyRuleDenialConditionLocation(original["location"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLocation); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["location"] = transformedLocation
	}

	return transformed, nil
}

func expandIAM2DenyPolicyRulesDenyRuleDenialConditionExpression(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAM2DenyPolicyRulesDenyRuleDenialConditionTitle(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAM2DenyPolicyRulesDenyRuleDenialConditionDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAM2DenyPolicyRulesDenyRuleDenialConditionLocation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}