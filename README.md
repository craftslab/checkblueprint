# checkblueprint

[![Actions Status](https://github.com/craftslab/checkblueprint/workflows/CI/badge.svg?branch=master&event=push)](https://github.com/craftslab/checkblueprint/actions?query=workflow%3ACI)
[![Go Report Card](https://goreportcard.com/badge/github.com/craftslab/checkblueprint)](https://goreportcard.com/report/github.com/craftslab/checkblueprint)
[![License](https://img.shields.io/github/license/craftslab/checkblueprint.svg?color=brightgreen)](https://github.com/craftslab/checkblueprint/blob/master/LICENSE)
[![Tag](https://img.shields.io/github/tag/craftslab/checkblueprint.svg?color=brightgreen)](https://github.com/craftslab/checkblueprint/tags)



## Introduction

*checkblueprint* is a linter for Blueprint written in Go.



## Prerequisites

- Go >= 1.16.0



## Build

```bash
git clone https://github.com/craftslab/checkblueprint.git

cd checkblueprint
make build
```



## Run

```bash
./checkblueprint --config-file="config.yml" --output-file="output.json"
```



## Docker

```bash
git clone https://github.com/craftslab/checkblueprint.git

cd checkblueprint
docker build --no-cache -f Dockerfile -t craftslab/checkblueprint:latest .
docker run craftslab/checkblueprint:latest /checkblueprint --config-file="/config.yml" --output-file="/output.json"
```



## Usage

```
usage: checkblueprint --config-file=CONFIG-FILE [<flags>]

Lint Flow

Flags:
  --help                     Show context-sensitive help (also try --help-long
                             and --help-man).
  --version                  Show application version.
  --config-file=CONFIG-FILE  Config file (.yml)
  --output-file=OUTPUT-FILE  Output file (.json|.txt|.xlsx)
```



## Settings

*checkblueprint* parameters can be set in the directory [config](https://github.com/craftslab/checkblueprint/blob/master/config).

An example of configuration in [config.yml](https://github.com/craftslab/checkblueprint/blob/master/config/config.yml):

```yaml
TBD
```



## License

Project License can be found [here](LICENSE).



## Reference

- [blueprint](https://opensource.google/projects/blueprint)
- [checkmake](https://github.com/mrtazz/checkmake)
