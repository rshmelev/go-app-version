# go-app-version library

Used to simplify process of app version output when 'version' param is passed.
There are many useful options - for developer and common users.


```
yourapp version
yourapp version oneliner
yourapp version debug-check  <-- useful for deploy tools to check if this version is prod one!
yourapp version dev          <-- revision, branch, ...
yourapp version dev oneliner
yourapp version dev-details  <-- includes info about deps and modified files
```

NOTE: 'dev' and 'dev-details' param names are configurable

How to use in your app:

```
appversion.ProbablyOutputVersionAndExit(someBuildVars, "dev", "dev-details")
```

... where `someBuildVars` is `map[string]string` and can contain following keys:

- DebugMode - should be "true" or "false"
- Author
- AppName
- Version
- Tag (git tag)
- BuildTime
- GoVersion
- CodeRev
- ModifiedSources
- Deps
- Branch
- Site

Please look at `appversion/testing/all_possible_runs.sh` and `appversion/testing/main.go` to see it in action

Sample run of `appversion/testing/all_possible_runs.sh` should look like this:

```
$ sh all_possible_runs.sh
  
......... now testing 'app'

hm... looks like you didn't call this app with 'version' param ...

......... now testing 'app version'

testutil DEBUG BUILD v1.0.1 by John Doe
site: http://some-super-util-site.com
built in 2015-09-19 15:40:24

THIS IS DEVELOPMENT VERSION, NOT FOR PRODUCTION

......... now testing 'app version oneliner'

testutil DEBUG BUILD v1.0.1 by John Doe, built in 2015-09-19 15:40:24

......... now testing 'app version dev'

testutil DEBUG BUILD v1.0.1 by John Doe, built in 2015-09-19 15:40:24

go version: go1.5.1 darwin/amd64
DEBUG: true
revision: 757f527935f24ecc097cf1f93ccf1998c627700a
branch: master
tag: v1.0.1
modified files: 3 (build.sh, code/main.go, deletedfile.go)
modified dependencies: 3 (g/rshmelev/go-app-version/appversion, g/rshmelev/go-uniconv, g/rshmelev/go-uniconv/first)

......... now testing 'app version dev-details'

testutil DEBUG BUILD v1.0.1 by John Doe, built in 2015-09-19 15:40:24

go version: go1.5.1 darwin/amd64
DEBUG: true
revision: 757f527935f24ecc097cf1f93ccf1998c627700a
branch: master
tag: v1.0.1
modified files: 3 (build.sh, code/main.go, deletedfile.go)
modified dependencies: 3 (g/rshmelev/go-app-version/appversion, g/rshmelev/go-uniconv, g/rshmelev/go-uniconv/first)

modified sources:
- build.sh (Sat Sep 19 17:57:04 UTC 2015)
- code/main.go (Sat Sep 19 14:39:40 UTC 2015)
- deletedfile.go ()
dirty dependencies:
- github.com/rshmelev/go-uniconv rev master/b73eb0c2c31c76c613cb955bec63592ac02f0c08 (2015-08-07 22:35:18+03:00)
- github.com/rshmelev/go-uniconv/first rev master/b73eb0c2c31c76c613cb955bec63592ac02f0c08 (2015-08-07 22:35:18+03:00)
common dependencies:
- github.com/gorilla/websocket rev master/a3ec486e6a7a41858210b0fc5d7b5df593b3c4a3 (2015-05-29 20:03:52-07:00)
dependencies without repo:
- github.com/rshmelev/go-app-version/appversion

......... now testing 'app version dev oneliner'

testutil DEBUG BUILD v1.0.1 by John Doe, built in 2015-09-19 15:40:24; go1.5.1 darwin/amd64, rev 757f52, branch master, tag v1.0.1 +3 modified files +3 dirty libs

......... now testing 'app version debug-check'

1
exit status 1
```
