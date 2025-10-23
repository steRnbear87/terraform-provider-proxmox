package rootmount

import (
	pveSDK "github.com/Telmate/proxmox-api-go/proxmox"
	"github.com/Telmate/terraform-provider-proxmox/v2/proxmox/Internal/helper/size"
	"github.com/Telmate/terraform-provider-proxmox/v2/proxmox/Internal/resource/guest/lxc/_sub/acl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Terraform(rootMount *pveSDK.LxcBootMount, d *schema.ResourceData) {
	if rootMount == nil {
		d.Set(Root, nil)
		return
	}
	mounts := make([]map[string]any, 1)
	mounts[0] = map[string]any{
		schemaACL:       acl.Terraform(rootMount.ACL),
		schemaOptions:   TerraformOptions(rootMount.Options),
		schemaReplicate: *rootMount.Replicate,
		schemaSize:      size.String(int64(*rootMount.SizeInKibibytes)),
		schemaStorage:   *rootMount.Storage}
	if rootMount.Quota != nil {
		mounts[0][schemaQuota] = *rootMount.Quota
	}
	d.Set(Root, mounts)
}

func TerraformOptions(options *pveSDK.LxcBootMountOptions) []map[string]any {
	if options == nil {
		return nil
	}
	return []map[string]any{{
		schemaDiscard:  *options.Discard,
		schemaLazyTime: *options.LazyTime,
		schemaNoATime:  *options.NoATime,
		schemaNoSuid:   *options.NoSuid}}
}
