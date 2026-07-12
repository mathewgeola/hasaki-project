APP_NAME := hp
CMD_DIR  := . 
OUT_DIR  := ./bin

VERSION  := $(shell cat VERSION 2>/dev/null || echo "dev")

LDFLAGS  := -s -w -X 'hasaki-project/cmd.Version=$(VERSION)'

.PHONY: default
default: build

.PHONY: tidy
tidy:
	go mod tidy
	go fmt ./...

.PHONY: clean
clean:
	rm -rf $(OUT_DIR)

.PHONY: build
build: tidy clean
	go build -ldflags="$(LDFLAGS)" -o $(OUT_DIR)/$(APP_NAME) $(CMD_DIR)

.PHONY: build-all
build-all: tidy clean
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="$(LDFLAGS)" -o $(OUT_DIR)/$(APP_NAME)-linux-amd64 $(CMD_DIR)
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="$(LDFLAGS)" -o $(OUT_DIR)/$(APP_NAME)-windows-amd64.exe $(CMD_DIR)
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags="$(LDFLAGS)" -o $(OUT_DIR)/$(APP_NAME)-darwin-arm64 $(CMD_DIR)
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags="$(LDFLAGS)" -o $(OUT_DIR)/$(APP_NAME)-darwin-amd64 $(CMD_DIR)
	@ls -lh $(OUT_DIR)

.PHONY: install
install: build
	sudo mv $(OUT_DIR)/$(APP_NAME) /usr/local/bin/$(APP_NAME)

.PHONY: uninstall
uninstall:
	sudo rm -f /usr/local/bin/$(APP_NAME)