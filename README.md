# route2bimmer - Create route files for your BMW
![GitHub release (latest by date)](https://img.shields.io/github/v/release/organized92/route2bimmer) ![GitHub](https://img.shields.io/github/license/organized92/route2bimmer)

## Important Note
__Please note that this project has been discontinued.__

In its current state, route2bimmer is mostly compatible to NBT EVO and CIC navigation systems. Please be aware, that your final destination may not be too close to your route, because the navigation system will then probably take a shortcut :-) Roundtrips won't work as expected. The generated routes will not work at all on NBT navigation systems.

The reason for the discontinuation is the "AgoraCString" inside the route files. AGORA-C is a closed-source and patented industry standard for map-based location referencing. Without this "AgoraCString" the navigation system will not strictly follow the given route. NBT won't do anything.

If you know more about AGORA-C and want to contribute to route2bimmer, feel free to contact me.

## Description
route2bimmer is a command-line tool that converts a GPX file into a route file, which you can load onto your BMW (compatible navigation system required).

## Usage
route2bimmer requires at minimum a GPX file containing route data. The GPX file may also include track data so that route2bimmer is able to calculate the total length and driving duration of the route (which will be shown in your navigation system).
``` bash
route2bimmer --input="path-to-input.gpx" --output="path-to-output-route.zip"
```

If you prefer, you can also use stdin and stdout.
``` bash
route2bimmer < path-to-input.gpx > path-to-output-route.zip
```

You can also have a look at the built in usage help:
``` bash
route2bimmer -h
```

## Credits
Thanks to
* __Matthias Nord__ for testing
* __lanc0r__ for providing sample route files
