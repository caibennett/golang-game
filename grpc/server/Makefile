run:
	./gen.sh
	make -j 2 run/go run/proxy
run/go:
	go run .
run/proxy:
	grpcwebproxy --server_tls_cert_file=./cert/cert.crt --server_tls_key_file=./cert/cert.key --backend_addr=localhost:3000 --backend_tls_noverify