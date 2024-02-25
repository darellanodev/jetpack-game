# go-jetpack-game

![jetpack game banner](https://github.com/darellanodev/jetpack-game/blob/main/img_github_readme/banner.png?raw=true)

A simple game to learn Go and Ebitengine *(THIS APPLICATION IS IN AN EARLY STAGE OF DEVELOPMENT)*

![jetpack game screenshots](https://github.com/darellanodev/jetpack-game/blob/main/img_github_readme/screenshots.png?raw=true)

## Build for web

Execute with `./build_web.sh` It would generate `jetpackgame.wasm` and `wasm_exec.js` inside build_web folder

## Play in web browser

After building for web, use a web server like XAMPP or Live Server extension in VSCode `./build_web/index.html`

### Notes for web browser

- All resources are embedded into the wasm file. The browser must wait to load the entire wasm file. For that, the game can not have a preloading inside wasm.

## Execute locally

Execute `run.sh` or `go run .` command

## Execute the unit tests

Execute `run_tests.sh` or `go test` command

### Next tasks

- [x] Make pillars to the platforms
- [x] When gameover/gamecomplete the player can press enter to restart
- [x] Create a starting menu. The player must press ENTER key to start the game
- [x] Make a preloading screen because now in web version it shows a blank screen
- [x] Make a progress bar class more flexible (for example to allow set the width property) to reuse it in oxygen bar and preload game
- [x] Rotate the planets
- [x] Make vulcan floor
- [x] Show a screen traveling to the planet of fire
- [x] Some Lava Floor drops lava
- [ ] Burn animation of the player when touches fire (lava floor, lavadrop)
- [ ] The rocket scales down when it is aproaching to the planet
- [ ] Some fire floors are active and shows blinking light and then drops fireballs
- [ ] When typewritting make sounds
- [ ] Circle rotating and line below the typewritting text like a futuristic hud
- [ ] Second scene before landing. Show the objectives in this screen while rocket is descending
- [ ] Try to make a preloading system in html. It can not be done in game because the resouces are embedded in wasm file
- [ ] Trees in foreground/background
- [ ] Show a screen traveling to the planet of ice
- [ ] Ice floors The player slips on the ground
- [ ] Make procedural backgrounds (mountains, stars)
- [ ] More Unit testing and error handling
- [ ] In web version auto click the canvas to allow using keys without clicking first the canbas with mouse (view ebitengine web examples)
- [ ] Parachute with live (core icon) The icon blinks and after a time the live is removed
- [ ] When last level show congratulations game completed
- [ ] Make the collision box smaller to not collision so early
- [ ] Flip an image and use only one image instead two (left and right)
- [ ] Platform with a lake of lava that drops lava
- [ ] Other enemy that moves vertically an shots fire balls horizontally
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
