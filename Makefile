.PHONY: all generate tidy fmt clean

SPEC = openapi\openapi.yaml

GEN_DIR = internal\api

GEN_PKG = api

all: generate tidy fmt

generate:
	@echo ------------------------------------------------------------
	@echo ">>> Installing ogen@latest..."
	@go install github.com/ogen-go/ogen/cmd/ogen@latest
	@echo ">>> Removing \"$(GEN_DIR)\" if it exists..."
	@cmd /C "if exist \"$(GEN_DIR)\" (rmdir /s /q \"$(GEN_DIR)\") else (echo INFO: \"$(GEN_DIR)\" not found, skipping removal)"
	@echo ">>> Generating code from \"$(SPEC)\"..."
	@ogen --target internal/api --package $(GEN_PKG) --clean $(SPEC)

tidy:
	@echo ">>> Running go mod tidy..."
	@go mod tidy


fmt:
	@echo ">>> Running go fmt ./..."
	@go fmt ./...

clean:
	@echo ">>> Cleaning up \"$(GEN_DIR)\"..."
	@cmd /C "if exist \"$(GEN_DIR)\" (rmdir /s /q \"$(GEN_DIR)\") else (echo INFO: \"$(GEN_DIR)\" not found, nothing to delete)"
