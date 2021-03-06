Pipeline
========

This tool is pipeline framework.

[![Travis](https://img.shields.io/travis/linyows/pipeline.svg?style=flat-square)][travis]
[![GitHub release](http://img.shields.io/github/release/linyows/pipeline.svg?style=flat-square)][release]
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]
[![Go Documentation](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)][godocs]

[travis]: https://travis-ci.org/linyows/pipeline
[release]: https://github.com/linyows/pipeline/releases
[license]: https://github.com/linyows/pipeline/blob/master/LICENSE
[godocs]: http://godoc.org/github.com/linyows/pipeline

Built-in Plugins
----------------

- Github Comment
- Github Pull-Request
- Slack Notification

Usage
-----

```sh
cat << EOF > .pipeline.yml
variables:
  %BRANCH%: $BRANCH
  %BASE%: $BASE
  %ISSUE_ID%: $ISSUE_ID
  %GITHUB_TOKEN%: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
lines:
  parallels:
    - name: build branch
      commands: |
        git checkout -f %BRANCH%
        bin/rspec
        %BRANCH_PERCENT% = cat coverage/.last_run.json | grep covered_percent | awk '{print $2}'"
    - name: build master
      commands: |
        git checkout -f %BASE%
        bin/rspec
        %BASE_PERCENT% = cat coverage/.last_run.json | grep covered_percent | awk '{print $2}'"
  - name: calculate coverage
    commands: |
      %PERCENT% = $(./calculate.sh)
      %COMMENT% = $(./build_comment.sh)
      exit %EXIT_STATUS%
    plugins:
      - github-comment
      - github-pr
  - name: notify to slack
    plugins:
      - slack-notify
EOF
```

Install
-------

To install, use `go get`:

```sh
$ go get -d github.com/linyows/pipeline
```

Contribution
------------

1. Fork ([https://github.com/linyows/pipeline/fork](https://github.com/linyows/pipeline/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

Author
------

[linyows](https://github.com/linyows)
