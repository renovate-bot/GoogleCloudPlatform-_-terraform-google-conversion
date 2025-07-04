// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/bigquery/Dataset.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc_next/cai2hcl/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package bigquery

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/pkg/cai2hcl/converters/utils"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/pkg/cai2hcl/models"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/pkg/caiasset"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/pkg/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v6/pkg/transport"
)

func bigqueryDatasetAccessSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"condition": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `Condition for the binding. If CEL expression in this field is true, this
access binding will be considered.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"expression": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `Textual representation of an expression in Common Expression Language syntax.`,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
							Description: `Description of the expression. This is a longer text which describes the expression,
e.g. when hovered over it in a UI.`,
						},
						"location": {
							Type:     schema.TypeString,
							Optional: true,
							Description: `String indicating the location of the expression for error reporting, e.g. a file
name and a position in the file.`,
						},
						"title": {
							Type:     schema.TypeString,
							Optional: true,
							Description: `Title for the expression, i.e. a short string describing its purpose.
This can be used e.g. in UIs which allow to enter the expression.`,
						},
					},
				},
			},
			"dataset": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `Grants all resources of particular types in a particular dataset read access to the current dataset.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dataset": {
							Type:        schema.TypeList,
							Required:    true,
							Description: `The dataset this entry applies to`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dataset_id": {
										Type:        schema.TypeString,
										Required:    true,
										Description: `The ID of the dataset containing this table.`,
									},
									"project_id": {
										Type:        schema.TypeString,
										Required:    true,
										Description: `The ID of the project containing this table.`,
									},
								},
							},
						},
						"target_types": {
							Type:     schema.TypeList,
							Required: true,
							Description: `Which resources in the dataset this entry applies to. Currently, only views are supported,
but additional target types may be added in the future. Possible values: VIEWS`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"domain": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `A domain to grant access to. Any users signed in with the
domain specified will be granted the specified access`,
			},
			"group_by_email": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `An email address of a Google Group to grant access to.`,
			},
			"iam_member": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `Some other type of member that appears in the IAM Policy but isn't a user,
group, domain, or special group. For example: 'allUsers'`,
			},
			"role": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `Describes the rights granted to the user specified by the other
member of the access object. Basic, predefined, and custom roles
are supported. Predefined roles that have equivalent basic roles
are swapped by the API to their basic counterparts. See
[official docs](https://cloud.google.com/bigquery/docs/access-control).`,
			},
			"routine": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `A routine from a different dataset to grant access to. Queries
executed against that routine will have read access to tables in
this dataset. The role field is not required when this field is
set. If that routine is updated by any user, access to the routine
needs to be granted again via an update operation.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dataset_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `The ID of the dataset containing this table.`,
						},
						"project_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `The ID of the project containing this table.`,
						},
						"routine_id": {
							Type:     schema.TypeString,
							Required: true,
							Description: `The ID of the routine. The ID must contain only letters (a-z,
A-Z), numbers (0-9), or underscores (_). The maximum length
is 256 characters.`,
						},
					},
				},
			},
			"special_group": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `A special group to grant access to. Possible values include:
* 'projectOwners': Owners of the enclosing project.
* 'projectReaders': Readers of the enclosing project.
* 'projectWriters': Writers of the enclosing project.
* 'allAuthenticatedUsers': All authenticated BigQuery users.`,
			},
			"user_by_email": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `An email address of a user to grant access to. For example:
fred@example.com`,
			},
			"view": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `A view from a different dataset to grant access to. Queries
executed against that view will have read access to tables in
this dataset. The role field is not required when this field is
set. If that view is updated by any user, access to the view
needs to be granted again via an update operation.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dataset_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `The ID of the dataset containing this table.`,
						},
						"project_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `The ID of the project containing this table.`,
						},
						"table_id": {
							Type:     schema.TypeString,
							Required: true,
							Description: `The ID of the table. The ID must contain only letters (a-z,
A-Z), numbers (0-9), or underscores (_). The maximum length
is 1,024 characters.`,
						},
					},
				},
			},
		},
	}
}

const BigQueryDatasetAssetType string = "bigquery.googleapis.com/Dataset"

const BigQueryDatasetSchemaName string = "google_bigquery_dataset"

type BigQueryDatasetConverter struct {
	name   string
	schema map[string]*schema.Schema
}

func NewBigQueryDatasetConverter(provider *schema.Provider) models.Converter {
	schema := provider.ResourcesMap[BigQueryDatasetSchemaName].Schema

	return &BigQueryDatasetConverter{
		name:   BigQueryDatasetSchemaName,
		schema: schema,
	}
}

// Convert converts asset to HCL resource blocks.
func (c *BigQueryDatasetConverter) Convert(asset caiasset.Asset) ([]*models.TerraformResourceBlock, error) {
	var blocks []*models.TerraformResourceBlock
	block, err := c.convertResourceData(asset)
	if err != nil {
		return nil, err
	}
	blocks = append(blocks, block)
	return blocks, nil
}

