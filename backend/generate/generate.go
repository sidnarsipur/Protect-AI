package generate

import "github.com/sidnarsipur/protect-ai/models"

func SetPermissions(permissions []models.Permission) string {
	var permissionString string
	for _, permission := range permissions {
		permissionString += setAgentHeader(permission.Agent)
		permissionString += allowFileTypes(permission.AllowedPermission.FileTypes)
		permissionString += disallowFileTypes(permission.DisallowedPermission.FileTypes)
		permissionString += allowPaths(permission.AllowedPermission.Paths)
		permissionString += disallowPaths(permission.DisallowedPermission.Paths)
	}
	return permissionString
}

func setAgentHeader(agent string) string {
	return "User-Agent: " + agent + "\n"
}

func allowPaths(paths []string) string {
	var allowString string

	for _, path := range paths {
		allowString += "Allow: " + path + "\n"
	}

	return allowString
}

func disallowPaths(paths []string) string {
	var disallowString string

	for _, path := range paths {
		disallowString += "Disallow: " + path + "\n"
	}

	return disallowString
}

func allowAll() string {
	return "Allow: /\n"
}

func disallowAll() string {
	return "Disallow: /\n"
}

func allowFileTypes(fileTypes models.FileTypes) string {
	var disallowString string

	for _, fileType := range fileTypes.TextFileTypes {
		disallowString += "Allow: *." + string(fileType) + "\n"
	}

	for _, fileType := range fileTypes.ImageFileTypes {
		disallowString += "Allow: *." + string(fileType) + "\n"
	}

	for _, fileType := range fileTypes.AudioFileTypes {
		disallowString += "Allow: *." + string(fileType) + "\n"
	}

	for _, fileType := range fileTypes.VideoFileTypes {
		disallowString += "Allow: *." + string(fileType) + "\n"
	}

	return disallowString
}

func disallowFileTypes(fileTypes models.FileTypes) string {
	var disallowString string

	for _, fileType := range fileTypes.TextFileTypes {
		disallowString += "Disallow: *." + string(fileType) + "\n"
	}

	for _, fileType := range fileTypes.ImageFileTypes {
		disallowString += "Disallow: *." + string(fileType) + "\n"
	}

	for _, fileType := range fileTypes.AudioFileTypes {
		disallowString += "Disallow: *." + string(fileType) + "\n"
	}

	for _, fileType := range fileTypes.VideoFileTypes {
		disallowString += "Disallow: *." + string(fileType) + "\n"
	}

	return disallowString
}
