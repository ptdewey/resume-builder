#show heading: set text(font: "Libertinus Serif")

#show link: underline

#set page(
  margin: (x: 0.9cm, y: 1.3cm)
)

#set par(justify: true)

#let chiline() = {v(-3pt); line(length: 100%); v(-5pt)}

#let dual_grid(left_body, right_body, top_pad: 0cm, bottom_pad: 0cm) = {
  pad(top: top_pad, bottom: bottom_pad,
    grid(
      columns: (1fr, 1fr),
      gutter: auto,
      align(left)[#left_body],
      align(right)[#right_body],
    )
  )
}

#let tri_grid(left_body, center_body, right_body, columns: (1fr, 1.5fr, 1fr)) = {
  grid(
    columns: columns,
    gutter: auto,
    align(left)[#left_body],
    align(center)[#center_body],
    align(right)[#right_body],
  )
}

#let quad_grid(upper_left_body, upper_right_body, lower_left_body, lower_right_body, upper_top_pad: 0cm, upper_bottom_pad: 0cm, lower_top_pad: 0cm, lower_bottom_pad: 0cm) = {
  pad(top: upper_top_pad, bottom: upper_bottom_pad,
    grid(
      columns: (3fr, 1fr),
      gutter: auto,
      align(left)[#upper_left_body],
      align(right)[#upper_right_body],
    )
  )
  pad(top: lower_top_pad, bottom: lower_bottom_pad,
    grid(
      columns: (3fr, 1fr),
      gutter: auto,
      align(left)[#lower_left_body],
      align(right)[#lower_right_body],
    )
  )
}


// Header
#text(22pt)[*{{ .Personal.Name }}*] #h(1fr) Location: {{ .Personal.Location }}{{ .Personal.AdditionalInfo }}

// Socials and contact info
#dual_grid(top_pad: -0.20cm,
  [
    #link("{{ .Personal.Website }}", "Portfolio") |
    #link("{{ .Personal.LinkedIn }}", "LinkedIn") |
    #link("{{ .Personal.GitHub }}", "GitHub")
  ],
  [
    Email: #link("mailto:{{ .Personal.Email }}", "{{ .Personal.Email }}") |
    Phone: {{ .Personal.Phone }}
  ]
)


== Education
#chiline()
{{ range $item := .Education.EducationItems }}
#text(10pt)[
  // education item header
  #quad_grid(
    lower_top_pad: -0.2cm,
    lower_bottom_pad: -0.2cm,
    [*{{ $item.Institution }}*],
    [{{ $item.Location }}],
    [_{{ $item.Degree }}_],
    [_{{ $item.Dates }}_],
  )
  // education item details
  #list(
    [*Honors:* {{ $item.Honors }}], {{ range $bullet := $item.AdditionalInfo }}
    [{{ $bullet }}], {{ end }}
  )
]
{{ end }}

== Work Experience
#chiline()
{{ range $item := .Work.WorkItems }}
#text(10pt)[
  #quad_grid(
    lower_top_pad: -0.2cm,
    lower_bottom_pad: -0.2cm,
    // TODO: maybe separate to quad layout
    [*{{ $item.Company }}*],
    [{{ $item.Location }}],
    [_{{ $item.JobTitle }}_],
    [_{{ $item.Dates }}_]
  )
  #list({{ range $bullet := $item.Description }}
    [{{ $bullet }}], {{ end }}
  )
]
{{ end }}

== Skills
#chiline()

// TODO: allow different table names here for dynamic changes
#text(10pt)[
  #list({{ range $section, $values := .Skills.Sections }}
    [*{{ $section }}:* {{ $values.JoinedValues }}],{{end}} // REFACTOR: maybe change this to not a list
  )
]


== Projects
#chiline()
{{ range $item := .Projects.ProjectItems }}
#text(10pt)[
  #dual_grid(
    link("{{ $item.Link }}", [*{{ $item.Name }}*]),
    [{{ $item.Tools }}],
  )
  #pad(top: -0.2cm, list({{ range $bullet := $item.Description }}
    [{{ $bullet }}], {{ end }}
  ))
]
{{ end }}
