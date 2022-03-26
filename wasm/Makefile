build:
	GOOS=js GOARCH=wasm go build -o bf.wasm
	cp $(shell go env GOROOT)/misc/wasm/wasm_exec.js .

run: build
	goexec 'http.ListenAndServe(`:8080`, http.FileServer(http.Dir(`.`)))'
