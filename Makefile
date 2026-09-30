p:
	mkdir -p bin
	go build -o ./bin/p-simple ./producer/simple
	./bin/p-simple "$(m)"

c:
	mkdir -p bin
	go build -o ./bin/c-simple ./consumer/01-simple
	./bin/c-simple

sp:
	mkdir -p bin
	go build -o ./bin/p-secure ./producer/secure
	./bin/p-secure "$(m)"

sc:
	mkdir -p bin
	go build -o ./bin/c-secure ./consumer/secure
	./bin/c-secure