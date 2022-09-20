package modules

type Linux struct {
	Network Network `json:"network"`
	Sys     Sys     `json:"sys"`
}
