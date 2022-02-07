# checkblueprint

[![Build Status](https://github.com/craftslab/checkblueprint/workflows/CI/badge.svg?branch=main&event=push)](https://github.com/craftslab/checkblueprint/actions?query=workflow%3ACI)
[![codecov](https://codecov.io/gh/craftslab/checkblueprint/branch/main/graph/badge.svg?token=etdHpieNbG)](https://codecov.io/gh/craftslab/checkblueprint)
[![Go Report Card](https://goreportcard.com/badge/github.com/craftslab/checkblueprint)](https://goreportcard.com/report/github.com/craftslab/checkblueprint)
[![License](https://img.shields.io/github/license/craftslab/checkblueprint.svg)](https://github.com/craftslab/checkblueprint/blob/main/LICENSE)
[![Release](https://img.shields.io/github/release/craftslab/checkblueprint.svg)](https://github.com/craftslab/checkblueprint/releases/latest)
[![Gitter chat](https://badges.gitter.im/craftslab/craftslab.png)](https://gitter.im/craftslab/craftslab)



## Introduction

*checkblueprint* is a linter for Blueprint written in Go.



## Prerequisites

- Go >= 1.17.0



## Run

```bash
git clone https://github.com/craftslab/checkblueprint.git

cd checkblueprint
version=latest make build
./bin/checkblueprint --config-file="$PWD/config/config.yml"
```



## Docker

```bash
git clone https://github.com/craftslab/checkblueprint.git

cd checkblueprint
version=latest make docker
docker run -v "$PWD"/config:/tmp ghcr.io/craftslab/checkblueprint:latest --config-file="/tmp/config.yml"
```



## Usage

```
TBD
```



## Settings

*checkblueprint* parameters can be set in the directory [config](https://github.com/craftslab/checkblueprint/blob/main/config).

An example of configuration in [config.yml](https://github.com/craftslab/checkblueprint/blob/main/config/config.yml):

```yaml
TBD
```



## License

Project License can be found [here](LICENSE).



## Reference

- [blueprint](https://opensource.google/projects/blueprint)
- [checkmake](https://github.com/mrtazz/checkmake)
