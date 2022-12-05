test:
	@go test -v ./...


push:
	@git add -A && git commit -m "update" && git push origin master


build_test:
	@go test -c -race -timeout 1000s -run ^TestRelimit$ github.com/realjf/relimit

run_test:
	@sudo ./relimit.test

# make tag t=<your_version>
tag:
	@echo '${t}'
	@git tag -a ${t} -m "${t}" && git push origin ${t}

dtag:
	@echo 'delete ${t}'
	@git push --delete origin ${t} && git tag -d ${t}

.PHONY: test push
