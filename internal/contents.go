package internal

// TODO: could use JS (or Lua) -> JSON -> Go pipeline for data flow instead of toml
// - toml may not allow complex structures like I want for nested struct arrays and optional dropping and adding of others

type ResumeContents struct {
	Personal  PersonalInfo  `toml:"personal"`
	Education EducationInfo `toml:"education"`
	Work      WorkInfo      `toml:"work"`
	Projects  ProjectsInfo  `toml:"projects"`
	Skills    SkillsInfo    `toml:"skills"`
}

type PersonalInfo struct {
	Name     string
	Location string
	Email    string
	Phone    string

	// Socials
	Website  string
	LinkedIn string
	GitHub   string

	// TODO: maybe add section for citizenship
}

type EducationInfo struct {
	EducationItems []EducationItem `toml:"item"`
}

type EducationItem struct {
	Institution string
	Location    string
	Degree      string
	Dates       string
	GPA         string
}

type WorkInfo struct {
	WorkItems []WorkItem `toml:"item"`
}

type WorkItem struct {
	JobTitle    string
	Company     string
	Dates       string
	Location    string
	Description []string // NOTE: multiple strings will be bulleted
}

type ProjectsInfo struct {
	ProjectItems []ProjectItem `toml:"item"`
}

type ProjectItem struct {
	Name        string
	Link        string
	Languages   []string
	Description []string
}

type SkillsInfo struct {
	Languages []string
	Libraries []string
	Databases []string
	Tools     []string
	// TODO: figure out how to integrate other things here (libs, frameworks, techs, etc.)
}
