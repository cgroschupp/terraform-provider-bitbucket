package bitbucket

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	bb "github.com/ktrysmt/go-bitbucket"
)

func resourceProject() *schema.Resource {
	return &schema.Resource{
		Description: "`bitbucket_project` manages a bitbucket project.",

		CreateContext: resourceProjectCreate,
		ReadContext:   resourceProjectRead,
		DeleteContext: resourceProjectDelete,
		UpdateContext: resourceProjectUpdate,
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The ID of the project.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the project.",
			},
			"key": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The key of the project.",
			},
			"workspace": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the workspace where the project is created.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the project.",
			},
		},
	}
}

func resourceProjectCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*bb.Client)

	var diags diag.Diagnostics
	workspace := d.Get("workspace").(string)
	name := d.Get("name").(string)
	key := d.Get("key").(string)
	description := d.Get("description").(string)

	project, err := c.Workspaces.CreateProject(&bb.ProjectOptions{Owner: workspace, Name: name, Description: description, Key: key})

	if err != nil {
		return diag.FromErr(fmt.Errorf("unable to create repo: %s", err))
	}
	d.SetId(project.Uuid)

	return diags
}

func resourceProjectUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*bb.Client)

	var diags diag.Diagnostics
	workspace := d.Get("workspace").(string)
	name := d.Get("name").(string)
	key := d.Get("key").(string)
	description := d.Get("description").(string)

	data := &bb.ProjectOptions{Owner: workspace, Uuid: d.Id(), Key: key, Description: description, Name: name}
	repo, err := c.Workspaces.UpdateProject(data)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("description", repo.Description)
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceProjectRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*bb.Client)

	var diags diag.Diagnostics
	workspace := d.Get("workspace").(string)
	key := d.Get("key").(string)
	pv, err := c.Workspaces.GetProject(&bb.ProjectOptions{Owner: workspace, Key: key})
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("name", pv.Name)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("key", pv.Key)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("description", pv.Description)
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
func resourceProjectDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*bb.Client)

	var diags diag.Diagnostics
	workspace := d.Get("workspace").(string)
	key := d.Get("key").(string)
	_, err := c.Workspaces.DeleteProject(&bb.ProjectOptions{Owner: workspace, Key: key})
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")

	return diags
}
