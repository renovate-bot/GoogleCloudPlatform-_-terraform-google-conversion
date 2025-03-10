// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/identityplatform/Tenant.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package identityplatform

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const IdentityPlatformTenantAssetType string = "identitytoolkit.googleapis.com/Tenant"

func ResourceConverterIdentityPlatformTenant() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: IdentityPlatformTenantAssetType,
		Convert:   GetIdentityPlatformTenantCaiObject,
	}
}

func GetIdentityPlatformTenantCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//identitytoolkit.googleapis.com/projects/{{project}}/tenants/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetIdentityPlatformTenantApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: IdentityPlatformTenantAssetType,
			Resource: &cai.AssetResource{
				Version:              "v2",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/identitytoolkit/v2/rest",
				DiscoveryName:        "Tenant",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetIdentityPlatformTenantApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandIdentityPlatformTenantDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	allowPasswordSignupProp, err := expandIdentityPlatformTenantAllowPasswordSignup(d.Get("allow_password_signup"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("allow_password_signup"); !tpgresource.IsEmptyValue(reflect.ValueOf(allowPasswordSignupProp)) && (ok || !reflect.DeepEqual(v, allowPasswordSignupProp)) {
		obj["allowPasswordSignup"] = allowPasswordSignupProp
	}
	enableEmailLinkSigninProp, err := expandIdentityPlatformTenantEnableEmailLinkSignin(d.Get("enable_email_link_signin"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("enable_email_link_signin"); !tpgresource.IsEmptyValue(reflect.ValueOf(enableEmailLinkSigninProp)) && (ok || !reflect.DeepEqual(v, enableEmailLinkSigninProp)) {
		obj["enableEmailLinkSignin"] = enableEmailLinkSigninProp
	}
	disableAuthProp, err := expandIdentityPlatformTenantDisableAuth(d.Get("disable_auth"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("disable_auth"); !tpgresource.IsEmptyValue(reflect.ValueOf(disableAuthProp)) && (ok || !reflect.DeepEqual(v, disableAuthProp)) {
		obj["disableAuth"] = disableAuthProp
	}

	return obj, nil
}

func expandIdentityPlatformTenantDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformTenantAllowPasswordSignup(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformTenantEnableEmailLinkSignin(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformTenantDisableAuth(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
