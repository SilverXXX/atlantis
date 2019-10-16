package raw

import "github.com/runatlantis/atlantis/server/events/yaml/valid"

// DefaultAutoPlanWhenModified is the default element in the when_modified
// list if none is defined.
const DefaultAutoPlanWhenModified = "**/*.tf*"
const DefaultAutoPlanWhenTargetBranch = ".*"

type Autoplan struct {
	WhenModified     []string `yaml:"when_modified,omitempty"`
	Enabled          *bool    `yaml:"enabled,omitempty"`
	WhenTargetBranch *string  `yaml:"when_target_branch,omitempty"`
}

func (a Autoplan) ToValid() valid.Autoplan {
	var v valid.Autoplan
	if a.WhenModified == nil {
		v.WhenModified = []string{DefaultAutoPlanWhenModified}
	} else {
		v.WhenModified = a.WhenModified
	}

	if a.Enabled == nil {
		v.Enabled = true
	} else {
		v.Enabled = *a.Enabled
	}

	if a.WhenTargetBranch == nil {
		v.WhenTargetBranch = DefaultAutoPlanWhenTargetBranch
	} else {
		v.WhenTargetBranch = *a.WhenTargetBranch
	}

	return v
}

func (a Autoplan) Validate() error {
	return nil
}

// DefaultAutoPlan returns the default autoplan config.
func DefaultAutoPlan() valid.Autoplan {
	return valid.Autoplan{
		WhenModified:     []string{DefaultAutoPlanWhenModified},
		Enabled:          valid.DefaultAutoPlanEnabled,
		WhenTargetBranch: DefaultAutoPlanWhenTargetBranch,
	}
}
