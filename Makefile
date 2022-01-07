BIN=bin/terraform-provider-colors
TEST?=$$(go list ./... | grep -v 'vendor')

HOSTNAME=hashicorp.com
NAMESPACE=julien-jean
NAME=colors
VERSION=0.2

OS_ARCH=$(shell go env GOHOSTOS)_$(shell go env GOHOSTARCH)

INSTALL_DIR=~/.terraform.d/plugins/$(HOSTNAME)/$(NAMESPACE)/$(NAME)/$(VERSION)/$(OS_ARCH)

build:
	goreleaser build --snapshot --single-target --rm-dist

install: build generate
	mkdir -p $(INSTALL_DIR)
	cp $(shell ls dist/terraform-provider-colors_linux_amd64/*) $(INSTALL_DIR)
	chmod +x $(INSTALL_DIR)/*

test:
	go test -i $(TEST) || exit 1
	echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

provider_test: test install
	cd examples && $(MAKE) test

docs:
	go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs
