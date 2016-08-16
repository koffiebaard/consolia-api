
test:
	# run tests
	export consolia_env=test; godep go test ./...

testschemas:
	# Run the server on a dedicated test port
	export consolia_env=test; export consolia_port=9001; ./consolia-api test &

	# run the json schema tests
	export consolia_env=test; export consolia_port=9001; go run schemas/test_schemas.go NhBEsuAFQe-9Ly_vRjAxQw

	# kill the server
	ps aux | grep "consolia-api test" | awk '{print $$2}' | xargs kill

migrate_up:
	go run migrations/migrate.go up

migrate_down:
	go run migrations/migrate.go down

build:
	# save dependencies
	godep save

	# run tests
	godep go test $(go list ./... | grep -v vendor)

	# build
	godep go build

run:
	# save dependencies
	godep save

	# run tests
	godep go test $(go list ./... | grep -v vendor)

	# build
	godep go build

	# run
	./consolia-api

build_docker:
	docker build -t "consolia-api" .

run_docker:
	docker rm -f consolia-api;

	docker run  -e "consolia_db_port=$$consolia_db_port" \
				-e "consolia_db_name=$$consolia_db_name" \
				-e "consolia_db_username=$$consolia_db_username" \
				-e "consolia_db_password=$$consolia_db_password" \
				-e "consolia_env=$$consolia_env" \
			   -d --name consolia-api -p 3000:3003 consolia-api
