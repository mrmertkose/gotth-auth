package views

import "context"

func IsAuth(ctx context.Context) bool {
	isAuthed := ctx.Value("isAuth")
	if isAuthed == false {
		return false
	}
	return isAuthed.(bool)
}

func GetName(ctx context.Context) string {
	name := ctx.Value("name")
	if name == "" {
		return ""
	}
	return name.(string)
}