func (c *BigQueryDatasetConverter) convertResourceData(asset caiasset.Asset) (*models.TerraformResourceBlock, error) {
	if asset.Resource == nil || asset.Resource.Data == nil {
		return nil, fmt.Errorf("asset resource data is nil")
	}

	var err error
	res := asset.Resource.Data
	config := utils.NewConfig()
	d := &schema.ResourceData{}

	assetNameParts := strings.Split(asset.Name, "/")
	hclBlockName := assetNameParts[len(assetNameParts)-1]

	hclData := make(map[string]interface{})

	hclData["max_time_travel_hours"] = flattenBigQueryDatasetMaxTimeTravelHours(res["maxTimeTravelHours"], d, config)
	hclData["access"] = flattenBigQueryDatasetAccess(res["access"], d, config)
	// Terraform must set the top level schema field, but since this object contains collapsed properties
	// it's difficult to know what the top level should be. Instead we just loop over the map returned from flatten.
	if flattenedProp := flattenBigQueryDatasetDatasetReference(res["datasetReference"], d, config); flattenedProp != nil {
		flattenedPropSlice, ok := flattenedProp.([]interface{})
		if !ok || len(flattenedPropSlice) == 0 {
			return nil, fmt.Errorf("unexpected type returned from flattener: %T", flattenedProp)
		}
		flattedPropMap, ok := flattenedPropSlice[0].(map[string]interface{})
		if !ok || len(flattedPropMap) == 0 {
			return nil, fmt.Errorf("unexpected type returned from flattener: %T", flattenedPropSlice)
		}
		for k, v := range flattedPropMap {
			hclData[k] = v
		}
	}
	hclData["default_table_expiration_ms"] = flattenBigQueryDatasetDefaultTableExpirationMs(res["defaultTableExpirationMs"], d, config)
	hclData["default_partition_expiration_ms"] = flattenBigQueryDatasetDefaultPartitionExpirationMs(res["defaultPartitionExpirationMs"], d, config)
	hclData["description"] = flattenBigQueryDatasetDescription(res["description"], d, config)
	hclData["external_dataset_reference"] = flattenBigQueryDatasetExternalDatasetReference(res["externalDatasetReference"], d, config)
	hclData["friendly_name"] = flattenBigQueryDatasetFriendlyName(res["friendlyName"], d, config)
	hclData["labels"] = flattenBigQueryDatasetLabels(res["labels"], d, config)
	hclData["location"] = flattenBigQueryDatasetLocation(res["location"], d, config)
	hclData["default_encryption_configuration"] = flattenBigQueryDatasetDefaultEncryptionConfiguration(res["defaultEncryptionConfiguration"], d, config)
	hclData["is_case_insensitive"] = flattenBigQueryDatasetIsCaseInsensitive(res["isCaseInsensitive"], d, config)
	hclData["default_collation"] = flattenBigQueryDatasetDefaultCollation(res["defaultCollation"], d, config)
	hclData["storage_billing_model"] = flattenBigQueryDatasetStorageBillingModel(res["storageBillingModel"], d, config)
	hclData["resource_tags"] = flattenBigQueryDatasetResourceTags(res["resourceTags"], d, config)
	hclData["external_catalog_dataset_options"] = flattenBigQueryDatasetExternalCatalogDatasetOptions(res["externalCatalogDatasetOptions"], d, config)

	ctyVal, err := utils.MapToCtyValWithSchema(hclData, c.schema)
	if err != nil {
		return nil, err
	}
	return &models.TerraformResourceBlock{
		Labels: []string{c.name, hclBlockName},
		Value:  ctyVal,
	}, nil
}

