package appversion
import (
	"os"
	"strings"
)

var DangerousParamCheck = true

func GetVersionString_AccordingToCmdParams(BuildVars map[string]string, devModeActivationParam string, deepDevModeActivationParam string) string {
	args := os.Args[1:]
	if len(args) == 0 {
		return ""
	}
	if !paramsEq(args[0], "version") {
		args = args[1:]
		if len(args) == 0 {
			return ""
		}
	}

	devMode := false
	deepDevMode := false
	oneliner := false
	debugcheck := false

	if paramsEq(args[0], "version") {
		for _, arg := range args {
			if paramsEq(arg, devModeActivationParam) {
				devMode = true
			}
			if paramsEq(arg, deepDevModeActivationParam) {
				deepDevMode = true
				devMode = true
			}
			if paramsEq(arg, "oneliner") {
				oneliner = true
			}
			if paramsEq(arg, "debug-check") {
				debugcheck = true
			}
		}
	} else {
		return ""
	}

	if debugcheck {
		return GetVersionString(BuildVars, VERSION_DEBUG_CHECK)
	} else if oneliner {
		if devMode {
			return GetVersionString(BuildVars, VERSION_ONELINER_DEV)
		} else {
			return GetVersionString(BuildVars, VERSION_ONELINER)
		}
	} else {
		if deepDevMode {
			return GetVersionString(BuildVars, VERSION_DEBUG)
		} else if devMode {
			return GetVersionString(BuildVars, VERSION_DEV)
		} else {
			return GetVersionString(BuildVars, VERSION_STANDARD)
		}
	}
	return ""
}

func ProbablyOutputVersionAndExit(BuildVars map[string]string, devModeActivationParam string, deepDevModeActivationParam string) {
	x := GetVersionString_AccordingToCmdParams(BuildVars, devModeActivationParam, deepDevModeActivationParam)
	if x != "" {
		println(x)
		if (x == "1") {
			os.Exit(1) // debug-check confirmed that it is debug version
		}
		os.Exit(0)
	}
}

func paramsEq(param, shouldmatch string) bool {
	param = strings.ToLower(param)
	shouldmatch = strings.ToLower(shouldmatch)
	if param == "--"+shouldmatch || param=="-"+shouldmatch || (DangerousParamCheck && param == shouldmatch) {
		return true
	}
	return false
}