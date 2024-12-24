package main

import (
	"github.com/upeshchalise/go_blogs/config"
	"github.com/upeshchalise/go_blogs/internal/database"
	"github.com/upeshchalise/go_blogs/pkg/logger"
)

func main() {
	config.Load()
	logger.Init()

	database.Init()
}
