# go vet
vet:
	$(info #Running go vet...)
	go vet -c=0 ./...

.PHONY: vet