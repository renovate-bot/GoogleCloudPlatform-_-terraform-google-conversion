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

package oracledatabase

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const OracleDatabaseCloudVmClusterAssetType string = "oracledatabase.googleapis.com/CloudVmCluster"

func ResourceConverterOracleDatabaseCloudVmCluster() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: OracleDatabaseCloudVmClusterAssetType,
		Convert:   GetOracleDatabaseCloudVmClusterCaiObject,
	}
}

func GetOracleDatabaseCloudVmClusterCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//oracledatabase.googleapis.com/projects/{{project}}/locations/{{location}}/cloudVmClusters/{{cloud_vm_cluster_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetOracleDatabaseCloudVmClusterApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: OracleDatabaseCloudVmClusterAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/oracledatabase/v1/rest",
				DiscoveryName:        "CloudVmCluster",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetOracleDatabaseCloudVmClusterApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	exadataInfrastructureProp, err := expandOracleDatabaseCloudVmClusterExadataInfrastructure(d.Get("exadata_infrastructure"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("exadata_infrastructure"); !tpgresource.IsEmptyValue(reflect.ValueOf(exadataInfrastructureProp)) && (ok || !reflect.DeepEqual(v, exadataInfrastructureProp)) {
		obj["exadataInfrastructure"] = exadataInfrastructureProp
	}
	displayNameProp, err := expandOracleDatabaseCloudVmClusterDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	propertiesProp, err := expandOracleDatabaseCloudVmClusterProperties(d.Get("properties"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("properties"); !tpgresource.IsEmptyValue(reflect.ValueOf(propertiesProp)) && (ok || !reflect.DeepEqual(v, propertiesProp)) {
		obj["properties"] = propertiesProp
	}
	cidrProp, err := expandOracleDatabaseCloudVmClusterCidr(d.Get("cidr"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("cidr"); !tpgresource.IsEmptyValue(reflect.ValueOf(cidrProp)) && (ok || !reflect.DeepEqual(v, cidrProp)) {
		obj["cidr"] = cidrProp
	}
	backupSubnetCidrProp, err := expandOracleDatabaseCloudVmClusterBackupSubnetCidr(d.Get("backup_subnet_cidr"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("backup_subnet_cidr"); !tpgresource.IsEmptyValue(reflect.ValueOf(backupSubnetCidrProp)) && (ok || !reflect.DeepEqual(v, backupSubnetCidrProp)) {
		obj["backupSubnetCidr"] = backupSubnetCidrProp
	}
	networkProp, err := expandOracleDatabaseCloudVmClusterNetwork(d.Get("network"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("network"); !tpgresource.IsEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	labelsProp, err := expandOracleDatabaseCloudVmClusterEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandOracleDatabaseCloudVmClusterExadataInfrastructure(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudVmClusterDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudVmClusterProperties(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedOcid, err := expandOracleDatabaseCloudVmClusterPropertiesOcid(original["ocid"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOcid); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["ocid"] = transformedOcid
	}

	transformedLicenseType, err := expandOracleDatabaseCloudVmClusterPropertiesLicenseType(original["license_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLicenseType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["licenseType"] = transformedLicenseType
	}

	transformedGiVersion, err := expandOracleDatabaseCloudVmClusterPropertiesGiVersion(original["gi_version"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGiVersion); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["giVersion"] = transformedGiVersion
	}

	transformedTimeZone, err := expandOracleDatabaseCloudVmClusterPropertiesTimeZone(original["time_zone"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTimeZone); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["timeZone"] = transformedTimeZone
	}

	transformedSshPublicKeys, err := expandOracleDatabaseCloudVmClusterPropertiesSshPublicKeys(original["ssh_public_keys"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSshPublicKeys); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["sshPublicKeys"] = transformedSshPublicKeys
	}

	transformedNodeCount, err := expandOracleDatabaseCloudVmClusterPropertiesNodeCount(original["node_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNodeCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["nodeCount"] = transformedNodeCount
	}

	transformedShape, err := expandOracleDatabaseCloudVmClusterPropertiesShape(original["shape"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedShape); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["shape"] = transformedShape
	}

	transformedOcpuCount, err := expandOracleDatabaseCloudVmClusterPropertiesOcpuCount(original["ocpu_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOcpuCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["ocpuCount"] = transformedOcpuCount
	}

	transformedMemorySizeGb, err := expandOracleDatabaseCloudVmClusterPropertiesMemorySizeGb(original["memory_size_gb"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMemorySizeGb); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["memorySizeGb"] = transformedMemorySizeGb
	}

	transformedDbNodeStorageSizeGb, err := expandOracleDatabaseCloudVmClusterPropertiesDbNodeStorageSizeGb(original["db_node_storage_size_gb"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDbNodeStorageSizeGb); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["dbNodeStorageSizeGb"] = transformedDbNodeStorageSizeGb
	}

	transformedStorageSizeGb, err := expandOracleDatabaseCloudVmClusterPropertiesStorageSizeGb(original["storage_size_gb"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStorageSizeGb); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["storageSizeGb"] = transformedStorageSizeGb
	}

	transformedDataStorageSizeTb, err := expandOracleDatabaseCloudVmClusterPropertiesDataStorageSizeTb(original["data_storage_size_tb"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDataStorageSizeTb); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["dataStorageSizeTb"] = transformedDataStorageSizeTb
	}

	transformedDiskRedundancy, err := expandOracleDatabaseCloudVmClusterPropertiesDiskRedundancy(original["disk_redundancy"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDiskRedundancy); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["diskRedundancy"] = transformedDiskRedundancy
	}

	transformedSparseDiskgroupEnabled, err := expandOracleDatabaseCloudVmClusterPropertiesSparseDiskgroupEnabled(original["sparse_diskgroup_enabled"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSparseDiskgroupEnabled); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["sparseDiskgroupEnabled"] = transformedSparseDiskgroupEnabled
	}

	transformedLocalBackupEnabled, err := expandOracleDatabaseCloudVmClusterPropertiesLocalBackupEnabled(original["local_backup_enabled"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLocalBackupEnabled); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["localBackupEnabled"] = transformedLocalBackupEnabled
	}

	transformedHostnamePrefix, err := expandOracleDatabaseCloudVmClusterPropertiesHostnamePrefix(original["hostname_prefix"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHostnamePrefix); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["hostnamePrefix"] = transformedHostnamePrefix
	}

	transformedDiagnosticsDataCollectionOptions, err := expandOracleDatabaseCloudVmClusterPropertiesDiagnosticsDataCollectionOptions(original["diagnostics_data_collection_options"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDiagnosticsDataCollectionOptions); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["diagnosticsDataCollectionOptions"] = transformedDiagnosticsDataCollectionOptions
	}

	transformedState, err := expandOracleDatabaseCloudVmClusterPropertiesState(original["state"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedState); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["state"] = transformedState
	}

	transformedScanListenerPortTcp, err := expandOracleDatabaseCloudVmClusterPropertiesScanListenerPortTcp(original["scan_listener_port_tcp"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedScanListenerPortTcp); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["scanListenerPortTcp"] = transformedScanListenerPortTcp
	}

	transformedScanListenerPortTcpSsl, err := expandOracleDatabaseCloudVmClusterPropertiesScanListenerPortTcpSsl(original["scan_listener_port_tcp_ssl"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedScanListenerPortTcpSsl); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["scanListenerPortTcpSsl"] = transformedScanListenerPortTcpSsl
	}

	transformedDomain, err := expandOracleDatabaseCloudVmClusterPropertiesDomain(original["domain"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDomain); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["domain"] = transformedDomain
	}

	transformedScanDns, err := expandOracleDatabaseCloudVmClusterPropertiesScanDns(original["scan_dns"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedScanDns); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["scanDns"] = transformedScanDns
	}

	transformedHostname, err := expandOracleDatabaseCloudVmClusterPropertiesHostname(original["hostname"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHostname); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["hostname"] = transformedHostname
	}

	transformedCpuCoreCount, err := expandOracleDatabaseCloudVmClusterPropertiesCpuCoreCount(original["cpu_core_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCpuCoreCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["cpuCoreCount"] = transformedCpuCoreCount
	}

	transformedSystemVersion, err := expandOracleDatabaseCloudVmClusterPropertiesSystemVersion(original["system_version"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSystemVersion); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["systemVersion"] = transformedSystemVersion
	}

	transformedScanIpIds, err := expandOracleDatabaseCloudVmClusterPropertiesScanIpIds(original["scan_ip_ids"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedScanIpIds); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["scanIpIds"] = transformedScanIpIds
	}

	transformedScanDnsRecordId, err := expandOracleDatabaseCloudVmClusterPropertiesScanDnsRecordId(original["scan_dns_record_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedScanDnsRecordId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["scanDnsRecordId"] = transformedScanDnsRecordId
	}

	transformedOciUrl, err := expandOracleDatabaseCloudVmClusterPropertiesOciUrl(original["oci_url"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOciUrl); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["ociUrl"] = transformedOciUrl
	}

	transformedDbServerOcids, err := expandOracleDatabaseCloudVmClusterPropertiesDbServerOcids(original["db_server_ocids"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDbServerOcids); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["dbServerOcids"] = transformedDbServerOcids
	}

	transformedCompartmentId, err := expandOracleDatabaseCloudVmClusterPropertiesCompartmentId(original["compartment_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCompartmentId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["compartmentId"] = transformedCompartmentId
	}

	transformedDnsListenerIp, err := expandOracleDatabaseCloudVmClusterPropertiesDnsListenerIp(original["dns_listener_ip"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDnsListenerIp); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["dnsListenerIp"] = transformedDnsListenerIp
	}

	transformedClusterName, err := expandOracleDatabaseCloudVmClusterPropertiesClusterName(original["cluster_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedClusterName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["clusterName"] = transformedClusterName
	}

	return transformed, nil
}

func expandOracleDatabaseCloudVmClusterPropertiesOcid(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudVmClusterPropertiesLicenseType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudVmClusterPropertiesGiVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudVmClusterPropertiesTimeZone(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedId, err := expandOracleDatabaseCloudVmClusterPropertiesTimeZoneId(original["id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["id"] = transformedId
	}

	return transformed, nil
}

func expandOracleDatabaseCloudVmClusterPropertiesTimeZoneId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudVmClusterPropertiesSshPublicKeys(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudVmClusterPropertiesNodeCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudVmClusterPropertiesShape(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudVmClusterPropertiesOcpuCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudVmClusterPropertiesMemorySizeGb(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudVmClusterPropertiesDbNodeStorageSizeGb(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudVmClusterPropertiesStorageSizeGb(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudVmClusterPropertiesDataStorageSizeTb(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudVmClusterPropertiesDiskRedundancy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudVmClusterPropertiesSparseDiskgroupEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudVmClusterPropertiesLocalBackupEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudVmClusterPropertiesHostnamePrefix(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudVmClusterPropertiesDiagnosticsDataCollectionOptions(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDiagnosticsEventsEnabled, err := expandOracleDatabaseCloudVmClusterPropertiesDiagnosticsDataCollectionOptionsDiagnosticsEventsEnabled(original["diagnostics_events_enabled"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDiagnosticsEventsEnabled); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["diagnosticsEventsEnabled"] = transformedDiagnosticsEventsEnabled
	}

	transformedHealthMonitoringEnabled, err := expandOracleDatabaseCloudVmClusterPropertiesDiagnosticsDataCollectionOptionsHealthMonitoringEnabled(original["health_monitoring_enabled"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHealthMonitoringEnabled); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["healthMonitoringEnabled"] = transformedHealthMonitoringEnabled
	}

	transformedIncidentLogsEnabled, err := expandOracleDatabaseCloudVmClusterPropertiesDiagnosticsDataCollectionOptionsIncidentLogsEnabled(original["incident_logs_enabled"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIncidentLogsEnabled); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["incidentLogsEnabled"] = transformedIncidentLogsEnabled
	}

	return transformed, nil
}

func expandOracleDatabaseCloudVmClusterPropertiesDiagnosticsDataCollectionOptionsDiagnosticsEventsEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudVmClusterPropertiesDiagnosticsDataCollectionOptionsHealthMonitoringEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudVmClusterPropertiesDiagnosticsDataCollectionOptionsIncidentLogsEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudVmClusterPropertiesState(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudVmClusterPropertiesScanListenerPortTcp(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudVmClusterPropertiesScanListenerPortTcpSsl(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudVmClusterPropertiesDomain(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudVmClusterPropertiesScanDns(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudVmClusterPropertiesHostname(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudVmClusterPropertiesCpuCoreCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudVmClusterPropertiesSystemVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudVmClusterPropertiesScanIpIds(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudVmClusterPropertiesScanDnsRecordId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudVmClusterPropertiesOciUrl(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudVmClusterPropertiesDbServerOcids(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudVmClusterPropertiesCompartmentId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudVmClusterPropertiesDnsListenerIp(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudVmClusterPropertiesClusterName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudVmClusterCidr(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudVmClusterBackupSubnetCidr(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudVmClusterNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudVmClusterEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}