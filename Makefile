BIN=bin/terraform-provider-julien
TEST?=$$(go list ./... | grep -v 'vendor')

OS_ARCH=$(shell go env GOHOSTOS)_$(shell go env GOHOSTARCH)

build:
	go build -o $(BIN)

install: build
	mkdir -p ~/.terraform.d/plugins/hashicorp.com/edu/julien/0.2/$(OS_ARCH)
	mv $(BIN) ~/.terraform.d/plugins/hashicorp.com/edu/julien/0.2/$(OS_ARCH)

test:
	go test -i $(TEST) || exit 1
	echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

provider_test: test install
	cd examples && $(MAKE) test
