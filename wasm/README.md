# WASM bindings for brainfuck compiler

## Usage
You should get the compiler in your path:
```bash
go get github.com/amirali/bf
```
then you should make the wasm file:
```bash
make
```
and then you have the `bf.wasm` file and load it into your project by using `wasm_exec.js` and import method in `index.html`. Then you have `execute_brainfuck("<program>", "<input_characters>")` and you can use it in your javascript code.


## Demo
You can demo this by serving on `:8080` port using this command:
```bash
make run
```
