package buildinfo

import (
	"fmt"
	"runtime/debug"
	"strings"
)

var BuildInfo *debug.BuildInfo
var BuildTags map[string]struct{} = make(map[string]struct{})
var DebugMode bool

func init() {
	var ok bool
	BuildInfo, ok = debug.ReadBuildInfo()
	if !ok {
		panic("can not load build info")
	}
	fmt.Printf("BuildInfo: %+v\n", BuildInfo.Settings)
	for _, setting := range BuildInfo.Settings {
		if setting.Key == "-tags" {
			// A tag-ek szóközökkel elválasztva szerepelnek
			tags := strings.Fields(setting.Value)
			for _, tag := range tags {
				BuildTags[tag] = struct{}{}
			}
		}
	}
	_, DebugMode = BuildTags["DEBUG"]
}
