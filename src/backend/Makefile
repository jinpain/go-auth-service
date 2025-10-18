.DEFAULT_GOAL = local

ROOT_DIR=$(shell pwd)

local:
	CONFIG_PATH=$(ROOT_DIR)/config/config.yaml go run $(ROOT_DIR)/cmd/main.go
prod:
	CONFIG_PATH=$(ROOT_DIR)/config/config.yaml go run $(ROOT_DIR)/cmd/main.go