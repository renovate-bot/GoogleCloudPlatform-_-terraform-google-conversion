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

package dataform

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const DataformRepositoryWorkflowConfigAssetType string = "dataform.googleapis.com/RepositoryWorkflowConfig"

func ResourceConverterDataformRepositoryWorkflowConfig() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: DataformRepositoryWorkflowConfigAssetType,
		Convert:   GetDataformRepositoryWorkflowConfigCaiObject,
	}
}

func GetDataformRepositoryWorkflowConfigCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//dataform.googleapis.com/projects/{{project}}/locations/{{region}}/repositories/{{repository}}/workflowConfigs/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetDataformRepositoryWorkflowConfigApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: DataformRepositoryWorkflowConfigAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/dataform/v1beta1/rest",
				DiscoveryName:        "RepositoryWorkflowConfig",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetDataformRepositoryWorkflowConfigApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandDataformRepositoryWorkflowConfigName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	releaseConfigProp, err := expandDataformRepositoryWorkflowConfigReleaseConfig(d.Get("release_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("release_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(releaseConfigProp)) && (ok || !reflect.DeepEqual(v, releaseConfigProp)) {
		obj["releaseConfig"] = releaseConfigProp
	}
	invocationConfigProp, err := expandDataformRepositoryWorkflowConfigInvocationConfig(d.Get("invocation_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("invocation_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(invocationConfigProp)) && (ok || !reflect.DeepEqual(v, invocationConfigProp)) {
		obj["invocationConfig"] = invocationConfigProp
	}
	cronScheduleProp, err := expandDataformRepositoryWorkflowConfigCronSchedule(d.Get("cron_schedule"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("cron_schedule"); !tpgresource.IsEmptyValue(reflect.ValueOf(cronScheduleProp)) && (ok || !reflect.DeepEqual(v, cronScheduleProp)) {
		obj["cronSchedule"] = cronScheduleProp
	}
	timeZoneProp, err := expandDataformRepositoryWorkflowConfigTimeZone(d.Get("time_zone"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("time_zone"); !tpgresource.IsEmptyValue(reflect.ValueOf(timeZoneProp)) && (ok || !reflect.DeepEqual(v, timeZoneProp)) {
		obj["timeZone"] = timeZoneProp
	}

	return obj, nil
}

func expandDataformRepositoryWorkflowConfigName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataformRepositoryWorkflowConfigReleaseConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataformRepositoryWorkflowConfigInvocationConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedIncludedTargets, err := expandDataformRepositoryWorkflowConfigInvocationConfigIncludedTargets(original["included_targets"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIncludedTargets); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["includedTargets"] = transformedIncludedTargets
	}

	transformedIncludedTags, err := expandDataformRepositoryWorkflowConfigInvocationConfigIncludedTags(original["included_tags"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIncludedTags); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["includedTags"] = transformedIncludedTags
	}

	transformedTransitiveDependenciesIncluded, err := expandDataformRepositoryWorkflowConfigInvocationConfigTransitiveDependenciesIncluded(original["transitive_dependencies_included"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTransitiveDependenciesIncluded); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["transitiveDependenciesIncluded"] = transformedTransitiveDependenciesIncluded
	}

	transformedTransitiveDependentsIncluded, err := expandDataformRepositoryWorkflowConfigInvocationConfigTransitiveDependentsIncluded(original["transitive_dependents_included"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTransitiveDependentsIncluded); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["transitiveDependentsIncluded"] = transformedTransitiveDependentsIncluded
	}

	transformedFullyRefreshIncrementalTablesEnabled, err := expandDataformRepositoryWorkflowConfigInvocationConfigFullyRefreshIncrementalTablesEnabled(original["fully_refresh_incremental_tables_enabled"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFullyRefreshIncrementalTablesEnabled); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["fullyRefreshIncrementalTablesEnabled"] = transformedFullyRefreshIncrementalTablesEnabled
	}

	transformedServiceAccount, err := expandDataformRepositoryWorkflowConfigInvocationConfigServiceAccount(original["service_account"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedServiceAccount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["serviceAccount"] = transformedServiceAccount
	}

	return transformed, nil
}

func expandDataformRepositoryWorkflowConfigInvocationConfigIncludedTargets(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedDatabase, err := expandDataformRepositoryWorkflowConfigInvocationConfigIncludedTargetsDatabase(original["database"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDatabase); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["database"] = transformedDatabase
		}

		transformedSchema, err := expandDataformRepositoryWorkflowConfigInvocationConfigIncludedTargetsSchema(original["schema"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedSchema); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["schema"] = transformedSchema
		}

		transformedName, err := expandDataformRepositoryWorkflowConfigInvocationConfigIncludedTargetsName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["name"] = transformedName
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDataformRepositoryWorkflowConfigInvocationConfigIncludedTargetsDatabase(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataformRepositoryWorkflowConfigInvocationConfigIncludedTargetsSchema(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataformRepositoryWorkflowConfigInvocationConfigIncludedTargetsName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataformRepositoryWorkflowConfigInvocationConfigIncludedTags(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataformRepositoryWorkflowConfigInvocationConfigTransitiveDependenciesIncluded(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataformRepositoryWorkflowConfigInvocationConfigTransitiveDependentsIncluded(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataformRepositoryWorkflowConfigInvocationConfigFullyRefreshIncrementalTablesEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataformRepositoryWorkflowConfigInvocationConfigServiceAccount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataformRepositoryWorkflowConfigCronSchedule(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataformRepositoryWorkflowConfigTimeZone(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}