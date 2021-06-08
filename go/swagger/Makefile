.PHONY: *

serve:
	go run .

swagger.json:
	./swagger generate spec -o ./swagger.json

doc: swagger.json
	./swagger serve -F=swagger swagger.json

doc/validate: swagger.json
	./swagger validate ./swagger.json
