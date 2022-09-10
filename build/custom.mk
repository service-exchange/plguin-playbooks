# Include custom targets and environment variables here
GO_BUILD_FLAGS += -ldflags "-X main.rudderDataplaneURL=$(MM_RUDDER_DATAPLANE_URL) -X main.rudderWriteKey=$(MM_RUDDER_WRITE_KEY)"

## Generate mocks.
mocks:
ifneq ($(HAS_SERVER),)
	mockgen -destination server/bot/mocks/mock_poster.go github.com/service-exchange/service-exchange-plugin-playbooks/server/bot Poster
	mockgen -destination server/app/mocks/mock_job_once_scheduler.go github.com/service-exchange/service-exchange-plugin-playbooks/server/app JobOnceScheduler
	mockgen -destination server/sqlstore/mocks/mock_kvapi.go github.com/service-exchange/service-exchange-plugin-playbooks/server/sqlstore KVAPI
	mockgen -destination server/sqlstore/mocks/mock_storeapi.go github.com/service-exchange/service-exchange-plugin-playbooks/server/sqlstore StoreAPI
	mockgen -destination server/sqlstore/mocks/mock_configurationapi.go github.com/service-exchange/service-exchange-plugin-playbooks/server/sqlstore ConfigurationAPI
endif

## Runs the redocly server.
.PHONY: docs-server
docs-server:
	npx @redocly/openapi-cli@1.0.0-beta.3 preview-docs server/api/api.yaml

## Re-generate tests-e2e/db-setup/service-exchange.sql from the Postgres image expected to be running
## in the developer's Docker environment.
.PHONY: tests-e2e/db-setup/service-exchange.sql
tests-e2e/db-setup/service-exchange.sql:
	docker exec service-exchange-postgres pg_dump \
		--username=mmuser \
		--clean \
		--if-exists \
		--exclude-table ir_incident \
		--exclude-table ir_playbook \
		--exclude-table ir_playbookmember \
		--exclude-table ir_statusposts \
		--exclude-table ir_system \
		--exclude-table ir_timelineevent \
		--exclude-table ir_userinfo \
		--exclude-table ir_viewedchannel \
		service-exchange_test > tests-e2e/db-setup/service-exchange.sql
