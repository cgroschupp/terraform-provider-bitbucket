package bitbucket

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	bb "github.com/ktrysmt/go-bitbucket"
)

func resourceBranchRestriction() *schema.Resource {
	return &schema.Resource{
		Description:   "`bitbucket_branch_restriction` manages a branch restriction rule for a repository.",
		CreateContext: resourceBranchRestrictionCreate,
		ReadContext:   resourceBranchRestrictionRead,
		UpdateContext: resourceBranchRestrictionUpdate,
		DeleteContext: resourceBranchRestrictionDelete,
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The ID of the branch restriction.",
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
			"pattern": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Branch restrictions applied to branches of this pattern.",
			},
			"kind": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"require_tasks_to_be_completed", "force", "restrict_merges", "enforce_merge_checks", "reset_pullrequest_changes_requested_on_change", "require_approvals_to_merge", "allow_auto_merge_when_builds_pass", "delete", "require_all_dependencies_merged", "require_no_changes_requested", "push", "require_passing_builds_to_merge", "reset_pullrequest_approvals_on_change", "require_default_reviewer_approvals_to_merge"}, false),
				ForceNew:     true,
				Description:  "Branch restrictions of this type.",
			},
			"value": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Value with kind-specific semantics: `require_approvals_to_merge` uses it to require a minimum number of approvals on a PR; `require_default_reviewer_approvals_to_merge` uses it to require a minimum number of approvals from default reviewers on a PR; `require_passing_builds_to_merge` uses it to require a minimum number of passing builds.",
			},
		},
	}
}

func resourceBranchRestrictionCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*bb.Client)

	var diags diag.Diagnostics
	workspace := d.Get("workspace").(string)
	repo := d.Get("repository").(string)
	kind := d.Get("kind").(string)
	pattern := d.Get("pattern").(string)
	var value *int
	if v, _ := d.Get("value").(int); v != 0 {
		value = &v
	}

	branchRestriction, err := c.Repositories.BranchRestrictions.Create(&bb.BranchRestrictionsOptions{RepoSlug: repo, Owner: workspace, Kind: kind, Pattern: pattern, Value: value})

	if err != nil {
		return diag.FromErr(fmt.Errorf("unable to create repo: %s", err))
	}
	d.SetId(strconv.FormatInt(int64(branchRestriction.ID), 10))

	return diags
}

func resourceBranchRestrictionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*bb.Client)

	var diags diag.Diagnostics
	workspace := d.Get("workspace").(string)
	repo := d.Get("repository").(string)
	br, err := c.Repositories.BranchRestrictions.Get(&bb.BranchRestrictionsOptions{Owner: workspace, RepoSlug: repo, ID: d.Id()})
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("pattern", br.Pattern)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("kind", br.Kind)
	if err != nil {
		return diag.FromErr(err)
	}

	if br.Value != nil {
		err = d.Set("value", br.Value)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return diags
}

func resourceBranchRestrictionUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*bb.Client)

	var diags diag.Diagnostics
	workspace := d.Get("workspace").(string)
	repo := d.Get("repository").(string)
	pattern := d.Get("pattern").(string)
	kind := d.Get("kind").(string)

	data := &bb.BranchRestrictionsOptions{Owner: workspace, ID: d.Id(), RepoSlug: repo, Pattern: pattern, Kind: kind}
	_, err := c.Repositories.BranchRestrictions.Update(data)
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceBranchRestrictionDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*bb.Client)

	var diags diag.Diagnostics
	workspace := d.Get("workspace").(string)
	repo := d.Get("repository").(string)

	_, err := c.Repositories.BranchRestrictions.Delete(&bb.BranchRestrictionsOptions{Owner: workspace, ID: d.Id(), RepoSlug: repo})
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")

	return diags
}
