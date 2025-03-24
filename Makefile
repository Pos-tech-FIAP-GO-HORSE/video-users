generate-mocks:
	mockery --dir src/core/useCases/ --all --output src/mocks
	mockery --dir src/repositories/ --all --output src/mocks
