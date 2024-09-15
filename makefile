# 出力先のディレクトリ
BIN_DIR:=bin

GO          = go
GO_BUILD    = $(GO) build -trimpath
GO_GET      = $(GO) get
GO_FORMAT   = $(GO) fmt
GO_LIST     = $(GO) list
GO_TEST     = $(GO) test -v
GO_VET      = $(GO) vet
GO_TOOL     = $(GO) tool
GO_LDFLAGS  = -ldflags="-s -w -buildid="

BINARY_NAME:=rest-testify

TARGETS :=
TARGETS += $(addprefix $(BIN_DIR)/, $(BINARY_NAME:%=%))
DEPFILES :=
DEPFILES +=$(shell find . -type f -name "*.go")

GO_PKGROOT  = ./...
GO_PACKAGES = $(shell $(GO_LIST) $(GO_PKGROOT) | grep -v vendor)

UNITTEST_WORK := work


### PHONY ターゲットのビルドルール
setup:
	$(GO_GET) golang.org/x/lint/golint
	$(GO_GET) golang.org/x/tools/cmd/godoc

build: static_analysis $(TARGETS)

clean:
	rm -rf $(BIN_DIR) $(UNITTEST_WORK)

static_analysis: vet
vet:
	$(GO_VET) $(GO_PACKAGES)

fmt:
	$(GO_FORMAT) $(GO_PKGROOT)

doc:
	godoc --http=:8080

### 各バイナリファイル生成ルール
$(BIN_DIR)/%: main.go $(DEPFILES)
	$(GO_BUILD) $(GO_LDFLAGS) -o $@ ./$<
