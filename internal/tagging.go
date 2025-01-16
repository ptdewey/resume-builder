package internal

import (
	"slices"
	"strings"
)

var maxProjects = 4
var projectCount = 0

func SelectTags(contents *ResumeContents, tags []string) {
	var temp ResumeContents

	for _, tag := range tags {
		selected := selectTaggedProjects(*contents, tag)
		for _, p := range selected {
			if !contains(temp.Projects.ProjectItems, ProjectItem{Name: p.Name}, "Name") {
				temp.Projects.ProjectItems = append(temp.Projects.ProjectItems, p)
			} else {
				projectCount--
			}
		}
		// TODO: also allow tagging of skills
	}

	contents.Projects.ProjectItems = temp.Projects.ProjectItems
}

func selectTaggedProjects(contents ResumeContents, tag string) []ProjectItem {
	var projects []ProjectItem

	for _, p := range contents.Projects.ProjectItems {
		if strings.ToLower(p.Name) == tag {
			projects = append(projects, p)
			projectCount++
			// TODO: remove a project item if at maxProjects and matched string is found
		} else if projectCount >= maxProjects {
			break
		} else if slices.Contains(p.Tags, tag) {
			projects = append(projects, p)
			projectCount++
		}
	}

	return projects
}

// TODO: skills are represented in a map
//
// func selectTaggedSkills(contents ResumeContents, tag string) []skillValues {
// 	var skills []skillValues
//
// 	for i, sv := range contents.Skills.
//
// 	return nil
// }
