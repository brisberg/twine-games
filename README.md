# Twine Games
Set of simple browser games built in Twine

## Install

Install [golang](https://golang.org/doc/install) for your system.
Install the Tweego compiled
`go get github.com/tmedwards/tweego`

## Editing

### Tweego

Use the [Tweego](https://github.com/tmedwards/tweego) twee compiled to compile `.twee` files into build `.html` files ready for viewing.

* Make edits to the twee files for your game.
* Compile it for viewing/distribution
  * `go run github.com/tmedwards/tweego games/<game> -o dist/<game>.html (-f <format>)`

#### Story Formats

The Tweego compiled requires a local copy of any Storyformat you use, located in the `storyformats/` directory. If you need a new format you will need to get them yourself, and add them to the formats folder.

For a list of all formats Tweego has access to: `go run github.com/tmedwards/tweego formats`

### Online-editor

Less perferable, but if you'd like a graphical view of the story structure the compiled `.html` file can be edited using the online Twine 2 Editor.

**NOTE**: This will only work for games built in a single `.twee` file.

* Open the Editor at https://twinery.org/2
* `Import from File` and select the compiled `<story>.html` file.
* Make any edits / rearrange stories
* `Publish to File` and overrite the old `<story>.html` file.
* Run the Tweego decompiler to backport your changes to the `.twee file.
  * `go run github.com/tmedwards/tweego dist/<game>.html -o games/<game>/<game>.twee -d (-f <format>)`
