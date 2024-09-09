default: e2e-apply

FEATURE  				?= examples/provider
TERRAFORM_FLAGS			?= -auto-approve

e2e-init:
	rm -f $(FEATURE)/.terraform.lock.hcl
	terraform -chdir=$(FEATURE) init

e2e-apply: e2e-init
	terraform -chdir=$(FEATURE) plan
	terraform -chdir=$(FEATURE) apply $(TERRAFORM_FLAGS)

e2e-destroy:
	terraform -chdir=$(FEATURE) destroy $(TERRAFORM_FLAGS)
