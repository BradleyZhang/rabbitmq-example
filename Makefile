p:
	@mkdir -p bin
	@go build -o ./bin/p-simple ./producer/01-simple
	@./bin/p-simple "$(m)"

c:
	@mkdir -p bin
	@go build -o ./bin/c-simple ./consumer/01-simple
	@./bin/c-simple

sp:
	@mkdir -p bin
	@go build -o ./bin/p-secure ./producer/02-secure
	@./bin/p-secure "$(m)"

sc:
	@mkdir -p bin
	@go build -o ./bin/c-secure ./consumer/02-secure
	@./bin/c-secure

wp:
	@mkdir -p bin
	@go build -o ./bin/p-task ./producer/03-work-queue
	@./bin/p-task "$(m)"

wc:
	@mkdir -p bin
	@go build -o ./bin/c-w ./consumer/03-work-queue
	@./bin/c-w

4p:
	@mkdir -p bin
	@go build -o ./bin/p-4 ./producer/04-publish
	@./bin/p-4

4c:
	@mkdir -p bin
	@go build -o ./bin/c-4 ./consumer/04-subscribe
	@./bin/c-4

5p:
	@mkdir -p bin
	@go build -o ./bin/p-5 ./producer/05-routing
	@./bin/p-5 $(key)

5c:
	@mkdir -p bin
	@go build -o ./bin/c-5 ./consumer/05-routing
	@./bin/c-5 $(key)