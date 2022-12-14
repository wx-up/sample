
# makefile 详细文档：https://seisman.github.io/how-to-write-makefile/invoke.html

.DEFAULT_GOAL := help

.PHONY: controller
controller:
	@go run main.go make $(MAKECMDGOALS)

.PHONY: help
help:
	@go run main.go


.PHONY: model
model:
	@go run main.go make $(MAKECMDGOALS)

.PHONY: policy
policy:
	@go run main.go make $(MAKECMDGOALS)


.PHONY: request
request:
	@go run main.go make $(MAKECMDGOALS)


.PHONY: seed
seed:
	@go run main.go make $(MAKECMDGOALS)


.PHONY: factory
factory:
	@go run main.go make $(MAKECMDGOALS)