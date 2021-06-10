package bitbucket

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	bb "github.com/ktrysmt/go-bitbucket"
)

func resourceGroupPermission() *schema.Resource {
	return &schema.Resource{
		Description:   "`bitbucket_group_permission` manages group permission for a repository.",
		CreateContext: resourceGroupPermissionCreate,
		ReadContext:   resourceGroupPermissionRead,
		DeleteContext: resourceGroupPermissionDelete,
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The ID of the group permission.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "This can either be the repository slug or the UUID of the repository.",
			},
			"workspace": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "This can either be the workspace ID (slug) or the workspace UUID.",
			},
			"group": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "This can either be the group or the group UUID.",
			},
			"permission": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"read", "write", "admin"}, false),
				Description:  "Permission of the group.",
			},
		},
	}
}

func resourceGroupPermissionCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*bb.Client)

	var diags diag.Diagnostics
	workspace := d.Get("workspace").(string)
	repo := d.Get("repository").(string)
	group := d.Get("group").(string)
	permission := d.Get("permission").(string)

	_, err := c.GroupPrivileges.Add(bb.GroupPrivilegesOptions{RepoSlug: repo, Owner: workspace, Group: group, Permission: permission})

	if err != nil {
		return diag.FromErr(fmt.Errorf("unable to add group permission: %s", err))
	}
	d.SetId(fmt.Sprintf("%s-%s-%s", workspace, repo, group))

	return diags
}

func resourceGroupPermissionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*bb.Client)

	var diags diag.Diagnostics
	workspace := d.Get("workspace").(string)
	repo := d.Get("repository").(string)
	group := d.Get("group").(string)

	data, err := c.GroupPrivileges.Get(bb.GroupPrivilegesOptions{RepoSlug: repo, Owner: workspace, Group: group})
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("permission", data[0].Privilege)
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceGroupPermissionDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*bb.Client)

	var diags diag.Diagnostics
	workspace := d.Get("workspace").(string)
	repo := d.Get("repository").(string)
	group := d.Get("group").(string)

	_, err := c.GroupPrivileges.Delete(bb.GroupPrivilegesOptions{RepoSlug: repo, Owner: workspace, Group: group})
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")

	return diags
}
