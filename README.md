# ghproj

Add GitHub Issues and Pull Requests to GitHub Projects.

## Motivation

I manage a lot of OSS projects, so I have to handle a lot of issues and pull requests.
So I want to manage them using a GitHub Project.

I've developed this tool to gather issues and pull requests of all my OSS in a single GitHub Project.
By executing this tool periodically by [GitHub Actions schedule event](https://docs.github.com/en/actions/using-workflows/events-that-trigger-workflows#schedule), you can add issues and projects to GitHub Projects automatically.

## Blog Post

- [Japanese](https://zenn.dev/shunsuke_suzuki/articles/add-github-issue-pr-to-project)

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

There are two options.

1. (Recommendation) If you use GitHub **Organization** Project, you can use a `GitHub App`
1. If you use GitHub **User** Project, you can use a `classic personal access token`

GitHub Actions token is unavailable to manage GitHub Projects.

fine-grained personal access token is unavailable because it doesn't support GitHub Projects.

https://github.com/orgs/community/discussions/36441

> There are also some APIs that do not yet support the fine-grained permission model, that we'll be adding support for in time:
> - Packages
> - Projects
> - Notifications

GitHub App is unavailable for GitHub User Projects because the permission of GitHub User Project isn't supported.

### 1. GitHub App

Permissions:

- `Repository permissions`: `metadata: read-only`
- `Organization permissions`: `Projects: Read and write`

Installed repositories:

Please install the GitHub App into only a repository where `ghproj` is executed via GitHub Actions.

### 2. classic personal access token

The scope `read:org` and `project` are required.

## Configuration

ghproj.yaml

e.g.

```yaml
entries:
  - query: |
      is:open
      archived:false
      -project:suzuki-shunsuke/5
      -label:create
      owner:szksh-lab
      owner:lintnet
    expr: |
      (! Item.Repo.IsFork) &&
      (Item.Title != "Dependency Dashboard") &&
      ! (Item.Repo.Name startsWith "homebrew-") &&
      ! (Item.Repo.Name startsWith "test-")
    project_id: PVT_kwHOAMtMJ84AQCf4
```

- `query`: GitHub GraphQL Query to search issues and pull requests which are added to a GitHub Project
- `expr`: An expression to filter the search result. [expr-lang/expr](https://github.com/expr-lang/expr) is used. The expression is evaluated per item. The evaluation result must be a boolean. If the result is `false`, the item is excluded. `expr` is optional

`Item`:

```json
{
  "Title": "issue or pull request title",
  "Repo": {
    "Owner": "repository owner name",
    "Name": "repository name",
    "IsArchived": false,
    "IsFork": false
  }
}
```

- `project_id`: GitHub Project id which issues and pull requests are added. You can get your project id using GitHub CLI `gh project list`

```sh
gh project list
```

## Archive items

You can archive items by `ghproj add` command.

```sh
ghproj add
```

ghproj.yaml

```yaml
entries:
  - expr: |
      Item.Repo.IsArchived
    action: archive
    project_id: PVT_kwHOAMtMJ84AQCf4
```

`Item`:

```json
{
  "State": "CLOSED",
  "Title": "issue or pull request title",
  "Labels": ["enhancement"],
  "Open": false,
  "Author": "octokit",
  "Repo": {
    "Owner": "repository owner name",
    "Name": "repository name",
    "IsArchived": false,
    "IsFork": false
  }
}
```

## Run ghproj by GitHub Actions

- [GitHub Actions Workflow](https://github.com/szksh-lab/.github/blob/main/.github/workflows/update-project.yaml)
- [Configuraition](https://github.com/szksh-lab/.github/blob/main/ghproj.yaml)
- [GitHub Project](https://github.com/orgs/szksh-lab/projects/1)

The workflow is executed periodically by [GitHub Actions schedule event](https://docs.github.com/en/actions/using-workflows/events-that-trigger-workflows#schedule), and issues and pull requests are added to the GitHub Project.

## LICENSE

[MIT](LICENSE)
