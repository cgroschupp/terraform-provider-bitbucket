package bitbucket

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	bb "github.com/ktrysmt/go-bitbucket"
)

func resourcePipelineVariable() *schema.Resource {
	return &schema.Resource{
		Description:   "bitbucket_pipeline_variable` manages a bitbucket pipeline variable.",
		CreateContext: resourcePipelineVariableCreate,
		ReadContext:   resourcePipelineVariableRead,
		UpdateContext: resourcePipelineVariableUpdate,
		DeleteContext: resourcePipelineVariableDelete,
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The ID of the pipeline variable.",
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
			"key": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The unique name of the variable.",
			},
			"value": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The value of the variable. If the variable is secured, this will be empty.",
			},
			"secured": {
				Type:        schema.TypeBool,
				Default:     false,
				Optional:    true,
				Description: "If true, this variable will be treated as secured. The value will never be exposed in the logs or the REST API.",
			},
		},
	}
}

func resourcePipelineVariableCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*bb.Client)

	var diags diag.Diagnostics
	workspace := d.Get("workspace").(string)
	repo := d.Get("repository").(string)
	key := d.Get("key").(string)
	value := d.Get("value").(string)
	secured := d.Get("secured").(bool)

	pv, err := c.Repositories.Repository.AddPipelineVariable(&bb.RepositoryPipelineVariableOptions{RepoSlug: repo, Owner: workspace, Key: key, Value: value, Secured: secured})

	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(pv.Uuid)

	return diags
}

func resourcePipelineVariableRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*bb.Client)

	var diags diag.Diagnostics
	workspace := d.Get("workspace").(string)
	repo := d.Get("repository").(string)
	pv, err := c.Repositories.Repository.GetPipelineVariable(&bb.RepositoryPipelineVariableOptions{Owner: workspace, RepoSlug: repo, Uuid: d.Id()})
	if err != nil {
		return diag.FromErr(fmt.Errorf("unable to read pipeline variabe: %s", err))
	}
	err = d.Set("key", pv.Key)
	if err != nil {
		return diag.FromErr(err)
	}
	if !pv.Secured {
		err = d.Set("value", pv.Value)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	err = d.Set("secured", pv.Secured)
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourcePipelineVariableUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*bb.Client)

	var diags diag.Diagnostics
	workspace := d.Get("workspace").(string)
	repo := d.Get("repository").(string)
	key := d.Get("key").(string)
	value := d.Get("value").(string)
	secured := d.Get("secured").(bool)

	data := &bb.RepositoryPipelineVariableOptions{Owner: workspace, RepoSlug: repo, Uuid: d.Id(), Key: key, Value: value, Secured: secured}
	_, err := c.Repositories.Repository.UpdatePipelineVariable(data)

	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourcePipelineVariableDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*bb.Client)

	var diags diag.Diagnostics
	workspace := d.Get("workspace").(string)
	repo := d.Get("repository").(string)

	_, err := c.Repositories.Repository.DeletePipelineVariable(&bb.RepositoryPipelineVariableDeleteOptions{Owner: workspace, RepoSlug: repo, Uuid: d.Id()})
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")

	return diags
}
