#show heading: set text(font: "Libertinus Serif")

#show link: underline

#set page(
  margin: (x: 0.9cm, y: 1.3cm)
)

#set par(justify: true)

#let chiline() = {v(-3pt); line(length: 100%); v(-5pt)}

// Header
#text(22pt)[*{{ .Personal.Name }}*] #h(1fr) Location: {{ .Personal.Location }} {{ .Personal.AdditionalInfo }}

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

// TODO:

== Work Experience
#chiline()

// TODO:

== Skills
#chiline()

// TODO:

== Projects
#chiline()

// TODO:
