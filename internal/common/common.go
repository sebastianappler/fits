package common

import "net/url"

type Path struct {
	Url      url.URL
	UrlRaw   string
	Username string
	Password string
}
