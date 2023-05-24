# License Attributions

jjversion uses several externally maintained libraries.

This markdown file lists the dependencies of jjversion. Copies of each license can be found in the attributions folder, for instance: [attributions/golang-LICENSE.md](attributions/golang-LICENSE.md).

The externally maintained libraries used by jjversion are:

- golang, located at <https://golang.org>, license from: <https://golang.org/LICENSE>
- docker-alpine, located at <https://github.com/gliderlabs/docker-alpine>
- go-git, located at <https://github.com/go-git/go-git>
- gopkg.in/yaml.v3 - <https://github.com/go-yaml/yaml/tree/v3>
- josephspurrier/goversioninfo - <https://github.com/josephspurrier/goversioninfo>
- akavel/rsrc - <https://github.com/akavel/rsrc>
- Many dependencies of go-git, including:
  - <https://github.com/microsoft/go-winio>
  - <https://github.com/anmitsu/go-shlex>
  - <https://github.com/armon/go-socks5>
  - <https://github.com/creack/pty>
  - <https://github.com/davecgh/go-spew>
  - <https://github.com/emirpasic/gods>
  - <https://github.com/flynn-archive/go-shlex>
  - <https://github.com/gliderlabs/ssh>
  - <https://github.com/go-git/gcfg>
  - <https://github.com/go-git/go-billy>
  - <https://github.com/go-git/go-git-fixtures>
  - <https://github.com/google/go-cmp>
  - <https://github.com/imdario/mergo>
  - <https://github.com/jbenet/go-context>
  - <https://github.com/jessevdk/go-flags>
  - <https://github.com/jessevdk/go-flags>
  - <https://github.com/kevinburke/ssh_config>
  - <https://github.com/konsorten/go-windows-terminal-sequences>
  - <https://github.com/kr/pretty>
  - <https://github.com/kr/text>
  - <https://github.com/mitchellh/go-homedir>
  - <https://github.com/niemeyer/pretty>
  - <https://github.com/pkg/errors>
  - <https://github.com/pmezard/go-difflib>
  - <https://github.com/sergi/go-diff>
  - <https://github.com/sirupsen/logrus>
  - <https://github.com/stretchr/objx>
  - <https://github.com/stretchr/testify>
  - <https://github.com/xanzy/ssh-agent>
  - <https://cs.opensource.google/go/x/crypto>
  - <https://cs.opensource.google/go/x/mod>
  - <https://cs.opensource.google/go/x/net>
  - <https://cs.opensource.google/go/x/sys>
  - <https://cs.opensource.google/go/x/term>
  - <https://cs.opensource.google/go/x/text>
  - <https://cs.opensource.google/go/x/tools>
  - <https://github.com/go-check/check/tree/v1>
  - <https://github.com/go-warnings/warnings/tree/v0.1.2>
  - <https://github.com/go-yaml/yaml/tree/v2.4.0>
  - <https://github.com/ProtonMail/go-crypto>
  - <https://github.com/acomagu/bufpipe>
  - <https://github.com/matryer/is>
  - <https://github.com/cloudflare/circl>
  - <https://github.com/bwesterb/go-ristretto>
  - <https://github.com/mmcloughlin/avo>
  - <https://pkg.go.dev/golang.org/x/arch>
  - <https://pkg.go.dev/rsc.io/pdf>
  - <https://github.com/golang/groupcache>
  - <https://github.com/elazarl/goproxy>
  - <https://github.com/golang/protobuf>
  - <https://pkg.go.dev/google.golang.org/protobuf>

In addition, we utilize two third party alpine packages in the build Docker image:

  - <https://pkgs.alpinelinux.org/package/edge/main/x86_64/gcc>
  - <https://pkgs.alpinelinux.org/package/edge/main/x86_64/musl-dev>

In addition, some libraries are only developer dependencies, used for the GitHub Actions pipeline or the devcontainer. Note, this is in addition to the libraries listed above:

- Library scripts and Dockerfile template from <https://github.com/microsoft/vscode-dev-containers> to install dependencies

Note, VS Code Dev Containers is a development tool that utilizes a Docker container as a "full-featured development environment." Information regarding dev containers can be found at: <https://code.visualstudio.com/docs/remote/containers>

- jq - <https://stedolan.github.io/jq/> - <https://github.com/stedolan/jq> - including dependencies noted in the [jq-LICENSE.md](notices/jq-LICENSE.md)

- build-essential - <https://packages.ubuntu.com/bionic/build-essential>

- <https://github.com/actions/checkout>
- <https://github.com/actions/setup-go>
- <https://github.com/docker/setup-buildx-action>
- <https://github.com/actions/upload-artifact>
- <https://github.com/actions/download-artifact>
- <https://github.com/docker/login-action>
- <https://github.com/github/codeql-action>
- <https://github.com/devcontainers/ci>
- <https://github.com/marvinpinto/action-automatic-releases>
- <https://github.com/Checkmarx/kics-github-action>

The following extensions are included as part of the dev container:

- ms-azuretools.vscode-docker
- golang.Go
- eamodio.gitlens
- hbenl.vscode-test-explorer
- ethan-reesor.vscode-go-test-adapter
- ms-vscode.test-adapter-converter
