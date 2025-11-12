run:
	cd gateway && go run . start &
	cd auth-service && go run . start &
	cd umkm-service && go run . start &
	wait

stop:
	pkill gateway || true
	pkill auth-service || true
	pkill umkm-service || true

build:
	mkdir -p bin
	cd gateway && go build -o ../bin/gateway .
	cd auth-service && go build -o ../bin/auth-service .
	cd umkm-service && go build -o ../bin/umkm-service .

.PHONY: run stop build