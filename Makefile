test:
	@go test -v ./...


push:
	@git add -A && git commit -m "update" && git push origin master


build_test:
	@go test -c -race -timeout 1000s -run ^TestRelimit$ github.com/realjf/relimit

run_test:
	@sudo ./relimit.test

.PHONY: test push
