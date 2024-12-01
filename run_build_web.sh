env GOOS=js GOARCH=wasm go build -o ./build_web/jetpackgame.wasm github.com/darellanodev/jetpack-game
cp $(go env GOROOT)/misc/wasm/wasm_exec.js ./build_web