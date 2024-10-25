startdb:
	@docker run -d --rm \
			--volume "./resources/pg/init/init.sql:/docker-entrypoint-initdb.d/init.sql" \
			--volume "./resources/pg/data:/usr/pgdata" \
			--name=forecastDB \
			-p 5544:5432 \
			-e POSTGRES_DB=forecast \
			-e POSTGRES_USER=forecast \
			-e POSTGRES_PASSWORD=forecast postgres:16-alpine


stopdb:
	@docker stop forecastDB || exit 0
test:
	cd api/server && go test -v
	cd pkg/api && go test -v
	cd pkg/output && go test -v
	
