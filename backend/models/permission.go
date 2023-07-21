package models

type Permission struct {
	Agent                string
	AllowedPermission    PermissionBody
	DisallowedPermission PermissionBody
	CrawlDelay           int
	AllowAll             bool
	DisallowAll          bool
}

type PermissionBody struct {
	FileTypes FileTypes
	Paths     []string
}
