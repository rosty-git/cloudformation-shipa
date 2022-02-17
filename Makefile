.PHONY: build submit

build:
	cd app-cname && make build
	cd app-deploy && make build
	cd app-env && make build
	cd application && make build
	cd framework && make build
	cd network-policy && make build
	cd cluster && make build
	cd plan && make build
	cd job && make build

submit:
	cd app-cname && make submit
	cd app-deploy && make submit
	cd app-env && make submit
	cd application && make submit
	cd framework && make submit
	cd network-policy && make submit
	cd cluster && make submit
	cd plan && make submit
	cd job && make submit
