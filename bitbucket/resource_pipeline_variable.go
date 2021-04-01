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
		CreateContext: resourcePipelineVariableCreate,
		ReadContext:   resourcePipelineVariableRead,
		UpdateContext: resourcePipelineVariableUpdate,
		DeleteContext: resourcePipelineVariableDelete,
		Schema: map[string]*schema.Schema{
			"repo_uuid": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"workspace": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"value": {
				Type:     schema.TypeString,
				Required: true,
			},
			"secured": {
				Type:     schema.TypeBool,
				Default:  false,
				Optional: true,
			},
		},
	}
}

func resourcePipelineVariableCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*bb.Client)

	var diags diag.Diagnostics
	workspace := d.Get("workspace").(string)
	repoUuid := d.Get("repo_uuid").(string)
	key := d.Get("key").(string)
	value := d.Get("value").(string)
	secured := d.Get("secured").(bool)

	pv, err := c.Repositories.Repository.AddPipelineVariable(&bb.RepositoryPipelineVariableOptions{RepoSlug: repoUuid, Owner: workspace, Key: key, Value: value, Secured: secured})

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
	repoUuid := d.Get("repo_uuid").(string)
	pv, err := c.Repositories.Repository.GetPipelineVariable(&bb.RepositoryPipelineVariableOptions{Owner: workspace, RepoSlug: repoUuid, Uuid: d.Id()})
	if err != nil {
		return diag.FromErr(fmt.Errorf("unable to read pipeline variabe: %s", err))
	}
	d.Set("key", pv.Key)
	if !pv.Secured {
		d.Set("value", pv.Value)
	}
	d.Set("secured", pv.Secured)
	return diags
}

func resourcePipelineVariableUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*bb.Client)

	var diags diag.Diagnostics
	workspace := d.Get("workspace").(string)
	repoUuid := d.Get("repo_uuid").(string)
	key := d.Get("key").(string)
	value := d.Get("value").(string)
	secured := d.Get("secured").(bool)

	data := &bb.RepositoryPipelineVariableOptions{Owner: workspace, RepoSlug: repoUuid, Uuid: d.Id(), Key: key, Value: value, Secured: secured}
	c.Repositories.Repository.UpdatePipelineVariable(data)
	return diags
}

func resourcePipelineVariableDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*bb.Client)

	var diags diag.Diagnostics
	workspace := d.Get("workspace").(string)
	repoUuid := d.Get("repo_uuid").(string)

	_, err := c.Repositories.Repository.DeletePipelineVariable(&bb.RepositoryPipelineVariableDeleteOptions{Owner: workspace, RepoSlug: repoUuid, Uuid: d.Id()})
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")

	return diags
}
