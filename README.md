# Jetpack game

![Jetpack game banner](https://github.com/darellanodev/jetpack-game/blob/main/img_github_readme/banner.png?raw=true)

A simple game to learn Go and Ebitengine.

## THIS APPLICATION IS UNDER ACTIVE DEVELOPMENT, BUT STILL CONSIDERED BETA

## Github repository

- <https://github.com/darellanodev/jetpack-game>

## Technologies

[![Go](https://img.shields.io/badge/Go-00ADD8?style=flat&logo=Go&logoColor=white)](https://golang.org)
[![Ebitengine](https://img.shields.io/badge/Ebitengine-005f73?style=flat&logo=Go&logoColor=white)](https://ebitengine.org)

## Screenshots

![jet pack game screenshots](https://github.com/darellanodev/jetpack-game/blob/main/img_github_readme/screenshots.png?raw=true)

## Build for web

Execute with `./run_build_web.sh` It would generate `jetpackgame.wasm` and `wasm_exec.js` inside build_web folder

## Play in web browser

After building for web, use a web server like XAMPP or Live Server extension in VSCode `./build_web/index.html`

### Notes for web browser

- All resources are embedded into the wasm file. The browser must wait to load the entire wasm file. For that, the game can not have a preloading inside wasm.

## Execute locally

Execute `run.sh` or `go run .` command

## Execute the unit tests

Execute `run_tests.sh` or `go test` command

## How to contribute

Check out the contribution guidelines [here](./CONTRIBUTING.md).

## Discord channel

Join the discord channel: [![Discord](https://img.shields.io/badge/Discord-%235865F2.svg?logo=discord&logoColor=white)](https://discord.gg/YRjvggs6)

## Customize keybindings.json in VSCode

You can use this settings into VSCode `keybindings.json`:

```json
  {
    "key": "ctrl+t",
    "command": "workbench.action.terminal.sendSequence",
    "args": {
      "text": "./run_tests.sh\u000D"
    },
  },
  {
    "key": "ctrl+r",
    "command": "workbench.action.terminal.sendSequence",
    "args": {
      "text": "./run.sh\u000D"
    },
  },
```
