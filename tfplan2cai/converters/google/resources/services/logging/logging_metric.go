// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/logging/Metric.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package logging

import (
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const LoggingMetricAssetType string = "logging.googleapis.com/Metric"

func ResourceConverterLoggingMetric() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: LoggingMetricAssetType,
		Convert:   GetLoggingMetricCaiObject,
	}
}

func GetLoggingMetricCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//logging.googleapis.com/projects/{{project}}/metrics/{{%name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetLoggingMetricApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: LoggingMetricAssetType,
			Resource: &cai.AssetResource{
				Version:              "v2",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/logging/v2/rest",
				DiscoveryName:        "Metric",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetLoggingMetricApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandLoggingMetricName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandLoggingMetricDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	bucketNameProp, err := expandLoggingMetricBucketName(d.Get("bucket_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("bucket_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(bucketNameProp)) && (ok || !reflect.DeepEqual(v, bucketNameProp)) {
		obj["bucketName"] = bucketNameProp
	}
	disabledProp, err := expandLoggingMetricDisabled(d.Get("disabled"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("disabled"); !tpgresource.IsEmptyValue(reflect.ValueOf(disabledProp)) && (ok || !reflect.DeepEqual(v, disabledProp)) {
		obj["disabled"] = disabledProp
	}
	filterProp, err := expandLoggingMetricFilter(d.Get("filter"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("filter"); !tpgresource.IsEmptyValue(reflect.ValueOf(filterProp)) && (ok || !reflect.DeepEqual(v, filterProp)) {
		obj["filter"] = filterProp
	}
	metricDescriptorProp, err := expandLoggingMetricMetricDescriptor(d.Get("metric_descriptor"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("metric_descriptor"); !tpgresource.IsEmptyValue(reflect.ValueOf(metricDescriptorProp)) && (ok || !reflect.DeepEqual(v, metricDescriptorProp)) {
		obj["metricDescriptor"] = metricDescriptorProp
	}
	labelExtractorsProp, err := expandLoggingMetricLabelExtractors(d.Get("label_extractors"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("label_extractors"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelExtractorsProp)) && (ok || !reflect.DeepEqual(v, labelExtractorsProp)) {
		obj["labelExtractors"] = labelExtractorsProp
	}
	valueExtractorProp, err := expandLoggingMetricValueExtractor(d.Get("value_extractor"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("value_extractor"); !tpgresource.IsEmptyValue(reflect.ValueOf(valueExtractorProp)) && (ok || !reflect.DeepEqual(v, valueExtractorProp)) {
		obj["valueExtractor"] = valueExtractorProp
	}
	bucketOptionsProp, err := expandLoggingMetricBucketOptions(d.Get("bucket_options"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("bucket_options"); !tpgresource.IsEmptyValue(reflect.ValueOf(bucketOptionsProp)) && (ok || !reflect.DeepEqual(v, bucketOptionsProp)) {
		obj["bucketOptions"] = bucketOptionsProp
	}

	return obj, nil
}

func expandLoggingMetricName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandLoggingMetricDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandLoggingMetricBucketName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandLoggingMetricDisabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandLoggingMetricFilter(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandLoggingMetricMetricDescriptor(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedUnit, err := expandLoggingMetricMetricDescriptorUnit(original["unit"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUnit); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["unit"] = transformedUnit
	}

	transformedValueType, err := expandLoggingMetricMetricDescriptorValueType(original["value_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedValueType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["valueType"] = transformedValueType
	}

	transformedMetricKind, err := expandLoggingMetricMetricDescriptorMetricKind(original["metric_kind"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMetricKind); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["metricKind"] = transformedMetricKind
	}

	transformedLabels, err := expandLoggingMetricMetricDescriptorLabels(original["labels"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLabels); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["labels"] = transformedLabels
	}

	transformedDisplayName, err := expandLoggingMetricMetricDescriptorDisplayName(original["display_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDisplayName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["displayName"] = transformedDisplayName
	}

	return transformed, nil
}

func expandLoggingMetricMetricDescriptorUnit(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandLoggingMetricMetricDescriptorValueType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandLoggingMetricMetricDescriptorMetricKind(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandLoggingMetricMetricDescriptorLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedKey, err := expandLoggingMetricMetricDescriptorLabelsKey(original["key"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedKey); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["key"] = transformedKey
		}

		transformedDescription, err := expandLoggingMetricMetricDescriptorLabelsDescription(original["description"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["description"] = transformedDescription
		}

		transformedValueType, err := expandLoggingMetricMetricDescriptorLabelsValueType(original["value_type"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedValueType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["valueType"] = transformedValueType
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandLoggingMetricMetricDescriptorLabelsKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandLoggingMetricMetricDescriptorLabelsDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandLoggingMetricMetricDescriptorLabelsValueType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandLoggingMetricMetricDescriptorDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandLoggingMetricLabelExtractors(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandLoggingMetricValueExtractor(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandLoggingMetricBucketOptions(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedLinearBuckets, err := expandLoggingMetricBucketOptionsLinearBuckets(original["linear_buckets"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLinearBuckets); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["linearBuckets"] = transformedLinearBuckets
	}

	transformedExponentialBuckets, err := expandLoggingMetricBucketOptionsExponentialBuckets(original["exponential_buckets"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExponentialBuckets); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["exponentialBuckets"] = transformedExponentialBuckets
	}

	transformedExplicitBuckets, err := expandLoggingMetricBucketOptionsExplicitBuckets(original["explicit_buckets"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExplicitBuckets); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["explicitBuckets"] = transformedExplicitBuckets
	}

	return transformed, nil
}

func expandLoggingMetricBucketOptionsLinearBuckets(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedNumFiniteBuckets, err := expandLoggingMetricBucketOptionsLinearBucketsNumFiniteBuckets(original["num_finite_buckets"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNumFiniteBuckets); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["numFiniteBuckets"] = transformedNumFiniteBuckets
	}

	transformedWidth, err := expandLoggingMetricBucketOptionsLinearBucketsWidth(original["width"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedWidth); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["width"] = transformedWidth
	}

	transformedOffset, err := expandLoggingMetricBucketOptionsLinearBucketsOffset(original["offset"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOffset); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["offset"] = transformedOffset
	}

	return transformed, nil
}

func expandLoggingMetricBucketOptionsLinearBucketsNumFiniteBuckets(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandLoggingMetricBucketOptionsLinearBucketsWidth(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandLoggingMetricBucketOptionsLinearBucketsOffset(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandLoggingMetricBucketOptionsExponentialBuckets(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedNumFiniteBuckets, err := expandLoggingMetricBucketOptionsExponentialBucketsNumFiniteBuckets(original["num_finite_buckets"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNumFiniteBuckets); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["numFiniteBuckets"] = transformedNumFiniteBuckets
	}

	transformedGrowthFactor, err := expandLoggingMetricBucketOptionsExponentialBucketsGrowthFactor(original["growth_factor"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGrowthFactor); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["growthFactor"] = transformedGrowthFactor
	}

	transformedScale, err := expandLoggingMetricBucketOptionsExponentialBucketsScale(original["scale"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedScale); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["scale"] = transformedScale
	}

	return transformed, nil
}

func expandLoggingMetricBucketOptionsExponentialBucketsNumFiniteBuckets(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandLoggingMetricBucketOptionsExponentialBucketsGrowthFactor(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandLoggingMetricBucketOptionsExponentialBucketsScale(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandLoggingMetricBucketOptionsExplicitBuckets(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedBounds, err := expandLoggingMetricBucketOptionsExplicitBucketsBounds(original["bounds"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBounds); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["bounds"] = transformedBounds
	}

	return transformed, nil
}

func expandLoggingMetricBucketOptionsExplicitBucketsBounds(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
