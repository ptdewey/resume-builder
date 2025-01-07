#show heading: set text(font: "Libertinus Serif")

#show link: underline

#set page(
  margin: (x: 0.9cm, y: 1.3cm)
)

#set par(justify: true)

#let chiline() = {v(-3pt); line(length: 100%); v(-5pt)}

// Header
#text(22pt)[*{{ .Personal.Name }}*] #h(1fr) Location: {{ .Personal.Location }}{{ .Personal.AdditionalInfo }}

// Socials and contact info
#grid(
  columns: (1fr, 1fr),
  align(left, pad(top: -0.20cm)[
    #link("{{ .Personal.Website }}", "Portfolio") |
    #link("{{ .Personal.LinkedIn }}", "LinkedIn") |
    #link("{{ .Personal.GitHub }}", "GitHub")
  ]),
  align(right, pad(top: -0.20cm)[
    Email: #link("mailto:{{ .Personal.Email }}", "{{ .Personal.Email }}") |
    Phone: {{ .Personal.Phone }}
  ])
)


== Education
#chiline()

{{ range $item := .Education.EducationItems }}
#text(10pt)[
  // work item header
  #grid(
    columns: (1fr, 1.5fr, 1fr), // NOTE: adjust column fractions as necessary
    gutter: auto,
    align(left)[*{{ $item.Institution }}*],
    align(center)[_{{ $item.Location }}_],
    align(right)[{{ $item.Dates }}],
  )
  // work item details (and location?)
  #pad(top: -0.25cm, list(
    [{{ $item.Degree }} #h(1fr) GPA: {{ $item.GPA }}],
  ))
]
{{ end }}

== Work Experience
#chiline()

{{ range $item := .Work.WorkItems }}
#text(10pt)[
  // work item header
  #grid(
    columns: (1fr, 1.5fr, 1fr), // NOTE: adjust column fractions as necessary
    gutter: auto,
    align(left)[*{{ $item.JobTitle }}*],
    align(center)[_{{ $item.Company }}_],
    align(right)[{{ $item.Dates }}],
  )
  // work item details (and location?)
  // TODO: improve spacing between bullet points (and headk)
  #list(
    {{ range $bullet := $item.Description }}
    [{{ $bullet }}],
    {{ end }}
  )
]
{{ end }}
// TODO:


// TODO:


== Skills
#chiline()

// TODO:


== Projects
#chiline()

// TODO:
