
test:
	# run tests
	export consolia_env=test; godep go test ./...

testschemas:
	# Run the server on a dedicated test port
	export consolia_env=test; export consolia_port=9001; ./consolia-api test &

	# run the json schema tests
	export consolia_env=test; export consolia_port=9001; go run schemas/test_schemas.go ihqy1nguQm6IWM5P7rI59g

	# kill the server
	ps aux | grep "consolia-api test" | awk '{print $$2}' | xargs kill

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
