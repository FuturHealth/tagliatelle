TAG=us-docker.pkg.dev/futurhealth/cloud-run/tagliatelle:$(VERSION)

.PHONY: build
build:
	docker build --platform linux/amd64 -t $(TAG) .

.PHONY: push
push:
	docker push $(TAG)
