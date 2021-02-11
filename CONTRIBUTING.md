# :alembic: Contributing

- [:alembic: Contributing](#alembic-contributing)
  - [:see_no_evil: Introduction](#see_no_evil-introduction)
  - [:package: Third-Party Packages](#package-third-party-packages)
  - [:seedling: Requirements](#seedling-requirements)
  - [:construction_worker: Setup](#construction_worker-setup)
  - [:hammer: Development Workflow](#hammer-development-workflow)
  - [:building_construction: Build](#building_construction-build)
  - [:card_file_box: Structure](#card_file_box-structure)
  - [:wrench: Tests](#wrench-tests)
  - [:triangular_flag_on_post: Pull Request Process](#triangular_flag_on_post-pull-request-process)

---

## :see_no_evil: Introduction

Kinto CLI is written in [GoLang](https://golang.org/) using the popular package [Cobra](https://github.com/spf13/cobra) and uses [Go Modules](https://github.com/golang/go/wiki/Modules) to make working with external dependencies easy.

## :package: Third-Party Packages

Kinto CLI uses some third-party packages to enable extra features. A list of all those packages can be found below -

- [spf13/cobra](https://github.com/spf13/cobra)
- [spf13/viper](https://github.com/spf13/viper)
- [jpillora/chisel](https://github.com/jpillora/chisel)
- [AlecAivazis/Survey](https://github.com/AlecAivazis/survey)
- [BrianDowns/Spinner](https://github.com/briandowns/spinner)
- [olekukonko/tablewriter](https://github.com/olekukonko/tablewriter)
- [Terry-Mao/goconf](https://github.com/Terry-Mao/goconf)

## :seedling: Requirements

- [Go v1.13+](https://golang.org/doc/install)
- [Git](https://git-scm.com/downloads)

## :construction_worker: Setup

To start contributing to Kinto CLI you can follow your existing GoLang workflow by forking, working on a branch and submitting a PR.

If you are new to working with GoLang repositories, follow the below-given instructions to set up the project locally.

- Fork the repo on GitHub
- Clone the repo and cd into it
  
  ```
  git clone https://github.com/<your-username>/kinto-cli

  cd kinto-cli
  ```
- Add remotes for the upstream branch

  ```
  git remote add upstream https://github.com/kintohub/kinto-cli
  
  git checkout -b <branch-name>
  ```
- Work on the feature/fix
- Commit, push and submit PR

## :hammer: Development Workflow

Once you have a local copy of Kinto-CLI repo on your machine, you can use an IDE such as [Goland](https://www.jetbrains.com/go/download/) to make working with GoLang easier or you can do the setup manually by -

```
cd kinto-cli

go mod download
```

## :building_construction: Build

You can build the binary file for your Operating system by running the command

```
go build
```

And install it to `$GOPATH/go/bin/` by

```
go install
```

## :card_file_box: Structure

Kinto CLI follows the following structure

```
kinto-cli
├── internal
│   ├── api
│   ├── cli
│   ├── config
│   ├── controller
│   ├── types
│   └── utils
└── proto
```

- **internal** houses all the code for the CLI
  - **api** provides access to different components of the CLI such as `auth`, `clusters`, `services/blocks`, etc via a single API.
  - **cli** contains the root file that houses all the command declarations.
  - **config** stores the consts and configs for the CLI.
  - **controller** hosts the business logic for all the functions declared in **cli**
  - **types** has the compiled proto files needed for the CLI.
  - **utils** contains the basic utils functions needed by the CLI.
- **proto** contains the proto files needed to connect to KintoHub servers.

## :wrench: Tests

When you're adding a new feature/fixing bugs, it is encouraged to add integration tests (unit tests also if possible) for the CLI. You should run all the tests and make sure everything passes before submitting the PR.

## :triangular_flag_on_post: Pull Request Process

1. Ensure any install or build dependencies are removed before the end of the layer when doing a
   build.
2. Increase the version numbers in any examples files and the README.md to the new version that this
   Pull Request would represent. The versioning scheme we use is [SemVer](http://semver.org/).
3. You may merge the Pull Request in once you have the sign-off of two other developers, or if you
   do not have permission to do that, you may request the second reviewer to merge it for you.
