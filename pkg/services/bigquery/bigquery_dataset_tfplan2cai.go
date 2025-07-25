// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/bigquery/Dataset.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc_next/tfplan2cai/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package bigquery

import (
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/pkg/caiasset"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/pkg/tfplan2cai/converters/cai"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/pkg/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v6/pkg/transport"
)

func BigQueryDatasetTfplan2caiConverter() cai.Tfplan2caiConverter {
	return cai.Tfplan2caiConverter{
		Convert: GetBigQueryDatasetCaiAssets,
	}
}

func GetBigQueryDatasetCaiAssets(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]caiasset.Asset, error) {
	name, err := cai.AssetName(d, config, "//bigquery.googleapis.com/projects/{{project}}/datasets/{{dataset_id}}")
	if err != nil {
		return []caiasset.Asset{}, err
	}
	if obj, err := GetBigQueryDatasetCaiObject(d, config); err == nil {
		location, _ := tpgresource.GetLocation(d, config)
		if location == "" {
			location = "global"
		}
		return []caiasset.Asset{
			{
				Name: name,
				Type: BigQueryDatasetAssetType,
				Resource: &caiasset.AssetResource{
					Version:              "v2",
					DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/bigquery/v2/rest",
					DiscoveryName:        "Dataset",
					Data:                 obj,
					Location:             location,
				},
			},
		}, nil
	} else {
		return []caiasset.Asset{}, err
	}
}

func GetBigQueryDatasetCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	maxTimeTravelHoursProp, err := expandBigQueryDatasetMaxTimeTravelHours(d.Get("max_time_travel_hours"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("max_time_travel_hours"); !tpgresource.IsEmptyValue(reflect.ValueOf(maxTimeTravelHoursProp)) && (ok || !reflect.DeepEqual(v, maxTimeTravelHoursProp)) {
		obj["maxTimeTravelHours"] = maxTimeTravelHoursProp
	}
	accessProp, err := expandBigQueryDatasetAccess(d.Get("access"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("access"); !tpgresource.IsEmptyValue(reflect.ValueOf(accessProp)) && (ok || !reflect.DeepEqual(v, accessProp)) {
		obj["access"] = accessProp
	}
	datasetReferenceProp, err := expandBigQueryDatasetDatasetReference(nil, d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("dataset_reference"); !tpgresource.IsEmptyValue(reflect.ValueOf(datasetReferenceProp)) && (ok || !reflect.DeepEqual(v, datasetReferenceProp)) {
		obj["datasetReference"] = datasetReferenceProp
	}
	defaultTableExpirationMsProp, err := expandBigQueryDatasetDefaultTableExpirationMs(d.Get("default_table_expiration_ms"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("default_table_expiration_ms"); !tpgresource.IsEmptyValue(reflect.ValueOf(defaultTableExpirationMsProp)) && (ok || !reflect.DeepEqual(v, defaultTableExpirationMsProp)) {
		obj["defaultTableExpirationMs"] = defaultTableExpirationMsProp
	}
	defaultPartitionExpirationMsProp, err := expandBigQueryDatasetDefaultPartitionExpirationMs(d.Get("default_partition_expiration_ms"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("default_partition_expiration_ms"); !tpgresource.IsEmptyValue(reflect.ValueOf(defaultPartitionExpirationMsProp)) && (ok || !reflect.DeepEqual(v, defaultPartitionExpirationMsProp)) {
		obj["defaultPartitionExpirationMs"] = defaultPartitionExpirationMsProp
	}
	descriptionProp, err := expandBigQueryDatasetDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	externalDatasetReferenceProp, err := expandBigQueryDatasetExternalDatasetReference(d.Get("external_dataset_reference"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("external_dataset_reference"); !tpgresource.IsEmptyValue(reflect.ValueOf(externalDatasetReferenceProp)) && (ok || !reflect.DeepEqual(v, externalDatasetReferenceProp)) {
		obj["externalDatasetReference"] = externalDatasetReferenceProp
	}
	friendlyNameProp, err := expandBigQueryDatasetFriendlyName(d.Get("friendly_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("friendly_name"); ok || !reflect.DeepEqual(v, friendlyNameProp) {
		obj["friendlyName"] = friendlyNameProp
	}
	locationProp, err := expandBigQueryDatasetLocation(d.Get("location"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("location"); !tpgresource.IsEmptyValue(reflect.ValueOf(locationProp)) && (ok || !reflect.DeepEqual(v, locationProp)) {
		obj["location"] = locationProp
	}
	defaultEncryptionConfigurationProp, err := expandBigQueryDatasetDefaultEncryptionConfiguration(d.Get("default_encryption_configuration"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("default_encryption_configuration"); !tpgresource.IsEmptyValue(reflect.ValueOf(defaultEncryptionConfigurationProp)) && (ok || !reflect.DeepEqual(v, defaultEncryptionConfigurationProp)) {
		obj["defaultEncryptionConfiguration"] = defaultEncryptionConfigurationProp
	}
	isCaseInsensitiveProp, err := expandBigQueryDatasetIsCaseInsensitive(d.Get("is_case_insensitive"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("is_case_insensitive"); !tpgresource.IsEmptyValue(reflect.ValueOf(isCaseInsensitiveProp)) && (ok || !reflect.DeepEqual(v, isCaseInsensitiveProp)) {
		obj["isCaseInsensitive"] = isCaseInsensitiveProp
	}
	defaultCollationProp, err := expandBigQueryDatasetDefaultCollation(d.Get("default_collation"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("default_collation"); !tpgresource.IsEmptyValue(reflect.ValueOf(defaultCollationProp)) && (ok || !reflect.DeepEqual(v, defaultCollationProp)) {
		obj["defaultCollation"] = defaultCollationProp
	}
	storageBillingModelProp, err := expandBigQueryDatasetStorageBillingModel(d.Get("storage_billing_model"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("storage_billing_model"); !tpgresource.IsEmptyValue(reflect.ValueOf(storageBillingModelProp)) && (ok || !reflect.DeepEqual(v, storageBillingModelProp)) {
		obj["storageBillingModel"] = storageBillingModelProp
	}
	resourceTagsProp, err := expandBigQueryDatasetResourceTags(d.Get("resource_tags"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("resource_tags"); !tpgresource.IsEmptyValue(reflect.ValueOf(resourceTagsProp)) && (ok || !reflect.DeepEqual(v, resourceTagsProp)) {
		obj["resourceTags"] = resourceTagsProp
	}
	externalCatalogDatasetOptionsProp, err := expandBigQueryDatasetExternalCatalogDatasetOptions(d.Get("external_catalog_dataset_options"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("external_catalog_dataset_options"); !tpgresource.IsEmptyValue(reflect.ValueOf(externalCatalogDatasetOptionsProp)) && (ok || !reflect.DeepEqual(v, externalCatalogDatasetOptionsProp)) {
		obj["externalCatalogDatasetOptions"] = externalCatalogDatasetOptionsProp
	}
	labelsProp, err := expandBigQueryDatasetEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandBigQueryDatasetMaxTimeTravelHours(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetAccess(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedDomain, err := expandBigQueryDatasetAccessDomain(original["domain"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDomain); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["domain"] = transformedDomain
		}

		transformedGroupByEmail, err := expandBigQueryDatasetAccessGroupByEmail(original["group_by_email"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedGroupByEmail); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["groupByEmail"] = transformedGroupByEmail
		}

		transformedRole, err := expandBigQueryDatasetAccessRole(original["role"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedRole); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["role"] = transformedRole
		}

		transformedSpecialGroup, err := expandBigQueryDatasetAccessSpecialGroup(original["special_group"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedSpecialGroup); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["specialGroup"] = transformedSpecialGroup
		}

		transformedIamMember, err := expandBigQueryDatasetAccessIamMember(original["iam_member"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedIamMember); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["iamMember"] = transformedIamMember
		}

		transformedUserByEmail, err := expandBigQueryDatasetAccessUserByEmail(original["user_by_email"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedUserByEmail); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["userByEmail"] = transformedUserByEmail
		}

		transformedView, err := expandBigQueryDatasetAccessView(original["view"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedView); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["view"] = transformedView
		}

		transformedDataset, err := expandBigQueryDatasetAccessDataset(original["dataset"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDataset); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["dataset"] = transformedDataset
		}

		transformedRoutine, err := expandBigQueryDatasetAccessRoutine(original["routine"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedRoutine); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["routine"] = transformedRoutine
		}

		transformedCondition, err := expandBigQueryDatasetAccessCondition(original["condition"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedCondition); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["condition"] = transformedCondition
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandBigQueryDatasetAccessDomain(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetAccessGroupByEmail(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetAccessRole(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetAccessSpecialGroup(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetAccessIamMember(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetAccessUserByEmail(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetAccessView(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDatasetId, err := expandBigQueryDatasetAccessViewDatasetId(original["dataset_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDatasetId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["datasetId"] = transformedDatasetId
	}

	transformedProjectId, err := expandBigQueryDatasetAccessViewProjectId(original["project_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProjectId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["projectId"] = transformedProjectId
	}

	transformedTableId, err := expandBigQueryDatasetAccessViewTableId(original["table_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTableId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["tableId"] = transformedTableId
	}

	return transformed, nil
}

func expandBigQueryDatasetAccessViewDatasetId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetAccessViewProjectId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetAccessViewTableId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetAccessDataset(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDataset, err := expandBigQueryDatasetAccessDatasetDataset(original["dataset"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDataset); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["dataset"] = transformedDataset
	}

	transformedTargetTypes, err := expandBigQueryDatasetAccessDatasetTargetTypes(original["target_types"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTargetTypes); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["targetTypes"] = transformedTargetTypes
	}

	return transformed, nil
}

func expandBigQueryDatasetAccessDatasetDataset(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDatasetId, err := expandBigQueryDatasetAccessDatasetDatasetDatasetId(original["dataset_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDatasetId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["datasetId"] = transformedDatasetId
	}

	transformedProjectId, err := expandBigQueryDatasetAccessDatasetDatasetProjectId(original["project_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProjectId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["projectId"] = transformedProjectId
	}

	return transformed, nil
}

func expandBigQueryDatasetAccessDatasetDatasetDatasetId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetAccessDatasetDatasetProjectId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetAccessDatasetTargetTypes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetAccessRoutine(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDatasetId, err := expandBigQueryDatasetAccessRoutineDatasetId(original["dataset_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDatasetId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["datasetId"] = transformedDatasetId
	}

	transformedProjectId, err := expandBigQueryDatasetAccessRoutineProjectId(original["project_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProjectId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["projectId"] = transformedProjectId
	}

	transformedRoutineId, err := expandBigQueryDatasetAccessRoutineRoutineId(original["routine_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRoutineId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["routineId"] = transformedRoutineId
	}

	return transformed, nil
}

func expandBigQueryDatasetAccessRoutineDatasetId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetAccessRoutineProjectId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetAccessRoutineRoutineId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetAccessCondition(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedExpression, err := expandBigQueryDatasetAccessConditionExpression(original["expression"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExpression); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["expression"] = transformedExpression
	}

	transformedTitle, err := expandBigQueryDatasetAccessConditionTitle(original["title"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTitle); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["title"] = transformedTitle
	}

	transformedDescription, err := expandBigQueryDatasetAccessConditionDescription(original["description"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["description"] = transformedDescription
	}

	transformedLocation, err := expandBigQueryDatasetAccessConditionLocation(original["location"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLocation); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["location"] = transformedLocation
	}

	return transformed, nil
}

func expandBigQueryDatasetAccessConditionExpression(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetAccessConditionTitle(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetAccessConditionDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetAccessConditionLocation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetDatasetReference(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	transformed := make(map[string]interface{})
	transformedDatasetId, err := expandBigQueryDatasetDatasetReferenceDatasetId(d.Get("dataset_id"), d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDatasetId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["datasetId"] = transformedDatasetId
	}

	return transformed, nil
}

func expandBigQueryDatasetDatasetReferenceDatasetId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetDefaultTableExpirationMs(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetDefaultPartitionExpirationMs(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetExternalDatasetReference(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedExternalSource, err := expandBigQueryDatasetExternalDatasetReferenceExternalSource(original["external_source"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExternalSource); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["externalSource"] = transformedExternalSource
	}

	transformedConnection, err := expandBigQueryDatasetExternalDatasetReferenceConnection(original["connection"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedConnection); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["connection"] = transformedConnection
	}

	return transformed, nil
}

func expandBigQueryDatasetExternalDatasetReferenceExternalSource(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetExternalDatasetReferenceConnection(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetFriendlyName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetLocation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetDefaultEncryptionConfiguration(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedKmsKeyName, err := expandBigQueryDatasetDefaultEncryptionConfigurationKmsKeyName(original["kms_key_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKeyName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kmsKeyName"] = transformedKmsKeyName
	}

	return transformed, nil
}

func expandBigQueryDatasetDefaultEncryptionConfigurationKmsKeyName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetIsCaseInsensitive(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetDefaultCollation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetStorageBillingModel(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetResourceTags(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandBigQueryDatasetExternalCatalogDatasetOptions(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedParameters, err := expandBigQueryDatasetExternalCatalogDatasetOptionsParameters(original["parameters"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedParameters); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["parameters"] = transformedParameters
	}

	transformedDefaultStorageLocationUri, err := expandBigQueryDatasetExternalCatalogDatasetOptionsDefaultStorageLocationUri(original["default_storage_location_uri"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDefaultStorageLocationUri); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["defaultStorageLocationUri"] = transformedDefaultStorageLocationUri
	}

	return transformed, nil
}

func expandBigQueryDatasetExternalCatalogDatasetOptionsParameters(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandBigQueryDatasetExternalCatalogDatasetOptionsDefaultStorageLocationUri(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigQueryDatasetEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
