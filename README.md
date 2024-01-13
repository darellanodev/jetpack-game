# go-jetpack-game

A simple game to learn Go and Ebitengine *(THIS APPLICATION IS IN AN EARLY STAGE OF DEVELOPMENT)*

## Build for web

Execute with `./build_web.sh` It would generate `jetpackgame.wasm` and `wasm_exec.js` inside build_web folder

## Play in web browser

After building for web, use a web server like XAMPP or live server extension in VSCode `./build_web/index.html`

## Execute locally

Execute with `go run .`

## Execute the unit tests

Execute with `go test`

### Next Tasks

- [ ] Level editor
- [ ] Level with a vulcan that throws fire balls.
- [ ] Other enemy that moves vertically an shots fire balls horizontally.
- [ ] Platform with a lake of lava that drops lava.
- [ ] Flip an image and use only one image instead two (left and right)
- [ ] Make the collision box smaller to not collision so early
- [ ] When last level show congratulations game completed.
- [ ] Make more white the background leg of the player
- [ ] In web version auto click the canvas to allow using keys without clicking first the canbas with mouse (view ebitengine web examples)
- [ ] Unit testing and error handling.
- [ ] Extract the levels data to txt files and embed them
- [ ] Upload the build web version into my portfolio to allow play in web browser.
- [ ] Create a basic starting menu. The player must press the space key to start the game.
- [ ] When game over the player can press spacebar or enter to start.
- [x] Reduce to two the arguments of collisions functions: (obj1, obj2). All the necessary information will be inside the objects.
- [x] Fix bug: when the player is inside the rocket, not allow to kill the player by an enemy.
- [x] Use my image utilities to flip images and delete unused flipped images.
- [x] Use my image utilities to move images.
