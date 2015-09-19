package appversion

import (
	"os"
	"runtime"
	"strings"
	"strconv"
	"path/filepath"
)

const (
	VERSION_NONE = iota
	VERSION_ONELINER
	VERSION_ONELINER_DEV
	VERSION_STANDARD
	VERSION_DEV
	VERSION_DEBUG
	VERSION_DEBUG_CHECK
)

// set to "" to remove version mentioning
var Version = "0.0.0"
var VersionVDelimiter = "" // can be . for v.1.0.0

func GetVersionString(buildVars map[string]string, mode int) string {
	debugVersionStr := buildVars["DebugMode"]
	debugVersion := debugVersionStr != "" && debugVersionStr!="false" && debugVersionStr!="0"
	if mode == VERSION_DEBUG_CHECK {
		if debugVersion {
			return "1"
		} else {
			return "0"
		}
	}

	author := buildVars["Author"]
	app := buildVars["AppName"]
	if unknown(app) {
		app = filepath.Base(os.Args[0])
	}
	version := buildVars["Version"]
	tag := buildVars["Tag"]
	if unknown(version) && unknown(tag) {
		version = Version
	}

	buildTime := buildVars["BuildTime"]
	goVersion := buildVars["GoVersion"]
	if unknown(goVersion) {
		goVersion = runtime.Version()
	}
	codeRev := buildVars["CodeRev"]
	modifiedSourcesStr := buildVars["ModifiedSources"]
	modifiedSources := []string{}
	if modifiedSourcesStr != "" {
		modifiedSourcesStr = strings.Trim(modifiedSourcesStr, ";")
		modifiedSources = strings.Split(modifiedSourcesStr, ";")
	}

	depsStr := buildVars["Deps"]
	depsStrs := []string{}
	if depsStr != "" {
		depsStr = strings.Trim(depsStr, ";")
		depsStrs = strings.Split(depsStr, ";")
	}
	dirtyDeps := []string{}
	for _, dep := range depsStrs {
		if !strings.Contains(dep, "clean/") {
			dirtyDeps = append(dirtyDeps, dep)
		}
	}

	branch := buildVars["Branch"]

	site := buildVars["Site"]

	//-----------------------------------------------

	appversion := app
	if debugVersion {
		appversion += " DEBUG BUILD"
	}

	if known(version) {
		if version[0] == 'v' {
			appversion += " "+version
		} else {
			appversion += " v"+VersionVDelimiter+version
		}
	}

	if known(tag) && version != tag && "v"+version != tag && version != "v"+tag  {
		appversion += " tagged '"+tag+"'"
	}

	if author != "" {
		appversion += " by "+author
	}

	appversiondate := appversion
	if known(buildTime) {
		appversiondate += ", built in "+buildTime
	}

	devoneliner := goVersion
	devmulti := "go version: "+goVersion
	if known(codeRev) {
		devoneliner += ", rev "+codeRev[0:6]+", branch "+branch
		if debugVersionStr != "" {
			devmulti += "\nDEBUG: "+debugVersionStr
		}
		devmulti += "\nrevision: "+codeRev+"\nbranch: "+branch+"\ntag: "+tag
		devmulti += "\nmodified files: "+itoa(len(modifiedSources))+enumlistPretty(modifiedSources,3)
		devmulti += "\nmodified dependencies: "+itoa(len(dirtyDeps))+enumlistPretty(dirtyDeps,3)

		if known(tag) {
			devoneliner += ", tag "+tag
		}
		if modifiedSourcesStr != "" {
			devoneliner += " +"+itoa(len(modifiedSources))+" modified files"
		}
		if len(dirtyDeps) != 0 {
			devoneliner += " +"+itoa(len(dirtyDeps))+" dirty libs"
		}
	}

	switch mode {
		case VERSION_ONELINER:
			return appversiondate
		case VERSION_ONELINER_DEV:
			return appversiondate+"; "+devoneliner
		case VERSION_STANDARD:
			t := ""
			if known(site) {
				t += "\nsite: "+site
			}
			t += "\nbuilt in "+buildTime
			if unknown(buildTime) { t = ""}
			if debugVersion { t += "\n\nTHIS IS DEVELOPMENT VERSION, NOT FOR PRODUCTION"}
			return appversion+t
		case VERSION_DEV:
			return appversiondate+"\n\n"+devmulti
		case VERSION_DEBUG:
			return appversiondate+"\n\n"+devmulti+"\n"+dumpEverything(modifiedSources, depsStrs)
		default:
			return GetVersionString(buildVars, VERSION_DEBUG) // bug somewhere
	}
}


func known(s string) bool {
	return !unknown(s)
}

func unknown(s string) bool {
	return s=="" || s=="<unknown>" || s=="<undefined>" || s=="null"
}

func enumlistPretty (s []string, limit int) string {
	if len(s) > 0 {
		return " ("+enumlist(s, limit)+")"
	}
	return ""
}

func enumlist(s []string, limit int) string {
	ss := s
	add := ""
	if len(ss) > limit {
		ss = ss[0:limit]
		add = ", ..."
	}
	sss := []string{}
	for _, v := range ss {
		sss = append(sss, strings.Replace(strings.Split(v, "=")[0], "github.com/","g/", -1))
	}

	return strings.Join(sss, ", ") + add
}

func itoa(i int) string {
	return strconv.Itoa(i)
}