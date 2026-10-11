# Graph Report - jetpack-game  (2026-10-09)

## Corpus Check
- 116 files · ~71,359 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 38 file(s) not represented in the graph (top: .piskel 28, .blend 5, (none) 2)

## Summary
- 398 nodes · 821 edges · 32 communities (20 shown, 12 thin omitted)
- Extraction: 88% EXTRACTED · 12% INFERRED · 0% AMBIGUOUS · INFERRED: 98 edges (avg confidence: 0.83)
- Token cost: 0 input · 0 output

## Community Hubs (Navigation)
- Game Loop & Typewriter
- Background & Image Utils
- HUD & Progress Bar
- Player Character
- Particle System
- Web Build & Docs
- Audio & Resource Loading
- Level Loading
- Sprite & Collision
- Floor System
- Rocket
- Enemy & Hazard Sprites
- Fuel Collectible
- Character SVG Rig
- Game Lifecycle
- Lava Drop
- Planet & Starfield Sprites
- Menu & HUD Imagery
- HUD Sprite Assets
- Player Sprite Assets
- Tree Sprites & README
- Vertical Text Drawing
- Platform & Collectible Sprites
- Particle VFX
- Editor Config
- Floor Tiles
- Banner & Icon
- Run Script
- Web Build Script
- Test Script
- Module Root

## God Nodes (most connected - your core abstractions)
1. `Planets` - 39 edges
2. `Player` - 38 edges
3. `Floor` - 31 edges
4. `Game` - 27 edges
5. `Hud` - 27 edges
6. `Progressbar` - 27 edges
7. `Rocket` - 27 edges
8. `Fuel` - 25 edges
9. `Lavadrop` - 21 edges
10. `DrawNormalImage()` - 20 edges

## Surprising Connections (you probably didn't know these)
- `Lava Drop Projectile Sprite` --references--> `NewLavadrop()`  [INFERRED]
  assets/img/enemies/lava_drop.png → objects/lavadrop.go
- `Mainmenu (1024x768 Pixel-Art Menu Backdrop)` --references--> `NewMainmenu()`  [INFERRED]
  assets/img/backgrounds/mainmenu.png → objects/mainmenu.go
- `Blinking Star (4-Frame Twinkle Sprite Sheet)` --references--> `NewBlinkingStar()`  [INFERRED]
  assets/img/backgrounds/blinking_star.png → objects/blinkingstar.go
- `Enemy1 Base Sprite (Idle / Open Eyes)` --references--> `NewEnemy()`  [INFERRED]
  assets/img/enemies/enemy1.png → objects/enemy.go
- `Pillar (30x75 Stone Column Platform Sprite)` --references--> `NewPlatform()`  [INFERRED]
  _design/platforms/pillar/pillar.png → objects/platform.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Web Build Pipeline** — readme_build_for_web, build_web_main_wasm_page, build_web_main_jetpackgame_wasm, build_web_main_wasm_exec [INFERRED 0.85]
- **Contribution Process** — contributing_guidelines, contributing_tdd, contributing_commit_conventions, contributing_pull_request_workflow [EXTRACTED 1.00]
- **Game Technology Stack** — readme_go, readme_ebitengine, readme_bfxr, readme_inkscape, readme_piskel, readme_blender [INFERRED 0.75]
- **Composed Jetpack Character Rig** — _design_character_animation_character, _design_character_animation_parts_background, _design_character_animation_parts_fire, _design_character_animation_parts_head, _design_character_animation_parts_jetpack, _design_character_animation_parts_left_arm, _design_character_animation_parts_left_leg, _design_character_animation_parts_right_arm, _design_character_animation_parts_right_leg, _design_character_animation_parts_torso [EXTRACTED 1.00]
- **HUD UI Assets** — _design_hud_hud, _design_hud_live, _design_hud_circle_planet_selector [INFERRED 0.85]
- **Space Background Assets** — _design_planets_starfield1, _design_planets_starfield2, _design_mainmenu_mainmenu [INFERRED 0.75]
- **Pixel-Art Space Scene Asset Set** — assets_img_backgrounds_background1, assets_img_backgrounds_background2, assets_img_backgrounds_blinking_star, assets_img_backgrounds_mainmenu [INFERRED 0.75]
- **Enemy1 Blink Animation Cycle** — assets_img_enemies_enemy1, assets_img_enemies_enemy1_closing_eyes, assets_img_enemies_enemy1_opening_eyes [EXTRACTED 1.00]
- **HUD Sprite Asset Set** — assets_img_hud_circle_planet_selector, assets_img_hud_cloud, assets_img_hud_heart, assets_img_hud_hud, assets_img_hud_live [INFERRED 0.85]
- **Sprites Rendered Through particles.ParticlesSystem** — assets_img_particles_smoke, assets_img_particles_explosion, assets_img_particles_fire [INFERRED 0.85]
- **Planet Sprite Set** — assets_img_planets_fire_planet, assets_img_planets_green_planet, assets_img_planets_ice_planet [INFERRED 0.85]
- **Starfield Parallax Layers** — assets_img_planets_starfield1, assets_img_planets_starfield2 [INFERRED 0.95]
- **Rocket Fuel Indicator States** — assets_img_rocket_rocket, assets_img_rocket_rocket_fuel_indicator_on, assets_img_rocket_rocket_fuel_indicator_off [INFERRED 0.85]
- **Player Character Sprite Set** — assets_img_player_player_center, assets_img_player_player_right, assets_img_player_player_right_with_fuel, assets_img_player_player_walk_right, assets_img_player_player_walk_right_with_fuel [INFERRED 0.85]
- **Jetpack Thruster Flame Overlays Attached to Player Sprites** — assets_img_player_fire_center, assets_img_player_fire_right, assets_img_player_player_center, assets_img_player_player_right [INFERRED 0.75]
- **Fire Tree Sprite Set (variants/frames)** — assets_img_trees_fire_tree_01, assets_img_trees_fire_tree_02, assets_img_trees_fire_tree_03, assets_img_trees_fire_tree_04, assets_img_trees_fire_tree_05 [INFERRED 0.85]
- **GitHub README Showcase Media** — img_github_readme_banner, img_github_readme_screenshots [INFERRED 0.75]

