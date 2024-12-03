package web

import v1 "love_knot/internal/app/api/handler/web/v1"

type V1 struct {
	Common *v1.Common
	User   *v1.User
	Friend *v1.Friend
}

type Handler struct {
	V1 *V1
}
