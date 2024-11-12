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

package backupdr

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const BackupDRBackupPlanAssetType string = "backupdr.googleapis.com/BackupPlan"

func ResourceConverterBackupDRBackupPlan() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: BackupDRBackupPlanAssetType,
		Convert:   GetBackupDRBackupPlanCaiObject,
	}
}

func GetBackupDRBackupPlanCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//backupdr.googleapis.com/projects/{{project}}/locations/{{location}}/backupPlans/{{backup_plan_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetBackupDRBackupPlanApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: BackupDRBackupPlanAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/backupdr/v1/rest",
				DiscoveryName:        "BackupPlan",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetBackupDRBackupPlanApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandBackupDRBackupPlanDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); ok || !reflect.DeepEqual(v, descriptionProp) {
		obj["description"] = descriptionProp
	}
	backupVaultProp, err := expandBackupDRBackupPlanBackupVault(d.Get("backup_vault"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("backup_vault"); !tpgresource.IsEmptyValue(reflect.ValueOf(backupVaultProp)) && (ok || !reflect.DeepEqual(v, backupVaultProp)) {
		obj["backupVault"] = backupVaultProp
	}
	resourceTypeProp, err := expandBackupDRBackupPlanResourceType(d.Get("resource_type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("resource_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(resourceTypeProp)) && (ok || !reflect.DeepEqual(v, resourceTypeProp)) {
		obj["resourceType"] = resourceTypeProp
	}
	backupRulesProp, err := expandBackupDRBackupPlanBackupRules(d.Get("backup_rules"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("backup_rules"); !tpgresource.IsEmptyValue(reflect.ValueOf(backupRulesProp)) && (ok || !reflect.DeepEqual(v, backupRulesProp)) {
		obj["backupRules"] = backupRulesProp
	}

	return obj, nil
}

func expandBackupDRBackupPlanDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBackupDRBackupPlanBackupVault(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBackupDRBackupPlanResourceType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBackupDRBackupPlanBackupRules(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedRuleId, err := expandBackupDRBackupPlanBackupRulesRuleId(original["rule_id"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedRuleId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["ruleId"] = transformedRuleId
		}

		transformedBackupRetentionDays, err := expandBackupDRBackupPlanBackupRulesBackupRetentionDays(original["backup_retention_days"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedBackupRetentionDays); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["backupRetentionDays"] = transformedBackupRetentionDays
		}

		transformedStandardSchedule, err := expandBackupDRBackupPlanBackupRulesStandardSchedule(original["standard_schedule"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedStandardSchedule); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["standardSchedule"] = transformedStandardSchedule
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandBackupDRBackupPlanBackupRulesRuleId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBackupDRBackupPlanBackupRulesBackupRetentionDays(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBackupDRBackupPlanBackupRulesStandardSchedule(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRecurrenceType, err := expandBackupDRBackupPlanBackupRulesStandardScheduleRecurrenceType(original["recurrence_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRecurrenceType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["recurrenceType"] = transformedRecurrenceType
	}

	transformedHourlyFrequency, err := expandBackupDRBackupPlanBackupRulesStandardScheduleHourlyFrequency(original["hourly_frequency"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHourlyFrequency); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["hourlyFrequency"] = transformedHourlyFrequency
	}

	transformedDaysOfWeek, err := expandBackupDRBackupPlanBackupRulesStandardScheduleDaysOfWeek(original["days_of_week"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDaysOfWeek); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["daysOfWeek"] = transformedDaysOfWeek
	}

	transformedDaysOfMonth, err := expandBackupDRBackupPlanBackupRulesStandardScheduleDaysOfMonth(original["days_of_month"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDaysOfMonth); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["daysOfMonth"] = transformedDaysOfMonth
	}

	transformedWeekDayOfMonth, err := expandBackupDRBackupPlanBackupRulesStandardScheduleWeekDayOfMonth(original["week_day_of_month"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedWeekDayOfMonth); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["weekDayOfMonth"] = transformedWeekDayOfMonth
	}

	transformedMonths, err := expandBackupDRBackupPlanBackupRulesStandardScheduleMonths(original["months"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMonths); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["months"] = transformedMonths
	}

	transformedTimeZone, err := expandBackupDRBackupPlanBackupRulesStandardScheduleTimeZone(original["time_zone"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTimeZone); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["timeZone"] = transformedTimeZone
	}

	transformedBackupWindow, err := expandBackupDRBackupPlanBackupRulesStandardScheduleBackupWindow(original["backup_window"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBackupWindow); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["backupWindow"] = transformedBackupWindow
	}

	return transformed, nil
}

func expandBackupDRBackupPlanBackupRulesStandardScheduleRecurrenceType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBackupDRBackupPlanBackupRulesStandardScheduleHourlyFrequency(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBackupDRBackupPlanBackupRulesStandardScheduleDaysOfWeek(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBackupDRBackupPlanBackupRulesStandardScheduleDaysOfMonth(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBackupDRBackupPlanBackupRulesStandardScheduleWeekDayOfMonth(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedWeekOfMonth, err := expandBackupDRBackupPlanBackupRulesStandardScheduleWeekDayOfMonthWeekOfMonth(original["week_of_month"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedWeekOfMonth); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["weekOfMonth"] = transformedWeekOfMonth
	}

	transformedDayOfWeek, err := expandBackupDRBackupPlanBackupRulesStandardScheduleWeekDayOfMonthDayOfWeek(original["day_of_week"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDayOfWeek); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["dayOfWeek"] = transformedDayOfWeek
	}

	return transformed, nil
}

func expandBackupDRBackupPlanBackupRulesStandardScheduleWeekDayOfMonthWeekOfMonth(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBackupDRBackupPlanBackupRulesStandardScheduleWeekDayOfMonthDayOfWeek(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBackupDRBackupPlanBackupRulesStandardScheduleMonths(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBackupDRBackupPlanBackupRulesStandardScheduleTimeZone(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBackupDRBackupPlanBackupRulesStandardScheduleBackupWindow(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedStartHourOfDay, err := expandBackupDRBackupPlanBackupRulesStandardScheduleBackupWindowStartHourOfDay(original["start_hour_of_day"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["startHourOfDay"] = transformedStartHourOfDay
	}

	transformedEndHourOfDay, err := expandBackupDRBackupPlanBackupRulesStandardScheduleBackupWindowEndHourOfDay(original["end_hour_of_day"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEndHourOfDay); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["endHourOfDay"] = transformedEndHourOfDay
	}

	return transformed, nil
}

func expandBackupDRBackupPlanBackupRulesStandardScheduleBackupWindowStartHourOfDay(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBackupDRBackupPlanBackupRulesStandardScheduleBackupWindowEndHourOfDay(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}