## Communities (32 total, 12 thin omitted)

### Community 0 - "Game Loop & Typewriter"
Cohesion: 0.08
Nodes (10): TypeWriter, NewTypeWriter(), Blackfader, Mainmenu, NewMainmenu(), CalculateTreePositionX(), Tree, NewTree() (+2 more)

### Community 1 - "Background & Image Utils"
Cohesion: 0.11
Nodes (16): Pillar (30x75 Stone Column Platform Sprite), Background1 (Near Space Pixel-Art Layer), Background2 (Far Deep-Space Pixel-Art Layer), Blinking Star (4-Frame Twinkle Sprite Sheet), Mainmenu (1024x768 Pixel-Art Menu Backdrop), DrawHorizontalFlippedImage(), DrawLightenImage(), DrawNormalImage() (+8 more)

### Community 2 - "HUD & Progress Bar"
Cohesion: 0.12
Nodes (4): Hud, NewHud(), Progressbar, NewProgressbar()

### Community 3 - "Player Character"
Cohesion: 0.13
Nodes (3): Player, NewPlayer(), PlayerStatus

### Community 4 - "Particle System"
Cohesion: 0.12
Nodes (8): NewBlackfader(), NewRocket(), Particle, newParticle(), ParticlesSystem, NewExplosion(), NewSmoke(), NewPlanets()

### Community 5 - "Web Build & Docs"
Cohesion: 0.12
Nodes (20): First Lava Floor (Level 1), Fire Everywhere (Level 2), Jetpack Game Landing Page, jetpackgame.wasm, instantiateStreaming Polyfill, wasm_exec.js, WASM Loader Script, WASM Game Page (+12 more)

### Community 6 - "Audio & Resource Loading"
Cohesion: 0.12
Nodes (5): Sound, LoadStaticResource(), LoadFonts(), LoadSounds(), NewSoundFromFile()

### Community 7 - "Level Loading"
Cohesion: 0.15
Nodes (10): Level, getLevelEntries(), getLevelFiles(), isLevelValid(), isLineValid(), LoadLevels(), NewLevel(), assertError() (+2 more)

### Community 8 - "Sprite & Collision"
Cohesion: 0.17
Nodes (8): Game, GameObject, GameStatus, checkCollision(), isColliding(), LoadFolderSprites(), loadImage(), LoadSprites()

### Community 9 - "Floor System"
Cohesion: 0.23
Nodes (5): GetFloorType(), Floor, InitializeFloor(), NewFloor(), FloorType

### Community 11 - "Enemy & Hazard Sprites"
Cohesion: 0.22
Nodes (9): Enemy1 Base Sprite (Idle / Open Eyes), Enemy1 Closing Eyes Animation Sheet, Enemy1 Opening Eyes Animation Sheet, Lava Drop Projectile Sprite, Ice Floor Sprite (Unreferenced Variant), Lava Floor Animated Sprite Sheet, Vulcan (Normal) Floor Sprite, Enemy (+1 more)

### Community 13 - "Character SVG Rig"
Cohesion: 0.20
Nodes (11): Composed Jetpack Character (SVG), Character Part Decomposition Rig, Character Background Layer, Jetpack Flame Animation, Character Head, Jetpack, Left Arm, Left Leg (+3 more)

