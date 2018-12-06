BUILD_TIME = $(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
BUILD_DISTROS = darwin linux
GIT_COMMIT = $(shell git rev-parse HEAD)
GIT_TREE_STATE = $(shell test -n "`git status --porcelain`" && echo "dirty" || echo "clean")
GIT_IMPORT = github.com/AirHelp/filler/version
GO_LDFLAGS = -X $(GIT_IMPORT).gitCommit=$(GIT_COMMIT) \
	-X $(GIT_IMPORT).gitTreeState=$(GIT_TREE_STATE) \
	-X $(GIT_IMPORT).buildDate=$(BUILD_TIME)

FILLER_VERSION?=$(shell awk -F\" '/^const version/ { print $$2; exit }' version/version.go)

default: test

docker-test-build:
	docker-compose -f docker-compose.test.yml build --pull

fmt:
	docker-compose -f docker-compose.test.yml run --rm tests gofmt -s -w .

vet:
	docker-compose -f docker-compose.test.yml run --rm tests go vet -v ./...

test: docker-test-build fmt vet
	docker-compose -f docker-compose.test.yml run --rm tests

testall: test dev
	echo '{{ getEnv "TEST1" }}' > test/output/a.conf.tpl
	echo '{{ getEnv "TEST2" }}' > test/output/b.conf.tpl
	echo '{{ getEnv "TEST2" }}' > test/output/c.conf.tpl_new
	echo '{{ getEnv "TEST1" }}' > test/output/d.conf.tpl_single
	echo '{{ range getEnvArray "ARRAY"}}{{.}}{{ end }}' > test/output/e.conf.tpl_array
	bats test/bats/tests.bats

build: test
	@rm -fr pkg
	@mkdir pkg
	@for distro in ${BUILD_DISTROS}; do \
		GOOS=$${distro} go build -ldflags "${GO_LDFLAGS}" -o pkg/$${distro}/filler; \
		cd pkg/$${distro}; \
		tar -czf ../filler-$${distro}-amd64.tar.gz filler; \
		cd ../..; \
	done

release: build
	@for distro in ${BUILD_DISTROS}; do \
		AWS_PROFILE=production aws s3 cp --acl public-read \
			pkg/filler-$${distro}-amd64.tar.gz s3://airhelp-devops-binaries/filler/${FILLER_VERSION}/filler-$${distro}-amd64.tar.gz; \
		shasum -a 256 pkg/filler-$${distro}-amd64.tar.gz; \
	done
dev:
	go build
