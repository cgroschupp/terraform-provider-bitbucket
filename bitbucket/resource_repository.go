package bitbucket

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	bb "github.com/ktrysmt/go-bitbucket"
)

func resourceRepository() *schema.Resource {
	return &schema.Resource{
		Description: "`bitbucket_repository` manages a bitbucket repository.",

		CreateContext: resourceRepositoryCreate,
		ReadContext:   resourceRepositoryRead,
		UpdateContext: resourceRepositoryUpdate,
		DeleteContext: resourceRepositoryDelete,
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The ID of the repository.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the repository.",
			},
			"workspace": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the workspace where the repository is created.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the repository.",
			},
		},
	}
}

func resourceRepositoryCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*bb.Client)

	var diags diag.Diagnostics
	workspace := d.Get("workspace").(string)
	name := d.Get("name").(string)
	description := d.Get("description").(string)

	repo, err := c.Repositories.Repository.Create(&bb.RepositoryOptions{Owner: workspace, RepoSlug: name, Description: description})

	if err != nil {
		return diag.FromErr(fmt.Errorf("unable to create repo: %s", err))
	}
	d.SetId(repo.Uuid)

	return diags
}

func resourceRepositoryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*bb.Client)

	var diags diag.Diagnostics
	workspace := d.Get("workspace").(string)
	pv, err := c.Repositories.Repository.Get(&bb.RepositoryOptions{Owner: workspace, RepoSlug: d.Id()})
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("name", pv.Name)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("description", pv.Description)
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceRepositoryUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*bb.Client)

	var diags diag.Diagnostics
	workspace := d.Get("workspace").(string)
	name := d.Get("name").(string)
	description := d.Get("description").(string)

	data := &bb.RepositoryOptions{Owner: workspace, Uuid: d.Id(), RepoSlug: name, Description: description}
	repo, err := c.Repositories.Repository.Update(data)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("description", repo.Description)
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceRepositoryDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*bb.Client)

	var diags diag.Diagnostics
	workspace := d.Get("workspace").(string)

	_, err := c.Repositories.Repository.Delete(&bb.RepositoryOptions{Owner: workspace, Uuid: d.Id()})
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")

	return diags
}
