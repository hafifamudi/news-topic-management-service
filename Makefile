.PHONY: dep run run-test-in-windows run-test-in-linux

test_dir := github.com/hafifamudi/news-topic-management-service/internal/core/news/controller \
        github.com/hafifamudi/news-topic-management-service/internal/core/news/repository \
        github.com/hafifamudi/news-topic-management-service/internal/core/news/service \
        github.com/hafifamudi/news-topic-management-service/internal/core/topic/controller \
        github.com/hafifamudi/news-topic-management-service/internal/core/topic/repository \
        github.com/hafifamudi/news-topic-management-service/internal/core/topic/service

dep:
	@echo ">> Downloading Dependencies"
	@go mod tidy

run: dep
	@air

seed:
	@go run ./cmd/seed/main.go

run-test: dep
	@go test -v $(test_dir)
