# Leetgo

Best LeetCode friend for geek.

---

[![CI](https://github.com/j178/leetgo/actions/workflows/ci.yaml/badge.svg)](https://github.com/j178/leetgo/actions/workflows/ci.yaml)

中文 | [English](./README.md)

Leetgo is a command line tool that generates skeleton code for LeetCode questions. You can run and debug test cases locally with your favorite IDE.
Then you can submit your code to LeetCode directly.

**This project is in its early development stage, and anything is likely to change.**

## Main features

- Search for and view a question by its ID or slug.
- Generate skeleton code and testing code for a question.
- Run test cases on your local machine.
- Generate contest questions just in time.

## Install `leetgo`

You can download the latest binary from the [release page](https://github.com/j178/leetgo/releases).

### go install

```shell
go install github.com/j178/leetgo@latest
```

### brew install

```shell
brew install j178/tap/leetgo
```

## Supported languages

- Golang
- Python
- Rust

### Planning

- Java
- C++

## Usage
<!-- BEGIN USAGE -->
```
Usage:
  leetgo [command]

Available Commands:
  init                    Init a leetcode workspace
  pick                    Generate a new question
  today                   Generate the question of today
  info                    Show question info
  test                    Run question test cases
  submit                  Submit solution
  contest                 Generate contest questions
  cache                   Manage local questions cache
  config                  Show leetgo config dir

Flags:
  -v, --version      version for leetgo
  -g, --gen string   language to generate: cpp, go, python ...

Use "leetgo [command] --help" for more information about a command.
```
<!-- END USAGE -->

## Config file
<!-- BEGIN CONFIG -->
```yaml
# Generate code for questions, go, python, ... (will be override by project config and flag --gen)
gen: go
# Language of the questions, zh or en
language: zh
# LeetCode configuration
leetcode:
  # LeetCode site, https://leetcode.com or https://leetcode.cn
  site: https://leetcode.cn
editor: {}
go:
  # Output directory for Go files
  out_dir: go
  # Generate separate package for each question
  separate_package: true
  # Filename template for Go files
  filename_template: ""
python:
  # Output directory for Python files
  out_dir: python
```
<!-- END CONFIG -->

## Credits

- https://github.com/EndlessCheng/codeforces-go
- https://github.com/clearloop/leetcode-cli
- https://github.com/budougumi0617/leetgode