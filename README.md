# go-conduit

[Phabricator](http://phabricator.org/) is "an open source, software engineering
platform".  It provides a JSON API for interacting with the system.

This repository provides a bare bones SDK for interacting with the Phabricator
conduit API in go.

## Build

Building an installing this package is as easy as

    go build ./...
    go install ./...

## Disclaimer

This SDK is not production ready and is missing a TON of functionality.  It was
originally developed to support one particular use case I had [at my day
job](https://researchsquare.com).  As such, a full implementation of the
conduit API has not been completed at this time.

If there are new endpoints you want to support, or current ones you want to
extend, I would welcome any pull requests you want to send my way.

Also, I am very new to go, so if you know of a cleaner way to implement any of
this, I would love to hear about it.
