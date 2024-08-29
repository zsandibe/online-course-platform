package main

import (
	"github.com/zsandibe/online-course-platform/internal/app"
	logger "github.com/zsandibe/online-course-platform/pkg"
)

func main() {
	if err := app.Start(); err != nil {
		logger.Error(err)
		return
	}
}
