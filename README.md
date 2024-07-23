# ghproj

Add GitHub Issues and Pull Requests to GitHub Projects.

## Status

This project is still under development.
Please don't use this yet.

## Motivation

I manage a lot of OSS projects, so I have to handle a lot of issues and pull requests.
So I want to manage them using GitHub User and Organization Projects.

e.g.

- [suzuki-shunsuke](https://github.com/users/suzuki-shunsuke/projects/5)
- [aquaproj](https://github.com/orgs/aquaproj/projects/8)
- [lintnet](https://github.com/orgs/lintnet/projects/1)

By executing `ghproj` periodically by [GitHub Actions schedule event](https://docs.github.com/en/actions/using-workflows/events-that-trigger-workflows#schedule), you can add issues and projects to GitHub Projects automatically.

## Install

`ghproj` is a single binary written in Go.
So you only need to put the executable binary into `$PATH`.

1. [Homebrew](https://brew.sh/)

```sh
brew install suzuki-shunsuke/ghproj/ghproj
```

2. [Scoop](https://scoop.sh/)

```sh
scoop bucket add suzuki-shunsuke https://github.com/suzuki-shunsuke/scoop-bucket
scoop install ghproj
```

3. [aqua](https://aquaproj.github.io/)

```sh
aqua g -i suzuki-shunsuke/ghproj
```

4. Download a prebuilt binary from [GitHub Releases](https://github.com/suzuki-shunsuke/ghproj/releases) and install it into `$PATH`

5. Go

```sh
go install github.com/suzuki-shunsuke/ghproj/cmd/ghproj@latest
```

## Usage

```sh
ghproj init # Scafold a configuration file ghproj.yaml
ghproj add # Add issues and pull requests to GitHub Projects
```

## GitHub Access token

GitHub classic personal access token is required.
The scope `read:org` and `project` are required.

> [!CAUTION]
> fine-grained personal access token is unavailable because it doesn't support GitHub Projects.
> https://github.com/orgs/community/discussions/36441
> > There are also some APIs that do not yet support the fine-grained permission model, that we'll be adding support for in time:
> > - Packages
> > - Projects
> > - Notifications

> [!CAUTION]
> GitHub Access token generated by GitHub App is unavailable for GitHub User Project because the permission of GitHub User Project isn't supported

## Configuration

ghproj.yaml

e.g.

```yaml
# ghproj https://github.com/suzuki-shunsuke/ghproj
entries:
  - query: |
      is:open
      archived:false
      -project:suzuki-shunsuke/5
      -label:create
      owner:szksh-lab
      owner:lintnet
    project_id: PVT_kwHOAMtMJ84AQCf4
```

## Run ghproj by GitHub Actions

Please see [Workflow](.github/workflows/update-project.yaml).

## LICENSE

[MIT](LICENSE)
