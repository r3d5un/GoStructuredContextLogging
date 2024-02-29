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
}
