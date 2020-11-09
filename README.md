# Twine Games (Deprecated)

Set of simple browser games built in Twine

Play them [online](https://brisberg.github.io/twine-games/)!

Deprecated 11/2020 in favor of modular repository structure - [twine.brisberg.dev](https://github.com/brisberg/twine.brisberg.dev).

## Install

Install [golang](https://golang.org/doc/install) for your system.\
Install the Tweego compiled binary with:\
`go get github.com/tmedwards/tweego`

## Editing

### Tweego

Use the [Tweego](https://github.com/tmedwards/tweego) twee compiler to compile `.twee` or `.tw` files into built `.html` files ready for viewing.

* Make edits to the `.twee` or `.tw` files for your game.
* Compile it for viewing/distribution
  * `go run builder/*`

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

### Analytics

This project uses Google Analytics and Google Tag Manager to track user engagement.

The `analytics.html` file initializes a GTM dataLayer for the container of all TwineGames: `GTM-T6XR5LZ`. This does not need to be modified per game, just set up a unique trigger/tag for each game and use the GameName as the `event` name.

Be sure to forward the GTM tag to a specific GA Property for each game.

#### Events

Because Twine games are run as a single page, GTM will only pick up a single page view event by default. If you want to track other engagement, you must provide custom events (such as "Start", "Finish", "Dead", etc).

Follow this [naming scheme](https://mixedanalytics.com/blog/event-tracking-naming-strategy-for-google-analytics/) for custom events.
