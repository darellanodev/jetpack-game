# go-jetpack-game

![jetpack game banner](https://github.com/darellanodev/jetpack-game/blob/main/img_github_readme/banner.png?raw=true)

A simple game to learn Go and Ebitengine *(THIS APPLICATION IS IN AN EARLY STAGE OF DEVELOPMENT)*

![jetpack game screenshots](https://github.com/darellanodev/jetpack-game/blob/main/img_github_readme/screenshots.png?raw=true)

## Build for web

Execute with `./build_web.sh` It would generate `jetpackgame.wasm` and `wasm_exec.js` inside build_web folder

## Play in web browser

After building for web, use a web server like XAMPP or live server extension in VSCode `./build_web/index.html`

## Execute locally

Execute `run.sh` or `go run .` command

## Execute the unit tests

Execute `run_tests.sh` or `go test` command

### Next tasks

- [ ] Make a progress bar class more flexible (for example to allow set the width property) to reuse it in oxygen bar and preload game
- [ ] Make pillars to the platforms.
- [ ] Make procedural backgrounds (mountains, stars)
- [ ] Make a preloading screen because now in web version it shows a blank screen.
- [ ] Make procedurals trees, peaceful animals in foreground/background
- [ ] More Unit testing and error handling.
- [ ] When game over the player can press spacebar or enter to start.
- [ ] Create a basic starting menu. The player must press the space key to start the game.
- [ ] In web version auto click the canvas to allow using keys without clicking first the canbas with mouse (view ebitengine web examples)
- [ ] Ice floors. The player slips on the ground.
- [ ] Parachute with live (core icon). The icon blinks and after a time the live is removed.
- [ ] When last level show congratulations game completed.
- [ ] Make the collision box smaller to not collision so early
- [ ] Flip an image and use only one image instead two (left and right)
- [ ] Platform with a lake of lava that drops lava.
- [ ] Other enemy that moves vertically an shots fire balls horizontally.
- [ ] Level with a vulcan that throws fire balls.
- [ ] Level editor

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
