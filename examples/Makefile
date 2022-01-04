init:
	terraform init

plan:
	terraform plan

apply:
	terraform apply -auto-approve

clean:
	rm -rf .terraform/
	rm -rf .terraform.lock.hcl

test: clean init plan apply
