package appversion
import "strings"

type ModifiedSource struct {
	Name string
	Timestamp string
	Date string
}

type Dep struct {
	Name string
	State string
	Branch string
	Commit string
	ModificationDate string
}

func dumpEverything(modifiedSources, depsStrs []string) string {
	res := ""

	modsrcs := []*ModifiedSource{}
	for _, v := range modifiedSources {
		a := strings.Split(v,"=")
		ts := ""
		date := ""
		if len(a)>1 {
			b := strings.SplitN(a[1], "/", 2)
			if len(b) >1 {
				ts = b[0]
				date = b[1]
			}
		}
		modsrcs = append(modsrcs, &ModifiedSource{a[0], ts, date})
	}

	depsClean := []*Dep{}
	depsDirty := []*Dep{}
	depsUntracked := []*Dep{}
	for _, v := range depsStrs {
		a := strings.Split(v,"=")
		state := ""
		branch := ""
		commit := ""
		moddate := ""
		if len(a)>1 {
			b := strings.SplitN(a[1], "/", 4)
			if len(b) == 4 {
				state = b[0]
				branch = b[1]
				commit = b[2]
				moddate = b[3]
				if state == "clean" {
					depsClean = append(depsClean, &Dep{a[0], state, branch, commit, moddate})
				} else {
					depsDirty = append(depsDirty, &Dep{a[0], state, branch, commit, moddate})
				}
			} else {
				depsUntracked = append(depsUntracked, &Dep{a[0], "", "", "", ""})
			}
		} else {
			depsUntracked = append(depsUntracked, &Dep{a[0], "","","",""})
		}
	}

	res += "\nmodified sources: "
	for _, v := range modsrcs {
		res += "\n - "+v.Name+" ("+v.Date+")"
	}

	res += "\ndirty dependencies: "
	for _, v := range depsDirty {
		res += "\n - "+v.Name+" rev "+v.Branch+"/"+v.Commit+" ("+prettyTime(v.ModificationDate)+")"
	}

	res += "\ncommon dependencies: "
	for _, v := range depsClean {
		res += "\n - "+v.Name+" rev "+v.Branch+"/"+v.Commit+" ("+prettyTime(v.ModificationDate)+")"
	}

	res += "\ndependencies without repo: "
	for _, v := range depsUntracked {
		res += "\n - "+v.Name
	}


	return res
}

func prettyTime(s string)string {
	return strings.Replace(s, "T", " ", -1)
}