local resume = {}

-- Personal Info
resume.personal = {
    name = "John Doe",
    phone = "800-123-4567",
    email = "johndoe@email.com",
    location = "Richmond, VA",
    website = "https://example.com",
    github = "https://github.com/johndoe",
    linkedin = "https://linkedin.com/in/johndoe/",
    additional_info = "",
}

-- Education
resume.education = {
    -- Sub-table names are optional for the education table
    masters = {
        institution = "Virginia Tech",
        location = "Blacksburg, VA",
        degree = "Masters of Computer Science",
        dates = "August 2023 - May 2025",
        gpa = "4.0",
    },
}

-- Work Experience
resume.work = {
    -- Sub-table names are optional for the work table
    company2 = {
        job_title = "Senior Software Engineer",
        company = "Company Two",
        dates = "May 2023 - Present",
        location = "Richmond, VA",
        description = {
            "Description here",
            "Additional Info",
        },
    },

    {
        job_title = "Junior Software Engineer",
        company = "Company One",
        dates = "May 2020 - August 2022",
        location = "Richmond, VA",
        description = {
            "Built stuff",
        },
    },
}

-- Projects
resume.projects = {
    -- Sub-table names are optional for the projects table
    {
        name = "Resume Builder",
        link = "https://github.com/ptdewey/resume-builder",
        tools = { "Go", "Typst", "Lua" },
        description = {
            "Easy to use resume builder that builds your resume from a Lua configuration file consisting of a small number of tables",
        },
    },
}

-- Skills
resume.skills = {
    -- Sub-table names within skills section require names

    -- Programming languages
    languages = {
        "Go",
        "Python",
        "Lua",
        "JavaScript",
        "Typst",
    },

    -- Libraries
    libraries = {
        "Pandas",
        "SciKit Learn",
        "PyTorch",
        "React",
    },

    -- Technologies and Tools
    technologies = {
        "Docker",
        "Linux",
        "Git",
        "SQLite",
        "MongoDB",
    },
}

return resume
