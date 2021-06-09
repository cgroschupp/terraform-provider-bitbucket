package bitbucket

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	bb "github.com/ktrysmt/go-bitbucket"
)

func resourceDefaultReviewer() *schema.Resource {
	return &schema.Resource{
		Description:   "`bitbucket_default_reviewers` manages default reviewers for a repository.",
		CreateContext: resourceDefaultReviewerCreate,
		ReadContext:   resourceDefaultReviewerRead,
		DeleteContext: resourceDefaultReviewerDelete,
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The ID of the default reviewer.",
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
			"user": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "This can either be the username or the user UUID.",
			},
		},
	}
}

func resourceDefaultReviewerCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*bb.Client)

	var diags diag.Diagnostics
	workspace := d.Get("workspace").(string)
	repo := d.Get("repository").(string)
	user := d.Get("user").(string)

	defaultReviewer, err := c.Repositories.Repository.AddDefaultReviewer(&bb.RepositoryDefaultReviewerOptions{RepoSlug: repo, Owner: workspace, Username: user})

	if err != nil {
		return diag.FromErr(fmt.Errorf("unable to add default reviewer: %s", err))
	}
	d.SetId(defaultReviewer.Uuid)

	return diags
}

func resourceDefaultReviewerRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*bb.Client)

	var diags diag.Diagnostics
	workspace := d.Get("workspace").(string)
	repo := d.Get("repository").(string)

	_, err := c.Repositories.Repository.GetDefaultReviewer(&bb.RepositoryDefaultReviewerOptions{Owner: workspace, RepoSlug: repo, Username: d.Id()})
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceDefaultReviewerDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*bb.Client)

	var diags diag.Diagnostics
	workspace := d.Get("workspace").(string)
	repo := d.Get("repository").(string)

	_, err := c.Repositories.Repository.DeleteDefaultReviewer(&bb.RepositoryDefaultReviewerOptions{Owner: workspace, Username: d.Id(), RepoSlug: repo})
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")

	return diags
}
