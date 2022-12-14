
# makefile 详细文档：https://seisman.github.io/how-to-write-makefile/invoke.html

.DEFAULT_GOAL := help

cmd := $(MAKECMDGOALS)


.PHONY: controller
controller:
ifeq ($(cmd),controller)
	@go run main.go make $(cmd) -h
else
	@go run main.go make $(cmd)
endif

.PHONY: help
help:
	@go run main.go


.PHONY: model
model:
ifeq ($(cmd),model)
	@go run main.go make $(cmd) -h
else
	@go run main.go make $(cmd)
endif

.PHONY: policy
policy:
ifeq ($(cmd),policy)
	@go run main.go make $(cmd) -h
else
	@go run main.go make $(cmd)
endif



.PHONY: request
request:
ifeq ($(cmd),request)
	@go run main.go make $(cmd) -h
else
	@go run main.go make $(cmd)
endif


.PHONY: seed
seed:
ifeq ($(cmd),seed)
	@go run main.go make $(cmd) -h
else
	@go run main.go make $(cmd)
endif


.PHONY: factory
factory:
ifeq ($(cmd),factory)
	@go run main.go make $(cmd) -h
else
	@go run main.go make $(cmd)
endif