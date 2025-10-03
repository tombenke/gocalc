package buildinfo

import (
	"fmt"
	"io"
	"os"
	"runtime/debug"
	"strings"
)

// Provides the complete debug.BuildInfo struct
var BuildInfo *debug.BuildInfo

// A dictionary with keys that are listed with the `-tags ...` switch during the build process
var BuildTags map[string]struct{} = make(map[string]struct{})

// It will be true of the build process is executed with the `-tags DEBUG` flag, otherwise it is false
var DebugMode bool
var DebugWriter io.Writer = io.Discard

func init() {
	var ok bool
	BuildInfo, ok = debug.ReadBuildInfo()
	if !ok {
		panic("can not load build info")
	}
	fmt.Printf("BuildInfo: %+v\n", BuildInfo.Settings)
	for _, setting := range BuildInfo.Settings {
		if setting.Key == "-tags" {
			tags := strings.Fields(setting.Value)
			for _, tag := range tags {
				BuildTags[tag] = struct{}{}
			}
		}
	}
	_, DebugMode = BuildTags["DEBUG"]

	if DebugMode {
		DebugWriter = os.Stdout
	}

}
