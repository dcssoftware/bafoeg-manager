package assets

import "embed"

//go:embed terminal-assets/*
var TerminalFS embed.FS

//go:embed swagger-css/gopher-coffee.gif
var GopherCoffee string

//go:embed swagger-css/logo.png
var Logo string

//go:embed swagger-css/logo-branding.png
var LogoBranding string

//go:embed default-data/profilepicture.jpg
var DefaultProfilePicture []byte

//go:embed licenses/licenses.json
var SoftwareLicenses []byte
