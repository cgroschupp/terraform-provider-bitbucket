terraform {
  required_providers {
    bitbucket = {
      version = "0.2"
      source  = "github.com/cgroschupp/bitbucket"
    }
  }
}

provider "bitbucket" {

}
data "bitbucket_workspace" "workspace" {
  name = "cgroschupp"
}

resource "bitbucket_repository" "repo" {
  workspace   = data.bitbucket_workspace.workspace.id
  name        = "test"
  description = "test"
}

output "repo_name" {
  value = bitbucket_repository.repo.name
}

output "repo_id" {
  value = bitbucket_repository.repo.id
}

resource "bitbucket_pipeline_variable" "var" {
  workspace  = bitbucket_repository.repo.workspace
  repository = bitbucket_repository.repo.id
  key        = "var"
  value      = "var_value"
  secured    = false
}

resource "bitbucket_branch_restriction" "var" {
  workspace  = bitbucket_repository.repo.workspace
  repository = bitbucket_repository.repo.id
  kind       = "force"
  pattern    = "master1"
}
