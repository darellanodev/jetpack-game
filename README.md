# go-jetpack-game

A simple game to learn Go and Ebitengine *(THIS APPLICATION IS IN AN EARLY STAGE OF DEVELOPMENT)*

## Build for web

Execute with `./build_web.sh` It would generate `jetpackgame.wasm` and `wasm_exec.js` inside build_web folder

## Play in web browser

After building for web, use a web server like XAMPP or live server extension in VSCode `./build_web/index.html`

## Execute locally

Execute with `go run .`

### Tasks

- [ ] Level editor
- [ ] Level with a vulcan that throws fire balls.
- [ ] Other enemy that moves vertically an shots fire balls horizontally.
- [ ] Platform with a lake of lava that drops lava.
- [ ] Flip an image and use only one image instead two (left and right)
- [ ] Make the collision box smaller to not collision so early
- [ ] When last level show congratulations game completed.
- [ ] When game over the player can press spacebar or enter to start.
- [ ] Make more white the background leg of the player
- [ ] In web version auto click the canvas to allow using keys without clicking first the canbas with mouse (view ebitengine web examples)
- [ ] Extract the levels data to txt files and embed them
- [ ] Create a basic starting menu. The player must press the space key to start the game.
- [x] Create a end game (game completed). The player must press the space key to restart the game.
- [x] HUD with 3 lives.
- [x] Embed audio, video, fonts.
- [x] Make some blinking stars in the background.
- [x] The lava floor kills the player.
- [x] Do not allow the player to leave the screen on the left, right or top.
- [x] Fire effect in lava with particles
- [x] Player explosion with particles
- [x] Smoke when Rocket goes up
- [x] Fire in bottom of the Rocket when it is landing or it goes up.
- [x] After landing, wait a time and then dont update/draw particles if we are in playing state
- [x] Trail of smoke particles when the Rocket is landing.
- [x] In the HUD there is a bar that displays the remaining oxygen for the player and it decreases as time passes.
- [x] When the player dies the player blinks and can not die during 5 seconds
- [x] Show text for "sound on" or "sound off" when the player press the S key.
- [x] Show text for "pause" when the player press P
- [x] When player gets the fuel his arms are in a fixed position getting the fuel (animation walk left/right with fixed arms)
- [x] The fuel lands in parachute
- [x] Put fuel near the hands of the player.
- [x] Enemy animation closing the eye.
- [x] Create a floor with a high of 200px aprox. to allow put holes with lava.
- [x] The player goes to the following level when completes the first level.
- [x] Animage the lava floor.
- [x] Floor with lava.
- [x] Player walking animation.
- [x] Play a sound when the player drops the fuel.
- [x] Play a sound when the rocket moves.
- [x] Play a sound when showing the traveling screen
- [x] Change the background for the second level
- [x] When player has 0 lives show Game Over.
- [x] Fuel sprite and set its initial position.
- [x] The fuel snaps to the player and follows them.
- [x] It disappears from its current position.
- [x] Play a sound when picking up the fuel.
- [x] Place the rocket on the floor.
- [x] If the player has fuel, the fuel reappears at its initial position.
- [x] Play a sound when the player dies.
- [x] Make a fuel indicator in the rocket.
- [x] The fuel indicator goes up when the player drops the fuel.
