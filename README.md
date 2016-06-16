# godog-issue

Repository to demonstrate https://github.com/DATA-DOG/godog/issues/35

## Steps to reproduce
* go to directory api
* run `godog ./features`

## observed behavior
it doesn't use the vendored copy of godog and instead reports an error (line breaks from me):
```
user@host:~/gopath/src/github.com/nightlyone/godog-issue/api (master)$ godog ./features/
failed to compile testmain package:
/tmp/go-build773782062/github.com/nightlyone/godog-issue/api/_test/_testmain.go:12: 
cannot use suite (type *"github.com/DATA-DOG/godog".Suite) as type *"github.com/nightlyone/godog-issue/vendor/github.com/DATA-DOG/godog".Suite in argument to main.FeatureContext
``` 
## expected behavior
Use the vendored copy of the godog package to run the features.

## versions
* godog: v0.5.0
* go: go version go1.6.2 linux/amd64
* govendor: https://github.com/kardianos/govendor

