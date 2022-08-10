ifeq (true, $(shell docker compose version > /dev/null 2>&1 && echo true))
	DOCKER_COMPOSE := docker compose
else ifeq (true, $(shell docker-compose version > /dev/null 2>&1 && echo true))
	DOCKER_COMPOSE := docker-compose
else
	ERR := $(error Docker compose is not installed. Refer to the documentation for instructions: https://docs.docker.com/compose/install/compose-plugin)
endif

ifneq ($(filter devsh devshroot,$(firstword $(MAKECMDGOALS))),)
	CONTAINER := $(firstword $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS)))
	EVAL := $(eval $(CONTAINER):;@:)
endif

devup:
	USER=$$(id -u):$$(id -g) docker compose up -d --remove-orphans

devdown:
	docker compose down --remove-orphans

devshroot:
	@$(DOCKER_COMPOSE) exec -it -u 0:0 $(CONTAINER) sh -c "command -v bash >/dev/null && exec bash || exec sh"

devsh:
	@user="-u $$(id -u):$$(id -g)";$(DOCKER_COMPOSE) exec -it $$user $(CONTAINER) sh -c "command -v bash >/dev/null && exec bash || exec sh"

build:
	CGO_ENABLED=0 go build -o .out/wgpm -ldflags "-s -w"
