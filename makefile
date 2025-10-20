run:
	cd gateway && go run . start &
	cd auth-service && go run . start &
	wait