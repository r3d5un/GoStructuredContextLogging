package main

import (
	"context"
	"log/slog"
	"os"
)

func main() {
	handler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(handler)
	slog.SetDefault(logger)

	// simple log statement
	slog.Info("Hello, World!")

	// log statement with extra attributes
	slog.Info("extra attributes", "msg", "possibilities limitless")

	// stronly typed attributes
	slog.Info("strongly typed attributes", slog.String("msg", "possibilities limited"))

	// guarantee strong types
	slog.LogAttrs(
		context.Background(), slog.LevelInfo,
		"truly, the best choice is no choice",
		slog.String("msg", "Do you have more microservices than customers?"),
		slog.Bool("answer", true),
	)

	// grouping attributes
	slog.LogAttrs(
		context.Background(), slog.LevelInfo,
		"these are not the droids you're looking for",
		slog.Group(
			"droids",
			slog.Int("droid_id", 1),
			slog.Int("droid_id", 2),
		),
	)

	// child loggers
	childLogger := logger.With(
		slog.Group(
			"inside",
			slog.String("thing", "oh, yeah!"),
		),
	)

	// child loggers include properties from it's declaration
	childLogger.Info("omg, somethings inside me!")

	// this is useful of embedding and grouping contextual data
	// without having to write repetitive log statements
	childLogger = logger.With(
		slog.Group(
			"request",
			"method", "GET",
			"path", "/totally/a/real/path",
			"request_id", 1234,
		),
	)
	childLogger.Info("request context embedded")

	// embed all attributes as part of a group
	loggerWithGroup := slog.New(handler).WithGroup("some_group")
	child := loggerWithGroup.With(
		slog.Int("some_number", 9876),
		slog.Int("another_number", 5432),
	)

	child.Info("some_statement")
}
