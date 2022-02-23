.PHONY: setup dry-run apply clean
setup: ./bin/mysqldef

dry-run:
	./bin/mysqldef --host 127.0.0.1 --user root --port 33060 exp1 --dry-run < ./schema.sql

apply:
	./bin/mysqldef --host 127.0.0.1 --user root --port 33060 exp1 < ./schema.sql

clean:
	rm -f ./bin/mysqldef

# no .PHONY
./bin/mysqldef: version=0.11.38
./bin/mysqldef:
	curl -sSfL -o /tmp/mysqldef_darwin_amd64.zip https://github.com/k0kubun/sqldef/releases/download/v$(version)/mysqldef_darwin_amd64.zip
	unzip /tmp/mysqldef_darwin_amd64.zip
	mv mysqldef ./bin

