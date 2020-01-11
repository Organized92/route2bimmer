# route2bimmer - Create route files for your BMW
![GitHub release (latest by date)](https://img.shields.io/github/v/release/organized92/route2bimmer) ![GitHub](https://img.shields.io/github/license/organized92/route2bimmer) 

***Please note that route2bimmer is not functional at it's current development state!***

## Description
route2bimmer is a command-line tool which converts GPX file into routes that you can load up into your BMW (compatible navigation system required).

## Usage
route2bimmer requires at minimum a GPX file including waypoint data. The GPX file may also include track data so that route2bimmer is able to calculate the total length of the route (which will be shown in your navigation system).
``` bash
route2bimmer -input=path-to-input.gpx -output=path-to-output-route.zip
```
In case you forget to add the zip file extension for the output file, route2bimmer will add it automatically.

You can also have a look at the built in usage help:
``` bash
route2bimmer -h
```