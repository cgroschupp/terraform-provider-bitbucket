package bitbucket

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	bb "github.com/ktrysmt/go-bitbucket"
)

func dataSourceRepository() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceRepositoryRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"repo_slug": {
				Type:     schema.TypeString,
				Required: true,
			},
			"workspace": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceRepositoryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*bb.Client)

	var diags diag.Diagnostics

	workspace := d.Get("workspace").(string)
	slug := d.Get("repo_slug").(string)
	repo, err := c.Repositories.Repository.Get(&bb.RepositoryOptions{Owner: workspace, RepoSlug: slug})

	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("name", repo.Name)
	d.SetId(repo.Uuid)

	return diags
}
