# License Attributions

jjversion uses several externally maintained libraries.

This markdown file lists the dependencies of jjversion. Copies of each license can be found in the notices folder, for instance: [notices/golang-LICENSE.md](notices/golang-LICENSE.md).

The externally maintained libraries used by jjversion are:

- golang, located at https://golang.org, license from: https://golang.org/LICENSE
- docker-alpine, located at https://github.com/gliderlabs/docker-alpine
- go-git, located at https://github.com/go-git/go-git
- gopkg.in/yaml.v3 - https://github.com/go-yaml/yaml/tree/v3
- Many dependencies of go-git, including:
    - https://github.com/microsoft/go-winio
    - https://github.com/alcortesm/tgz
    - https://github.com/anmitsu/go-shlex
    - https://github.com/armon/go-socks5
    - https://github.com/creack/pty
    - https://github.com/davecgh/go-spew
    - https://github.com/emirpasic/gods
    - https://github.com/flynn-archive/go-shlex
    - https://github.com/gliderlabs/ssh
    - https://github.com/go-git/gcfg
    - https://github.com/go-git/go-billy
    - https://github.com/go-git/go-git-fixtures
    - https://github.com/google/go-cmp
    - https://github.com/imdario/mergo
    - https://github.com/jbenet/go-context
    - https://github.com/jessevdk/go-flags
    - https://github.com/jessevdk/go-flags
    - https://github.com/kevinburke/ssh_config
    - https://github.com/konsorten/go-windows-terminal-sequences
    - https://github.com/kr/pretty
    - https://github.com/kr/text
    - https://github.com/mitchellh/go-homedir
    - https://github.com/niemeyer/pretty
    - https://github.com/pelletier/go-buffruneio
    - https://github.com/pkg/errors
    - https://github.com/pmezard/go-difflib
    - https://github.com/sergi/go-diff
    - https://github.com/sirupsen/logrus
    - https://github.com/src-d/gcfg
    - https://github.com/stretchr/objx
    - https://github.com/stretchr/testify
    - https://github.com/xanzy/ssh-agent
    - https://cs.opensource.google/go/x/crypto
    - https://cs.opensource.google/go/x/net
    - https://cs.opensource.google/go/x/sync
    - https://cs.opensource.google/go/x/sys
    - https://cs.opensource.google/go/x/term
    - https://cs.opensource.google/go/x/text
    - https://cs.opensource.google/go/x/tools
    - https://github.com/go-check/check/tree/v1
    - https://github.com/src-d/go-billy/tree/v4.3.2
    - https://github.com/src-d/go-git-fixtures/tree/v3.5.0
    - https://github.com/src-d/go-git/tree/v4.13.1
    - https://github.com/go-warnings/warnings/tree/v0.1.2
    - https://github.com/go-yaml/yaml/tree/v2.4.0

In addition, some libraries are only developer dependencies, used for the GitHub Actions pipeline or the devcontainer. Note, this is in addition to the libraries listed above:

- Library scripts and Dockerfile template from https://github.com/microsoft/vscode-dev-containers to install dependencies

Note, VS Code Dev Containers is a development tool that utilizes a Docker container as a "full-featured development environment." Information regarding dev containers can be found at: https://code.visualstudio.com/docs/remote/containers

- jq - https://stedolan.github.io/jq/ - https://github.com/stedolan/jq - including dependencies noted in the [jq-LICENSE.md](notices/jq-LICENSE.md)

- build-essential - https://packages.ubuntu.com/bionic/build-essential

- https://github.com/actions/checkout
- https://github.com/actions/setup-go
- https://github.com/docker/setup-buildx-action
- https://github.com/actions/upload-artifact
- https://github.com/actions/download-artifact
- https://github.com/docker/login-action
- https://github.com/github/codeql-action
- https://github.com/stuartleeks/devcontainer-build-run

The following extensions are included as part of the dev container:

- ms-azuretools.vscode-docker
- golang.Go
- eamodio.gitlens
- hbenl.vscode-test-explorer
- ethan-reesor.vscode-go-test-adapter
- ms-vscode.test-adapter-converter
