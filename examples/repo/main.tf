terraform {
  required_providers {
    bitbucket = {
      version = "0.2"
      source  = "github.com/cgroschupp/bitbucket"
    }
  }
}

provider "bitbucket" {}

resource "bitbucket_repository" "repo" {
  workspace   = "cgroschupp"
  repo_slug   = "test"
  description = "test"
}

output "repo_name" {
  value = bitbucket_repository.repo.repo_slug
}

output "repo_id" {
  value = bitbucket_repository.repo.id
}

resource "bitbucket_pipeline_variable" "var" {
  workspace = bitbucket_repository.repo.workspace
  repo_uuid = bitbucket_repository.repo.id
  key       = "var"
  value     = "var_value"
  secured   = false
}