### Community 14 - "Game Lifecycle"
Cohesion: 0.27
Nodes (3): Game, NewGame(), main()

### Community 16 - "Planet & Starfield Sprites"
Cohesion: 0.50
Nodes (8): Fire Planet Sprite, Green Planet Sprite, Ice Planet Sprite, Starfield Layer 1, Starfield Layer 2, Rocket Sprite, Rocket Fuel Indicator (Off), Rocket Fuel Indicator (On)

### Community 17 - "Menu & HUD Imagery"
Cohesion: 0.29
Nodes (7): GitHub README Screenshots, Circle Planet Selector, HUD Overlay, Live Icon, Main Menu Background, Starfield 1, Starfield 2

### Community 18 - "HUD Sprite Assets"
Cohesion: 0.43
Nodes (7): Circle Planet Selector Sprite, Cloud Sprite, Heart Sprite, HUD Panel Sprite, HUD Sprite Asset Set, Live Indicator Sprite, Lives / Health Display Concept

### Community 19 - "Player Sprite Assets"
Cohesion: 0.43
Nodes (7): Fire Center Thruster Flame Sprite, Fire Right Thruster Flame Sprite, Player Center Idle Sprite, Player Right Facing Sprite, Player Right Facing Sprite with Fuel Indicator, Player Walk Right Animation Strip, Player Walk Right Animation Strip with Fuel Indicator

### Community 20 - "Tree Sprites & README"
Cohesion: 0.62
Nodes (7): Fire Tree Sprite 01, Fire Tree Sprite 02, Fire Tree Sprite 03, Fire Tree Sprite 04, Fire Tree Sprite 05, GitHub README Banner (700x140), GitHub README Screenshots Showcase (1100x432)

### Community 22 - "Platform & Collectible Sprites"
Cohesion: 0.50
Nodes (4): Fuel (Collectible Fuel Canister Sprite), Parachute (Canopy Drawn Above Falling Fuel), Pillar (Stacked Column Segment Beneath Platform), Platform (Floating Platform Surface Sprite)

### Community 23 - "Particle VFX"
Cohesion: 0.67
Nodes (3): Explosion (Player-Death Particle VFX Sprite), Fire (Lava Floor Flame Particle Sprite), Smoke (Rocket Exhaust Particle Sprite)

## Ambiguous Edges - Review These
- `Ice Floor Sprite (Unreferenced Variant)` → `Vulcan (Normal) Floor Sprite`  [AMBIGUOUS]
  assets/img/floors/ice_floor.png · relation: conceptually_related_to
- `Circle Planet Selector Sprite` → `Lives / Health Display Concept`  [AMBIGUOUS]
  assets/img/hud/circle_planet_selector.png · relation: conceptually_related_to

## Knowledge Gaps
- **30 isolated node(s):** `Game`, `github.com/darellanodev/jetpack-game`, `Bfxr`, `Inkscape`, `Piskel` (+25 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 113 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **12 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `Player` connect `Player Character` to `Game Loop & Typewriter`, `Background & Image Utils`, `Game Lifecycle`, `Audio & Resource Loading`?**
  _High betweenness centrality (0.091) - this node is a cross-community bridge._
- **What connects `Game`, `github.com/darellanodev/jetpack-game`, `Bfxr` to the rest of the system?**
  _30 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `Game Loop & Typewriter` be split into smaller, more focused modules?**
  _Cohesion score 0.08305647840531562 - nodes in this community are weakly interconnected._
- **Why does `Planets` connect `Background & Image Utils` to `Game Loop & Typewriter`, `HUD & Progress Bar`, `Player Character`, `Particle System`, `Audio & Resource Loading`, `Floor System`, `Rocket`, `Enemy & Hazard Sprites`, `Fuel Collectible`, `Game Lifecycle`, `Lava Drop`?**
  _High betweenness centrality (0.079) - this node is a cross-community bridge._
- **Should `Background & Image Utils` be split into smaller, more focused modules?**
  _Cohesion score 0.10668563300142248 - nodes in this community are weakly interconnected._

### Low-confidence Hints
_AMBIGUOUS edges — the extractor was unsure. Verify before acting on these._

- **What is the exact relationship between `Ice Floor Sprite (Unreferenced Variant)` and `Vulcan (Normal) Floor Sprite`?**
  _Edge tagged AMBIGUOUS (relation: conceptually_related_to) - confidence is low._
- **What is the exact relationship between `Circle Planet Selector Sprite` and `Lives / Health Display Concept`?**
  _Edge tagged AMBIGUOUS (relation: conceptually_related_to) - confidence is low._