go_version = $(shell go version)
commit_id = $(shell git rev-parse HEAD)
branch_name = $(shell git name-rev --name-only HEAD)
build_time = $(shell date -u '+%Y-%m-%d_%H:%M:%S')
app_version = 1.0.0
version_package = newframework/pkg/version
app_name = newframework
work_dir = target
all: package

build: target
	@go build -ldflags \
	"-X ${version_package}.CommitId=${commit_id} \
	-X ${version_package}.BranchName=${branch_name} \
	-X ${version_package}.BuildTime=${build_time} \
	-X ${version_package}.AppVersion=${app_version}" -v \
	-o ${work_dir}/${app_name} ./cmd/base/.

target:
	@mkdir ${work_dir}