# Pixel topographic map generator

The code was born during an improvised hackathon as a result of a friendly dispute about procedural generation and terrain maps. The application purpose is to generate pixel-art terrain maps of a given size using parameterized simplex noise for terrain altitudes and a predefined set of rules for terrain types and colors.

Work is still in progress.

## Badges

[![Built with Ebitengine](https://img.shields.io/badge/built%20with-Ebitengine-%23db5620?style=for-the-badge&logo=github)](https://github.com/hajimehoshi/ebiten)

![Commit activity](https://img.shields.io/github/commit-activity/m/oleg-vasiliev/mapgen?style=for-the-badge&logo=github)

![Last Release](https://img.shields.io/github/v/release/oleg-vasiliev/mapgen?include_prereleases&sort=semver&display_name=release&logo=github&style=for-the-badge)

[![GPLv3 License](https://img.shields.io/badge/License-GPL%20v3-yellow.svg?style=for-the-badge)](https://opensource.org/license/gpl-3-0/)


## Screenshots
- #### Main window with UI hints
![App Screenshot](/docs/screenshots/window-small.png)
- #### Maps exported to image file
![App Screenshot](/docs/screenshots/exported-a.jpg)
![App Screenshot](/docs/screenshots/exported-b.jpg)


## Tech Stack
[Go](https://go.dev/) with [Ebitengine](https://ebitengine.org/) and zero third-party dependencies.


## Build Locally 

Make sure you have the [Go](https://go.dev/) installed
```bash
go version
```
Clone the repo
```bash
  git clone https://github.com/oleg-vasiliev/mapgen.git
```

Go to the project directory
```bash
  cd mapgen
```

Download and verify dependencies 
```bash
  go mod download
  go mod verify
```

Compile and start the app
```bash
  go run ./cmd/mapgen
```

Or just build the binary
```bash
  go build -o=./bin/ ./cmd/mapgen
```

## Usage and cli-arguments

TODO

## Roadmap

- Configuration seed should include all various noise settings.

- Display in UI current terrain levels.

- Implement different color themes.

- Add visual indication for world exporting process.

- Add to map randomly placed objects: trees, animals, buildings, etc.

- Implement inertial viewport scrolling.


## Feedback

If you have any feedback please contact me by email.


## License

[GNU GPLv3](https://opensource.org/license/gpl-3-0/)
