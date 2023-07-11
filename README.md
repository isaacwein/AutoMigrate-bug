# GORM Playground

GORM Playground can be used to play GORM and reports issues, if you encounter a bug in GORM, please report it at [https://github.com/go-gorm/gorm/issues](https://github.com/go-gorm/gorm/issues) with the Playground Pull Request's link

[![test status](https://github.com/go-gorm/playground/workflows/tests/badge.svg?branch=master "test status")](https://github.com/go-gorm/playground/actions)

### Quick Start

##### 1. [Fork this repo](https://docs.github.com/en/free-pro-team@latest/github/getting-started-with-github/fork-a-repo)


```bash


# Run tests with specfied database
GORM_DIALECT=mysql go test
```

##### 5. Modify tests and make it fail

##### 6. [Create Playground Pull Request](https://docs.github.com/en/free-pro-team@latest/github/collaborating-with-issues-and-pull-requests/creating-a-pull-request) and [Create a GORM issue](https://github.com/go-gorm/gorm/issues/new?template=bug_report.md) with the link

### Advanced Usage



We have prepared some structs with relationships in [https://github.com/go-gorm/playground/blob/master/models.go](https://github.com/go-gorm/playground/blob/master/models.go) that you can use for your tests

## Happy Hacking!
