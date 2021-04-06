package bitbucket

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	bb "github.com/ktrysmt/go-bitbucket"
)

func dataSourceWorkspace() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceWorkspaceRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceWorkspaceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*bb.Client)

	var diags diag.Diagnostics

	name := d.Get("name").(string)

	workspace, err := c.Workspaces.Get(name)

	if err != nil {
		return diag.FromErr(err)
	}

	err = d.Set("name", workspace.Name)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(workspace.UUID)

	return diags
}
