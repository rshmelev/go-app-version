package main

import "github.com/rshmelev/go-app-version/appversion"

func main() {
	vars := map[string]string{}
	vars["DebugMode"] = "true"
	vars["Author"] = "John Doe"
	vars["AppName"] = "testutil"
	vars["Version"] = "v1.0.1"
	vars["Tag"] = "v1.0.1"
	vars["BuildTime"] = "2015-09-19 15:40:24"
	vars["GoVersion"] = "go1.5.1 darwin/amd64"
	vars["CodeRev"] = "757f527935f24ecc097cf1f93ccf1998c627700a"
	vars["Branch"] = "master"
	vars["Site"] = "http://some-super-util-site.com"

	vars["ModifiedSources"] = ";build.sh=1442685424/Sat Sep 19 17:57:04 UTC 2015;code/main.go=1442673580/Sat Sep 19 14:39:40 UTC 2015;deletedfile.go=deleted"
	vars["Deps"] = ";github.com/gorilla/websocket=clean/master/a3ec486e6a7a41858210b0fc5d7b5df593b3c4a3/2015-05-29T20:03:52-07:00;github.com/rshmelev/go-app-version/appversion=unknown;github.com/rshmelev/go-uniconv=dirty/master/b73eb0c2c31c76c613cb955bec63592ac02f0c08/2015-08-07T22:35:18+03:00;github.com/rshmelev/go-uniconv/first=dirty/master/b73eb0c2c31c76c613cb955bec63592ac02f0c08/2015-08-07T22:35:18+03:00"

	appversion.ProbablyOutputVersionAndExit(vars, "dev", "dev-details")
	println("hm... looks like you didn't call this app with 'version' param ... ")
}
