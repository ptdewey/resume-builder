package internal

import (
	"os"

	"github.com/BurntSushi/toml"
	lua "github.com/yuin/gopher-lua"
)

func ParseTomlResumeContents(contentsPath string) (ResumeContents, error) {
	var out ResumeContents

	bytes, err := os.ReadFile(contentsPath)
	if err != nil {
		return out, err
	}

	if _, err := toml.Decode(string(bytes), &out); err != nil {
		return out, nil
	}

	return out, nil
}

func ParseLuaResumeContents(scriptPath string) (ResumeContents, error) {
	var out ResumeContents

	L := lua.NewState()
	defer L.Close()

	if err := L.DoFile(scriptPath); err != nil {
		return out, err
	}

	luaTable := L.Get(-1)
	if tbl, ok := luaTable.(*lua.LTable); ok {
		tbl.ForEach(func(key lua.LValue, value lua.LValue) {
			switch key.String() {
			case "personal":
				if personalTbl, ok := value.(*lua.LTable); ok {
					parsePersonalLuaTable(&out, personalTbl)
				}
			case "education":
				if educationTbl, ok := value.(*lua.LTable); ok {
					parseEducationLuaTable(&out, educationTbl)
				}
			case "work":
				if workTbl, ok := value.(*lua.LTable); ok {
					parseWorkLuaTable(&out, workTbl)
				}
			case "projects":
				if projectsTbl, ok := value.(*lua.LTable); ok {
					parseProjectsLuaTable(&out, projectsTbl)
				}
			case "skills":
				if skillsTbl, ok := value.(*lua.LTable); ok {
					parseSkillsLuaTable(&out, skillsTbl)
				}
			}
		})
	}

	return out, nil
}

func parsePersonalLuaTable(contents *ResumeContents, tbl *lua.LTable) {
	tbl.ForEach(func(k lua.LValue, v lua.LValue) {
		switch k.String() {
		case "name":
			contents.Personal.Name = v.String()
		case "location":
			contents.Personal.Location = v.String()
		case "email":
			contents.Personal.Email = v.String()
		case "phone":
			contents.Personal.Phone = v.String()
		case "additional_info":
			contents.Personal.AdditionalInfo = v.String()
		case "website":
			contents.Personal.Website = v.String()
		case "github":
			contents.Personal.GitHub = v.String()
		case "linkedin":
			contents.Personal.LinkedIn = v.String()
		}
	})
}

func parseEducationLuaTable(contents *ResumeContents, tbl *lua.LTable) {
	tbl.ForEach(func(_, value lua.LValue) {
		if itemTable, ok := value.(*lua.LTable); ok {
			var item EducationItem
			itemTable.ForEach(func(key, value lua.LValue) {
				switch key.String() {
				case "institution":
					item.Institution = value.String()
				case "location":
					item.Location = value.String()
				case "degree":
					item.Degree = value.String()
				case "dates":
					item.Dates = value.String()
				case "gpa":
					item.GPA = value.String()
				}
			})
			contents.Education.EducationItems = append(contents.Education.EducationItems, item)
		}
	})
}

func parseWorkLuaTable(contents *ResumeContents, tbl *lua.LTable) {
	tbl.ForEach(func(_, value lua.LValue) {
		if itemTable, ok := value.(*lua.LTable); ok {
			var item WorkItem
			itemTable.ForEach(func(key, value lua.LValue) {
				switch key.String() {
				case "job_title":
					item.JobTitle = value.String()
				case "company":
					item.Company = value.String()
				case "dates":
					item.Dates = value.String()
				case "location":
					item.Location = value.String()
				case "description":
					if descriptionTable, ok := value.(*lua.LTable); ok {
						item.Description = luaTableToStringSlice(descriptionTable)
					}
				}
			})
			contents.Work.WorkItems = append(contents.Work.WorkItems, item)
		}
	})
}

func parseProjectsLuaTable(contents *ResumeContents, tbl *lua.LTable) {
	tbl.ForEach(func(_, value lua.LValue) {
		if itemTable, ok := value.(*lua.LTable); ok {
			var item ProjectItem
			itemTable.ForEach(func(key, value lua.LValue) {
				switch key.String() {
				case "name":
					item.Name = value.String()
				case "link":
					item.Link = value.String()
				case "languages":
					if languagesTable, ok := value.(*lua.LTable); ok {
						item.Languages = luaTableToStringSlice(languagesTable)
					}
				case "description":
					if descriptionTable, ok := value.(*lua.LTable); ok {
						item.Description = luaTableToStringSlice(descriptionTable)
					}
				}
			})
			contents.Projects.ProjectItems = append(contents.Projects.ProjectItems, item)
		}
	})
}

func parseSkillsLuaTable(contents *ResumeContents, tbl *lua.LTable) {
	tbl.ForEach(func(key, value lua.LValue) {
		switch key.String() {
		case "languages":
			if languagesTable, ok := value.(*lua.LTable); ok {
				contents.Skills.Languages = luaTableToStringSlice(languagesTable)
			}
		case "libraries":
			if librariesTable, ok := value.(*lua.LTable); ok {
				contents.Skills.Libraries = luaTableToStringSlice(librariesTable)
			}
		case "databases":
			if databasesTable, ok := value.(*lua.LTable); ok {
				contents.Skills.Databases = luaTableToStringSlice(databasesTable)
			}
		case "tools":
			if toolsTable, ok := value.(*lua.LTable); ok {
				contents.Skills.Tools = luaTableToStringSlice(toolsTable)
			}
		}
	})
}

func luaTableToStringSlice(tbl *lua.LTable) []string {
	var result []string
	tbl.ForEach(func(_, value lua.LValue) {
		result = append(result, value.String())
	})
	return result
}
