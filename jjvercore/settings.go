package jjvercore

import (
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v3"
)

type versionSettings struct {
	Major_Version_Bump_Message          string
	Minor_Version_Bump_Message          string
	Patch_Version_Bump_Message          string
	Commit_Message_Incrementing_Enabled bool
}

func getVersionSettings() versionSettings {
	// Default settings
	vs := versionSettings{Major_Version_Bump_Message: "^((breaking|major):)",
		Minor_Version_Bump_Message:          "^((feat|feature|minor):)",
		Patch_Version_Bump_Message:          "^((fix|patch|refactor):)",
		Commit_Message_Incrementing_Enabled: false,
	}

	settingsData := &versionSettings{}

	versioningFile, err := ioutil.ReadFile("versioning.yaml")
	if err == nil {
		err = yaml.Unmarshal(versioningFile, settingsData)

		if err == nil {
			if settingsData.Commit_Message_Incrementing_Enabled {
				vs.Commit_Message_Incrementing_Enabled = true
			}

			if strings.TrimSpace(settingsData.Major_Version_Bump_Message) != "" {
				vs.Major_Version_Bump_Message = settingsData.Major_Version_Bump_Message
			}

			if strings.TrimSpace(settingsData.Minor_Version_Bump_Message) != "" {
				vs.Minor_Version_Bump_Message = settingsData.Minor_Version_Bump_Message
			}

			if strings.TrimSpace(settingsData.Patch_Version_Bump_Message) != "" {
				vs.Patch_Version_Bump_Message = settingsData.Patch_Version_Bump_Message
			}
		}
	}

	return vs
}
