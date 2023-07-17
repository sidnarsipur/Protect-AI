package models

type Permission struct {
	Agent                string
	AllowedPermission    PermissionBody
	DisallowedPermission PermissionBody
}

type PermissionBody struct {
	FileTypes FileTypes
	Paths     []string
}
