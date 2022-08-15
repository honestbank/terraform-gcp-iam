lint: docs
	terraform fmt --recursive

validate: lint
	terraform init --upgrade
	terraform validate

docs:
	rm -rf modules/*/.terraform modules/*/.terraform.lock.hcl
	rm -rf examples/*/.terraform examples/*/.terraform.lock.hcl
	terraform-docs -c .terraform-docs.yml .
	terraform-docs -c .terraform-docs-examples.yml .

commit: docs lint validate
