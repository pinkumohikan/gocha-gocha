
default: setup
.PHONY: setup

setup: composer.phar
	./composer.phar install --no-dev --optimize-autoloader --no-interaction

composer.phar:
	./download-composer.sh
