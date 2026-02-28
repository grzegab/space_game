# Asteroids GO Game

A simple arcade game set in space, written in Go. The project is a modern interpretation of 80s classics, combining mechanics known from **Asteroids** (inertia, ship rotation) with the atmosphere of **Space Invaders** (incoming waves of aliens).

## About the Game

In the game, you take on the role of a spaceship pilot whose task is to survive in an asteroid belt overrun by hostile alien units. The game offers dynamic retro gameplay with a modern engine under the hood.

### Key Features:
- **Retro Mechanics:** Smooth ship control with physics (acceleration and rotation).
- **Diverse Enemies:** Different types of aliens with varying behavior patterns and intelligence (IQ).
- **Obstacles:** Meteors of various sizes that must be avoided or destroyed.
- **Defense Systems:** Energy shield protecting against collisions and the ability to perform a "hyperspace jump" in critical situations.
- **Audiovisual Setting:** Pixel-art graphics and sound effects in the style of arcade games.

## Tech Stack

The project was built using the modern Go language ecosystem:

- **Language:** [Go](https://go.dev/) (version 1.25+)
- **Game Engine:** [Ebitengine (Ebiten v2)](https://ebitengine.org/) – an open-source 2D game engine for Go, providing efficient rendering and input handling.
- **Collision Detection:** [Resolv](https://github.com/solarlune/resolv) – a library dedicated to 2D games for handling physics and collisions.
- **Supporting Dependencies:**
    - `golang.org/x/image`: Support for graphic formats and fonts.
    - `github.com/jfreymuth/oggvorbis`: OGG audio file decoding.
    - `github.com/ebitengine/oto/v3`: Low-level audio support.
- **Assets:** Images, sounds, and fonts are embedded directly into the executable using the `embed` mechanism.

## Running and Compilation

### Requirements
- Go environment installed (version 1.25 or newer recommended).
- System dependencies required by Ebitengine (details in the [Ebitengine documentation](https://ebitengine.org/en/documents/install.html) for your OS).

### Running in Development Mode
To quickly run the game without building a binary:
```bash
go run cmd/game/main.go
```

### Compilation
To build an optimized executable:
```bash
go build -o asteroids-go ./cmd/game
```
The `asteroids-go` file (or `asteroids-go.exe` on Windows) will be ready to run.

## Controls

- **Up Arrow** – Acceleration (main engines)
- **Left / Right Arrow** – Ship rotation
- **Space** – Laser shot
- **S Key** – Shield activation (consumes shield energy)
- **H Key** – Hyperspace jump (teleportation to a random location on the map)
- **Esc Key** – Exit the game

## Authors
- Greg - Lord of PHP Solutions, Master of Golang Services, Warden of RESTful Gates, Keeper of gRPC Streams, Protector of SQL Queries, Overlord of RabbitMQ Channels
