caddy_local_repo = 0x1c.caddy
caddy_tag = $(tag)
caddy_pub_repo = registry.gitlab.com/ox1c/caddy

.PHONY: hello
## Say 'Hello' to everybody
hello:
	@echo "Hello Arse!"

.PHONY: show.time
## Show date & time of your local machine
show.time:
	@echo "$$(date)"

.PHONY: show.docker-ps
## List the docker containers that runs on your local machine.
## Note:
## This action requires 'docker' installed as prerequisite.
show.docker-ps:
	@docker ps
	@echo
	@echo "And we are done here.\n\
			- Nobody\n"

.PHONY: show.docker-hello_internal
show.docker-hello_internal: ## This is an internal target
	@docker run --rm \
		hello-world

.PHONY: hello.docker
## Run the famous 'hello-world' docker image and then check the running docker container, where you'll hopefully see the one that runs the 'hello-world' image
hello.docker: show.docker-hello_internal show.docker-ps