package internal

type ResumeContents struct {
	Personal  PersonalInfo
	Education EducationInfo
	Work      WorkInfo
	Projects  ProjectsInfo
	Skills    SkillsInfo
	Extras    ExtraInfo
}

type PersonalInfo struct {
	Name           string
	Location       string
	Email          string
	Phone          string
	AdditionalInfo string
	Website        string
	LinkedIn       string
	GitHub         string
}

type EducationInfo struct {
	EducationItems []EducationItem
}

type EducationItem struct {
	Institution    string
	Location       string
	Degree         string
	Dates          string
	Honors         string
	AdditionalInfo []string
}

type WorkInfo struct {
	WorkItems []WorkItem
}

type WorkItem struct {
	JobTitle    string
	Company     string
	Dates       string
	Location    string
	Description []string // NOTE: multiple strings will be bulleted
}

type ProjectsInfo struct {
	ProjectItems []ProjectItem
}

type ProjectItem struct {
	Name        string
	Link        string
	Dates       string
	Tools       string
	Description []string
	Tags        []string
}

type SkillsInfo struct {
	Sections map[string]skillValues
}

type skillValues struct {
	Values       []string
	JoinedValues string
}

type ExtraInfo struct {
	Visible string
	Hidden  string
}
