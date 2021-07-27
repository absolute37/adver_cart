package core

import "github.com/gofiber/fiber/v2"

type Context struct {
	*fiber.Ctx
}

func NewFiberContext(ctx *fiber.Ctx) *Context {
	return &Context{ctx}
}
