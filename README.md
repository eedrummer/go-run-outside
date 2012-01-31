Go Run Outside
==============

This project is a command line application that moves run data from Nike+ and move it to RunKeeper.

Why?
----

This is a hobby project to teach myself the Go programming language. I am aware that there are solutions out there to move runs from Nike+ to RunKeeper already.

Install
-------

You'll need to have [Go](http://golang.org/) installed.

    git clone https://github.com/eedrummer/go-run-outside.git
    cd go-run-outside
    gomake
    ./main -u YOUR_NIKEPLUS_ID

Your Nike+ id is usually visible in as a query parameter when signed in and viewing the Nike+ web site.

About the Makefile
------------------

I am terrible with Makefiles. It is currently set up to build the executable. If you want to run the test suite change the last line from

    include $(GOROOT)/src/Make.cmd

to

    include $(GOROOT)/src/Make.pkg

Any Makefile help/suggestions would be appreciated.

License
-------

Copyright 2012 Andy Gregorowicz

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.