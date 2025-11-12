run:
	cd gateway && go run . start &
	cd auth-service && go run . start &
	cd umkm-service && go run . start &
	wait

stop:
	pkill gateway
	pkill auth-service
	pkill umkm-service