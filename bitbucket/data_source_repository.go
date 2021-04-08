package bitbucket

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	bb "github.com/ktrysmt/go-bitbucket"
)

func dataSourceRepository() *schema.Resource {
	return &schema.Resource{
		Description: "`bitbucket_repository` data source can be used to retrieve the UUID for a reposistory by name and workspace.",
		ReadContext: dataSourceRepositoryRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of the repository.",
			},
			"workspace": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The Workspace of the repository.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the repository.",
			},
		},
	}
}

func dataSourceRepositoryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*bb.Client)

	var diags diag.Diagnostics

	workspace := d.Get("workspace").(string)
	name := d.Get("name").(string)
	repo, err := c.Repositories.Repository.Get(&bb.RepositoryOptions{Owner: workspace, RepoSlug: name})

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(repo.Uuid)

	return diags
}
