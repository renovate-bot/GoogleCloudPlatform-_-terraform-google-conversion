// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/redis/Cluster.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package redis

import (
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const RedisClusterAssetType string = "redis.googleapis.com/Cluster"

func ResourceConverterRedisCluster() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: RedisClusterAssetType,
		Convert:   GetRedisClusterCaiObject,
	}
}

func GetRedisClusterCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//redis.googleapis.com/projects/{{project}}/locations/{{region}}/clusters/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetRedisClusterApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: RedisClusterAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/redis/v1beta1/rest",
				DiscoveryName:        "Cluster",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetRedisClusterApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	gcsSourceProp, err := expandRedisClusterGcsSource(d.Get("gcs_source"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("gcs_source"); !tpgresource.IsEmptyValue(reflect.ValueOf(gcsSourceProp)) && (ok || !reflect.DeepEqual(v, gcsSourceProp)) {
		obj["gcsSource"] = gcsSourceProp
	}
	managedBackupSourceProp, err := expandRedisClusterManagedBackupSource(d.Get("managed_backup_source"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("managed_backup_source"); !tpgresource.IsEmptyValue(reflect.ValueOf(managedBackupSourceProp)) && (ok || !reflect.DeepEqual(v, managedBackupSourceProp)) {
		obj["managedBackupSource"] = managedBackupSourceProp
	}
	automatedBackupConfigProp, err := expandRedisClusterAutomatedBackupConfig(d.Get("automated_backup_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("automated_backup_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(automatedBackupConfigProp)) && (ok || !reflect.DeepEqual(v, automatedBackupConfigProp)) {
		obj["automatedBackupConfig"] = automatedBackupConfigProp
	}
	authorizationModeProp, err := expandRedisClusterAuthorizationMode(d.Get("authorization_mode"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("authorization_mode"); !tpgresource.IsEmptyValue(reflect.ValueOf(authorizationModeProp)) && (ok || !reflect.DeepEqual(v, authorizationModeProp)) {
		obj["authorizationMode"] = authorizationModeProp
	}
	transitEncryptionModeProp, err := expandRedisClusterTransitEncryptionMode(d.Get("transit_encryption_mode"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("transit_encryption_mode"); !tpgresource.IsEmptyValue(reflect.ValueOf(transitEncryptionModeProp)) && (ok || !reflect.DeepEqual(v, transitEncryptionModeProp)) {
		obj["transitEncryptionMode"] = transitEncryptionModeProp
	}
	nodeTypeProp, err := expandRedisClusterNodeType(d.Get("node_type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("node_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(nodeTypeProp)) && (ok || !reflect.DeepEqual(v, nodeTypeProp)) {
		obj["nodeType"] = nodeTypeProp
	}
	zoneDistributionConfigProp, err := expandRedisClusterZoneDistributionConfig(d.Get("zone_distribution_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("zone_distribution_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(zoneDistributionConfigProp)) && (ok || !reflect.DeepEqual(v, zoneDistributionConfigProp)) {
		obj["zoneDistributionConfig"] = zoneDistributionConfigProp
	}
	allowFewerZonesDeploymentProp, err := expandRedisClusterAllowFewerZonesDeployment(d.Get("allow_fewer_zones_deployment"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("allow_fewer_zones_deployment"); !tpgresource.IsEmptyValue(reflect.ValueOf(allowFewerZonesDeploymentProp)) && (ok || !reflect.DeepEqual(v, allowFewerZonesDeploymentProp)) {
		obj["allowFewerZonesDeployment"] = allowFewerZonesDeploymentProp
	}
	pscConfigsProp, err := expandRedisClusterPscConfigs(d.Get("psc_configs"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("psc_configs"); !tpgresource.IsEmptyValue(reflect.ValueOf(pscConfigsProp)) && (ok || !reflect.DeepEqual(v, pscConfigsProp)) {
		obj["pscConfigs"] = pscConfigsProp
	}
	replicaCountProp, err := expandRedisClusterReplicaCount(d.Get("replica_count"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("replica_count"); ok || !reflect.DeepEqual(v, replicaCountProp) {
		obj["replicaCount"] = replicaCountProp
	}
	shardCountProp, err := expandRedisClusterShardCount(d.Get("shard_count"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("shard_count"); !tpgresource.IsEmptyValue(reflect.ValueOf(shardCountProp)) && (ok || !reflect.DeepEqual(v, shardCountProp)) {
		obj["shardCount"] = shardCountProp
	}
	deletionProtectionEnabledProp, err := expandRedisClusterDeletionProtectionEnabled(d.Get("deletion_protection_enabled"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("deletion_protection_enabled"); !tpgresource.IsEmptyValue(reflect.ValueOf(deletionProtectionEnabledProp)) && (ok || !reflect.DeepEqual(v, deletionProtectionEnabledProp)) {
		obj["deletionProtectionEnabled"] = deletionProtectionEnabledProp
	}
	redisConfigsProp, err := expandRedisClusterRedisConfigs(d.Get("redis_configs"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("redis_configs"); !tpgresource.IsEmptyValue(reflect.ValueOf(redisConfigsProp)) && (ok || !reflect.DeepEqual(v, redisConfigsProp)) {
		obj["redisConfigs"] = redisConfigsProp
	}
	persistenceConfigProp, err := expandRedisClusterPersistenceConfig(d.Get("persistence_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("persistence_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(persistenceConfigProp)) && (ok || !reflect.DeepEqual(v, persistenceConfigProp)) {
		obj["persistenceConfig"] = persistenceConfigProp
	}
	maintenancePolicyProp, err := expandRedisClusterMaintenancePolicy(d.Get("maintenance_policy"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("maintenance_policy"); !tpgresource.IsEmptyValue(reflect.ValueOf(maintenancePolicyProp)) && (ok || !reflect.DeepEqual(v, maintenancePolicyProp)) {
		obj["maintenancePolicy"] = maintenancePolicyProp
	}
	crossClusterReplicationConfigProp, err := expandRedisClusterCrossClusterReplicationConfig(d.Get("cross_cluster_replication_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("cross_cluster_replication_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(crossClusterReplicationConfigProp)) && (ok || !reflect.DeepEqual(v, crossClusterReplicationConfigProp)) {
		obj["crossClusterReplicationConfig"] = crossClusterReplicationConfigProp
	}
	kmsKeyProp, err := expandRedisClusterKmsKey(d.Get("kms_key"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("kms_key"); !tpgresource.IsEmptyValue(reflect.ValueOf(kmsKeyProp)) && (ok || !reflect.DeepEqual(v, kmsKeyProp)) {
		obj["kmsKey"] = kmsKeyProp
	}

	return resourceRedisClusterEncoder(d, config, obj)
}

func resourceRedisClusterEncoder(d tpgresource.TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {

	// if the automated_backup_config is not defined, automatedBackupMode needs to be passed and set to DISABLED in the expand
	if obj["automatedBackupConfig"] == nil {
		config := meta.(*transport_tpg.Config)
		automatedBackupConfigProp, _ := expandRedisClusterAutomatedBackupConfig(d.Get("automated_backup_config"), d, config)
		obj["automatedBackupConfig"] = automatedBackupConfigProp
	}

	return obj, nil
}

func expandRedisClusterGcsSource(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedUris, err := expandRedisClusterGcsSourceUris(original["uris"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUris); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["uris"] = transformedUris
	}

	return transformed, nil
}

func expandRedisClusterGcsSourceUris(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	return v, nil
}

func expandRedisClusterManagedBackupSource(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedBackup, err := expandRedisClusterManagedBackupSourceBackup(original["backup"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBackup); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["backup"] = transformedBackup
	}

	return transformed, nil
}

func expandRedisClusterManagedBackupSourceBackup(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandRedisClusterAutomatedBackupConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	// The automated_backup_config block is not specified, so automatedBackupMode should be DISABLED
	transformed := make(map[string]interface{})
	if len(d.Get("automated_backup_config").([]interface{})) < 1 {
		transformed["automatedBackupMode"] = "DISABLED"
		return transformed, nil
	}

	// The automated_backup_config block is specified, so automatedBackupMode should be ENALBED
	transformed["automatedBackupMode"] = "ENABLED"
	transformedFixedFrequencySchedule, err := expandRedisClusterAutomatedBackupConfigFixedFrequencySchedule(original["fixed_frequency_schedule"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFixedFrequencySchedule); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["fixedFrequencySchedule"] = transformedFixedFrequencySchedule
	}

	transformedRetention, err := expandRedisClusterAutomatedBackupConfigRetention(original["retention"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRetention); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["retention"] = transformedRetention
	}

	return transformed, nil
}

func expandRedisClusterAutomatedBackupConfigFixedFrequencySchedule(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedStartTime, err := expandRedisClusterAutomatedBackupConfigFixedFrequencyScheduleStartTime(original["start_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStartTime); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["startTime"] = transformedStartTime
	}

	return transformed, nil
}

func expandRedisClusterAutomatedBackupConfigFixedFrequencyScheduleStartTime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedHours, err := expandRedisClusterAutomatedBackupConfigFixedFrequencyScheduleStartTimeHours(original["hours"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHours); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["hours"] = transformedHours
	}

	return transformed, nil
}

func expandRedisClusterAutomatedBackupConfigFixedFrequencyScheduleStartTimeHours(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandRedisClusterAutomatedBackupConfigRetention(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandRedisClusterAuthorizationMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandRedisClusterTransitEncryptionMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandRedisClusterNodeType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandRedisClusterZoneDistributionConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMode, err := expandRedisClusterZoneDistributionConfigMode(original["mode"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMode); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["mode"] = transformedMode
	}

	transformedZone, err := expandRedisClusterZoneDistributionConfigZone(original["zone"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedZone); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["zone"] = transformedZone
	}

	return transformed, nil
}

func expandRedisClusterZoneDistributionConfigMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandRedisClusterZoneDistributionConfigZone(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandRedisClusterAllowFewerZonesDeployment(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandRedisClusterPscConfigs(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedNetwork, err := expandRedisClusterPscConfigsNetwork(original["network"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedNetwork); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["network"] = transformedNetwork
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandRedisClusterPscConfigsNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandRedisClusterReplicaCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandRedisClusterShardCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandRedisClusterDeletionProtectionEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandRedisClusterRedisConfigs(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandRedisClusterPersistenceConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMode, err := expandRedisClusterPersistenceConfigMode(original["mode"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMode); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["mode"] = transformedMode
	}

	transformedRdbConfig, err := expandRedisClusterPersistenceConfigRdbConfig(original["rdb_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRdbConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["rdbConfig"] = transformedRdbConfig
	}

	transformedAofConfig, err := expandRedisClusterPersistenceConfigAofConfig(original["aof_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAofConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["aofConfig"] = transformedAofConfig
	}

	return transformed, nil
}

func expandRedisClusterPersistenceConfigMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandRedisClusterPersistenceConfigRdbConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRdbSnapshotPeriod, err := expandRedisClusterPersistenceConfigRdbConfigRdbSnapshotPeriod(original["rdb_snapshot_period"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRdbSnapshotPeriod); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["rdbSnapshotPeriod"] = transformedRdbSnapshotPeriod
	}

	transformedRdbSnapshotStartTime, err := expandRedisClusterPersistenceConfigRdbConfigRdbSnapshotStartTime(original["rdb_snapshot_start_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRdbSnapshotStartTime); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["rdbSnapshotStartTime"] = transformedRdbSnapshotStartTime
	}

	return transformed, nil
}

func expandRedisClusterPersistenceConfigRdbConfigRdbSnapshotPeriod(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandRedisClusterPersistenceConfigRdbConfigRdbSnapshotStartTime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandRedisClusterPersistenceConfigAofConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAppendFsync, err := expandRedisClusterPersistenceConfigAofConfigAppendFsync(original["append_fsync"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAppendFsync); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["appendFsync"] = transformedAppendFsync
	}

	return transformed, nil
}

func expandRedisClusterPersistenceConfigAofConfigAppendFsync(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandRedisClusterMaintenancePolicy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCreateTime, err := expandRedisClusterMaintenancePolicyCreateTime(original["create_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCreateTime); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["createTime"] = transformedCreateTime
	}

	transformedUpdateTime, err := expandRedisClusterMaintenancePolicyUpdateTime(original["update_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUpdateTime); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["updateTime"] = transformedUpdateTime
	}

	transformedWeeklyMaintenanceWindow, err := expandRedisClusterMaintenancePolicyWeeklyMaintenanceWindow(original["weekly_maintenance_window"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedWeeklyMaintenanceWindow); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["weeklyMaintenanceWindow"] = transformedWeeklyMaintenanceWindow
	}

	return transformed, nil
}

func expandRedisClusterMaintenancePolicyCreateTime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandRedisClusterMaintenancePolicyUpdateTime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandRedisClusterMaintenancePolicyWeeklyMaintenanceWindow(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedDay, err := expandRedisClusterMaintenancePolicyWeeklyMaintenanceWindowDay(original["day"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDay); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["day"] = transformedDay
		}

		transformedDuration, err := expandRedisClusterMaintenancePolicyWeeklyMaintenanceWindowDuration(original["duration"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDuration); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["duration"] = transformedDuration
		}

		transformedStartTime, err := expandRedisClusterMaintenancePolicyWeeklyMaintenanceWindowStartTime(original["start_time"], d, config)
		if err != nil {
			return nil, err
		} else {
			transformed["startTime"] = transformedStartTime
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandRedisClusterMaintenancePolicyWeeklyMaintenanceWindowDay(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandRedisClusterMaintenancePolicyWeeklyMaintenanceWindowDuration(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandRedisClusterMaintenancePolicyWeeklyMaintenanceWindowStartTime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 {
		return nil, nil
	}

	if l[0] == nil {
		transformed := make(map[string]interface{})
		return transformed, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedHours, err := expandRedisClusterMaintenancePolicyWeeklyMaintenanceWindowStartTimeHours(original["hours"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHours); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["hours"] = transformedHours
	}

	transformedMinutes, err := expandRedisClusterMaintenancePolicyWeeklyMaintenanceWindowStartTimeMinutes(original["minutes"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMinutes); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["minutes"] = transformedMinutes
	}

	transformedSeconds, err := expandRedisClusterMaintenancePolicyWeeklyMaintenanceWindowStartTimeSeconds(original["seconds"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSeconds); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["seconds"] = transformedSeconds
	}

	transformedNanos, err := expandRedisClusterMaintenancePolicyWeeklyMaintenanceWindowStartTimeNanos(original["nanos"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNanos); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["nanos"] = transformedNanos
	}

	return transformed, nil
}

func expandRedisClusterMaintenancePolicyWeeklyMaintenanceWindowStartTimeHours(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandRedisClusterMaintenancePolicyWeeklyMaintenanceWindowStartTimeMinutes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandRedisClusterMaintenancePolicyWeeklyMaintenanceWindowStartTimeSeconds(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandRedisClusterMaintenancePolicyWeeklyMaintenanceWindowStartTimeNanos(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandRedisClusterCrossClusterReplicationConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedClusterRole, err := expandRedisClusterCrossClusterReplicationConfigClusterRole(original["cluster_role"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedClusterRole); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["clusterRole"] = transformedClusterRole
	}

	transformedPrimaryCluster, err := expandRedisClusterCrossClusterReplicationConfigPrimaryCluster(original["primary_cluster"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPrimaryCluster); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["primaryCluster"] = transformedPrimaryCluster
	}

	transformedSecondaryClusters, err := expandRedisClusterCrossClusterReplicationConfigSecondaryClusters(original["secondary_clusters"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSecondaryClusters); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["secondaryClusters"] = transformedSecondaryClusters
	}

	transformedMembership, err := expandRedisClusterCrossClusterReplicationConfigMembership(original["membership"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMembership); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["membership"] = transformedMembership
	}

	transformedUpdateTime, err := expandRedisClusterCrossClusterReplicationConfigUpdateTime(original["update_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUpdateTime); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["updateTime"] = transformedUpdateTime
	}

	return transformed, nil
}

func expandRedisClusterCrossClusterReplicationConfigClusterRole(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandRedisClusterCrossClusterReplicationConfigPrimaryCluster(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCluster, err := expandRedisClusterCrossClusterReplicationConfigPrimaryClusterCluster(original["cluster"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCluster); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["cluster"] = transformedCluster
	}

	transformedUid, err := expandRedisClusterCrossClusterReplicationConfigPrimaryClusterUid(original["uid"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUid); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["uid"] = transformedUid
	}

	return transformed, nil
}

func expandRedisClusterCrossClusterReplicationConfigPrimaryClusterCluster(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandRedisClusterCrossClusterReplicationConfigPrimaryClusterUid(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandRedisClusterCrossClusterReplicationConfigSecondaryClusters(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedCluster, err := expandRedisClusterCrossClusterReplicationConfigSecondaryClustersCluster(original["cluster"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedCluster); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["cluster"] = transformedCluster
		}

		transformedUid, err := expandRedisClusterCrossClusterReplicationConfigSecondaryClustersUid(original["uid"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedUid); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["uid"] = transformedUid
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandRedisClusterCrossClusterReplicationConfigSecondaryClustersCluster(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandRedisClusterCrossClusterReplicationConfigSecondaryClustersUid(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandRedisClusterCrossClusterReplicationConfigMembership(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPrimaryCluster, err := expandRedisClusterCrossClusterReplicationConfigMembershipPrimaryCluster(original["primary_cluster"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPrimaryCluster); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["primaryCluster"] = transformedPrimaryCluster
	}

	transformedSecondaryClusters, err := expandRedisClusterCrossClusterReplicationConfigMembershipSecondaryClusters(original["secondary_clusters"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSecondaryClusters); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["secondaryClusters"] = transformedSecondaryClusters
	}

	return transformed, nil
}

func expandRedisClusterCrossClusterReplicationConfigMembershipPrimaryCluster(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCluster, err := expandRedisClusterCrossClusterReplicationConfigMembershipPrimaryClusterCluster(original["cluster"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCluster); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["cluster"] = transformedCluster
	}

	transformedUid, err := expandRedisClusterCrossClusterReplicationConfigMembershipPrimaryClusterUid(original["uid"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUid); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["uid"] = transformedUid
	}

	return transformed, nil
}

func expandRedisClusterCrossClusterReplicationConfigMembershipPrimaryClusterCluster(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandRedisClusterCrossClusterReplicationConfigMembershipPrimaryClusterUid(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandRedisClusterCrossClusterReplicationConfigMembershipSecondaryClusters(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedCluster, err := expandRedisClusterCrossClusterReplicationConfigMembershipSecondaryClustersCluster(original["cluster"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedCluster); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["cluster"] = transformedCluster
		}

		transformedUid, err := expandRedisClusterCrossClusterReplicationConfigMembershipSecondaryClustersUid(original["uid"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedUid); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["uid"] = transformedUid
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandRedisClusterCrossClusterReplicationConfigMembershipSecondaryClustersCluster(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandRedisClusterCrossClusterReplicationConfigMembershipSecondaryClustersUid(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandRedisClusterCrossClusterReplicationConfigUpdateTime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandRedisClusterKmsKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
