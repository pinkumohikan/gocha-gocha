.PHONY: summary
summary:
	@./fetch-all.sh
	find ./proposal-pages -type f | xargs -n 1 ./parse.sh 2>/dev/null | sort | uniq -c
