// ----------------------------------------------------------------------------
//
//	***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//	This file is automatically generated by Magic Modules and manual
//	changes will be clobbered when the file is regenerated.
//
//	Please read more about how to change this file in
//	.github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------
package converters

import (
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/pkg/cai2hcl/converters/services/compute"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/pkg/cai2hcl/converters/services/resourcemanager"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/pkg/cai2hcl/models"

	tpg_provider "github.com/GoogleCloudPlatform/terraform-google-conversion/v6/pkg/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var provider *schema.Provider = tpg_provider.Provider()

// ConverterMap is a collection of converters instances, indexed by cai asset type.
var ConverterMap = map[string]models.Converter{
	// ####### START handwritten resources ###########
	resourcemanager.ProjectAssetType: resourcemanager.NewProjectConverter(provider),
	compute.ComputeInstanceAssetType: compute.NewComputeInstanceConverter(provider),
	// ####### END handwritten resources ###########
	compute.ComputeAddressAssetType:    compute.NewComputeAddressConverter(provider),
	compute.ComputeAutoscalerAssetType: compute.NewComputeAutoscalerConverter(provider),
	compute.ComputeDiskAssetType:       compute.NewComputeDiskConverter(provider),
}
