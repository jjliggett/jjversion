# jjversion

A basic versioning application for git projects, enabled by [go-git](https://github.com/go-git/go-git).

Versioning is obtained from the following sources, in this order:

1. Release branch names if ona release branch, e.g. release/1.42.0
2. Version tag on the current commit, e.g. v1.42.0
3. If commit_message_incrementing_enabled is **true** in ***versioning.yaml***: combination of semantic commits (major/minor/patch) and version tagged commits

Combining this with conventional commits can work well: <https://www.conventionalcommits.org/en/v1.0.0/>

## Sample Output

```json
{
  "Major": 0,
  "Minor": 1,
  "Patch": 0,
  "MajorMinorPatch": "0.1.0",
  "Sha": "f727c93d0afeefa52f2f28e03c3aae98b0854379",
  "ShortSha": "f727c93"
}
```

## Docker Image

In addition to binaries, a Docker image is available. The image can be run as follows:

```sh
docker run --rm -v "$(pwd):/repo" jjliggett/jjversion
```

Within the devcontainer, the image can be run as follows:

```sh
docker run --rm -v "$LOCAL_WORKSPACE_FOLDER:/repo" jjliggett/jjversion
```

The image is available on both GitHub Container Registry and Docker Hub.

## GitHub Action

A custom GitHub action is available which installs the jjversion package, executes jjversion, and parses the JSON output into several outputs for use in GitHub workflows. Details can be found in the action repository: <https://github.com/jjliggett/jjversion-action>.

## Licensing

Licensing for this application can be found at: [LICENSE.md](LICENSE.md).

The jjversion license applies to all parts of jjversion that are not
externally maintained libraries and dependencies.

Attributions for dependencies of jjversion and developer
dependencies can be found at [docs/NOTICE.md](docs/NOTICE.md).

## Build Locally

```sh
go build -ldflags "-X main.appVersion=42.10.0"
```
