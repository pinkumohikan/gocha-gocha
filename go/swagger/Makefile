.PHONY: *

serve:
	go run .

swagger.json:
	./swagger generate spec -o ./swagger.json

doc: swagger.json
	./swagger serve -F=swagger swagger.json

doc/validate: swagger.json
	./swagger validate ./swagger.json

client/gen: swagger.json
	docker run --rm -v ${PWD}:/local swaggerapi/swagger-codegen-cli generate     -i /local/swagger.json -l typescript-fetch -o /local/out/go
