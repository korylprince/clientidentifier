clientidentifier

https://github.com/korylprince/clientidentifier

A simple go program to make changing the [munki](http://code.google.com/p/munki/) ClientIdentifier easier.

# Installing

You should be able to go get clientidentifier:

	$ go get github.com/korylprince/clientidentifier
	$ go install github.com/korylprince/clientidentifier
	$ $GOPATH/bin/clientidentifier -h

Also included are `build.sh` which will build a binary, and `mkpkg.sh` which will build a OS X pkg. Note that [pkggen](https://github.com/korylprince/pkggen) is required for `mkpkg.sh`.

If you have any issues or questions, email the email address below, or open an issue at:
https://github.com/korylprince/clientidentifier/issues

# Usage

	$ clientidentifier -h
	Usage: clientidentifier [OPTION...] [IDENTIFIER]

		-h, --help	Display this help message

	Running this program without any options will display the current ClientIdentifier.
	The ClientIdentifier will be changed to IDENTIFIER if given.

# Copyright Information#

Copyright 2020 Kory Prince (korylprince AT gmail DAWT com).

License is BSD.
