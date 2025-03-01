package internal

import (
	"fmt"
	"strings"

	lua "github.com/yuin/gopher-lua"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func ParseLuaResumeContents(scriptPath string) (ResumeContents, []string, error) {
	var out ResumeContents
	var defaultTags []string

	L := lua.NewState()
	defer L.Close()

	if err := L.DoFile(scriptPath); err != nil {
		return out, nil, err
	}

	// FIX: ensure order of unnamed tables is always the same (order by dates where possible)
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
			case "default_tags":
				if tagsTbl, ok := value.(*lua.LTable); ok {
					defaultTags = parseDefaultTagsTable(tagsTbl)
				}
			case "extras":
				if extrasTbl, ok := value.(*lua.LTable); ok {
					parseExtraInfoTable(&out, extrasTbl)
				}
			}

		})
	}
	fmt.Println("parsed tags: ", defaultTags)

	return out, defaultTags, nil
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
			contents.Personal.Website = stringToURL(v.String())
		case "github":
			contents.Personal.GitHub = stringToURL(v.String())
		case "linkedin":
			contents.Personal.LinkedIn = stringToURL(v.String())
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
				case "honors":
					item.Honors = value.String()
				case "additional_info":
					if additionalInfoTbl, ok := value.(*lua.LTable); ok {
						item.AdditionalInfo = luaTableToStringSlice(additionalInfoTbl)
					}
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
			// TODO: programmatic switching of projects (add tag attr to struct?) (handle elsewhere)
			itemTable.ForEach(func(key, value lua.LValue) {
				switch key.String() {
				case "name":
					item.Name = value.String()
				case "link":
					// TODO: add https:// if not included in parsed link (also do this with personal links) (maybe use url lib?)
					item.Link = stringToURL(value.String())
				case "dates":
					item.Dates = value.String()
				case "tools":
					if toolsTable, ok := value.(*lua.LTable); ok {
						item.Tools = strings.Join(luaTableToStringSlice(toolsTable), ", ")
					}
				case "description":
					if descriptionTable, ok := value.(*lua.LTable); ok {
						item.Description = luaTableToStringSlice(descriptionTable)
					}
				case "tags":
					if tagsTable, ok := value.(*lua.LTable); ok {
						item.Tags = luaTableToStringSlice(tagsTable)
					}
				}
			})
			contents.Projects.ProjectItems = append(contents.Projects.ProjectItems, item)
		}
	})
}

func parseSkillsLuaTable(contents *ResumeContents, tbl *lua.LTable) {
	// TODO: Custom handling of sections, merging of string slices
	// - use subtables to filter by job?
	contents.Skills.Sections = make(map[string]skillValues)
	tbl.ForEach(func(key, value lua.LValue) {
		if sectionName := key.String(); sectionName != "" {
			if sectionTbl, ok := value.(*lua.LTable); ok {
				// TODO: filter out underscores, capitalize additional words
				name := cases.Title(language.English, cases.Compact).String(sectionName)
				values := luaTableToStringSlice(sectionTbl)
				contents.Skills.Sections[name] = skillValues{
					Values:       values,
					JoinedValues: strings.Join(values, ", "),
				}
			}
		}
	})
}

func parseDefaultTagsTable(tbl *lua.LTable) []string {
	var out []string
	tbl.ForEach(func(key, value lua.LValue) {
		fmt.Println(value)
		out = append(out, value.String())
	})
	return out
}

func parseExtraInfoTable(contents *ResumeContents, tbl *lua.LTable) {
	tbl.ForEach(func(k, v lua.LValue) {
		switch k.String() {
		case "visible":
			contents.Extras.Visible = v.String()
		case "hidden":
			contents.Extras.Hidden = v.String()
		}
	})
}
