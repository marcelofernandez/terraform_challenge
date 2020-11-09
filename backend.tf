terraform {
  backend "remote" {
    organization = "marcelofernandez"

    workspaces {
      name = "terraform_challenge"
    }
  }
}