func flattenBigQueryDatasetMaxTimeTravelHours(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigQueryDatasetAccess(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := schema.NewSet(schema.HashResource(bigqueryDatasetAccessSchema()), []interface{}{})
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed.Add(map[string]interface{}{
			"domain":         flattenBigQueryDatasetAccessDomain(original["domain"], d, config),
			"group_by_email": flattenBigQueryDatasetAccessGroupByEmail(original["groupByEmail"], d, config),
			"role":           flattenBigQueryDatasetAccessRole(original["role"], d, config),
			"special_group":  flattenBigQueryDatasetAccessSpecialGroup(original["specialGroup"], d, config),
			"iam_member":     flattenBigQueryDatasetAccessIamMember(original["iamMember"], d, config),
			"user_by_email":  flattenBigQueryDatasetAccessUserByEmail(original["userByEmail"], d, config),
			"view":           flattenBigQueryDatasetAccessView(original["view"], d, config),
			"dataset":        flattenBigQueryDatasetAccessDataset(original["dataset"], d, config),
			"routine":        flattenBigQueryDatasetAccessRoutine(original["routine"], d, config),
			"condition":      flattenBigQueryDatasetAccessCondition(original["condition"], d, config),
		})
	}
	return transformed
}
func flattenBigQueryDatasetAccessDomain(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigQueryDatasetAccessGroupByEmail(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigQueryDatasetAccessRole(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigQueryDatasetAccessSpecialGroup(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigQueryDatasetAccessIamMember(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigQueryDatasetAccessUserByEmail(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigQueryDatasetAccessView(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["dataset_id"] =
		flattenBigQueryDatasetAccessViewDatasetId(original["datasetId"], d, config)
	transformed["project_id"] =
		flattenBigQueryDatasetAccessViewProjectId(original["projectId"], d, config)
	transformed["table_id"] =
		flattenBigQueryDatasetAccessViewTableId(original["tableId"], d, config)
	return []interface{}{transformed}
}
func flattenBigQueryDatasetAccessViewDatasetId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigQueryDatasetAccessViewProjectId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigQueryDatasetAccessViewTableId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigQueryDatasetAccessDataset(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["dataset"] =
		flattenBigQueryDatasetAccessDatasetDataset(original["dataset"], d, config)
	transformed["target_types"] =
		flattenBigQueryDatasetAccessDatasetTargetTypes(original["targetTypes"], d, config)
	return []interface{}{transformed}
}
func flattenBigQueryDatasetAccessDatasetDataset(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["dataset_id"] =
		flattenBigQueryDatasetAccessDatasetDatasetDatasetId(original["datasetId"], d, config)
	transformed["project_id"] =
		flattenBigQueryDatasetAccessDatasetDatasetProjectId(original["projectId"], d, config)
	return []interface{}{transformed}
}
func flattenBigQueryDatasetAccessDatasetDatasetDatasetId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigQueryDatasetAccessDatasetDatasetProjectId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigQueryDatasetAccessDatasetTargetTypes(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigQueryDatasetAccessRoutine(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["dataset_id"] =
		flattenBigQueryDatasetAccessRoutineDatasetId(original["datasetId"], d, config)
	transformed["project_id"] =
		flattenBigQueryDatasetAccessRoutineProjectId(original["projectId"], d, config)
	transformed["routine_id"] =
		flattenBigQueryDatasetAccessRoutineRoutineId(original["routineId"], d, config)
	return []interface{}{transformed}
}
func flattenBigQueryDatasetAccessRoutineDatasetId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigQueryDatasetAccessRoutineProjectId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigQueryDatasetAccessRoutineRoutineId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigQueryDatasetAccessCondition(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["expression"] =
		flattenBigQueryDatasetAccessConditionExpression(original["expression"], d, config)
	transformed["title"] =
		flattenBigQueryDatasetAccessConditionTitle(original["title"], d, config)
	transformed["description"] =
		flattenBigQueryDatasetAccessConditionDescription(original["description"], d, config)
	transformed["location"] =
		flattenBigQueryDatasetAccessConditionLocation(original["location"], d, config)
	return []interface{}{transformed}
}
func flattenBigQueryDatasetAccessConditionExpression(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigQueryDatasetAccessConditionTitle(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigQueryDatasetAccessConditionDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigQueryDatasetAccessConditionLocation(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigQueryDatasetDatasetReference(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["dataset_id"] =
		flattenBigQueryDatasetDatasetReferenceDatasetId(original["datasetId"], d, config)
	return []interface{}{transformed}
}
func flattenBigQueryDatasetDatasetReferenceDatasetId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigQueryDatasetDefaultTableExpirationMs(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenBigQueryDatasetDefaultPartitionExpirationMs(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenBigQueryDatasetDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigQueryDatasetExternalDatasetReference(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["external_source"] =
		flattenBigQueryDatasetExternalDatasetReferenceExternalSource(original["externalSource"], d, config)
	transformed["connection"] =
		flattenBigQueryDatasetExternalDatasetReferenceConnection(original["connection"], d, config)
	return []interface{}{transformed}
}
func flattenBigQueryDatasetExternalDatasetReferenceExternalSource(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigQueryDatasetExternalDatasetReferenceConnection(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigQueryDatasetFriendlyName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigQueryDatasetLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return utils.RemoveTerraformAttributionLabel(v)
}

// Older Datasets in BigQuery have no Location set in the API response. This may be an issue when importing
// datasets created before BigQuery was available in multiple zones. We can safely assume that these datasets
// are in the US, as this was the default at the time.
func flattenBigQueryDatasetLocation(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return "US"
	}
	return v
}

func flattenBigQueryDatasetDefaultEncryptionConfiguration(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["kms_key_name"] =
		flattenBigQueryDatasetDefaultEncryptionConfigurationKmsKeyName(original["kmsKeyName"], d, config)
	return []interface{}{transformed}
}
func flattenBigQueryDatasetDefaultEncryptionConfigurationKmsKeyName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigQueryDatasetIsCaseInsensitive(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigQueryDatasetDefaultCollation(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigQueryDatasetStorageBillingModel(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigQueryDatasetResourceTags(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigQueryDatasetExternalCatalogDatasetOptions(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["parameters"] =
		flattenBigQueryDatasetExternalCatalogDatasetOptionsParameters(original["parameters"], d, config)
	transformed["default_storage_location_uri"] =
		flattenBigQueryDatasetExternalCatalogDatasetOptionsDefaultStorageLocationUri(original["defaultStorageLocationUri"], d, config)
	return []interface{}{transformed}
}
func flattenBigQueryDatasetExternalCatalogDatasetOptionsParameters(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenBigQueryDatasetExternalCatalogDatasetOptionsDefaultStorageLocationUri(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}
