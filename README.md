envprint
--------

[![Build Status](https://travis-ci.org/dylanmei/envprint.svg?branch=master)](https://travis-ci.org/dylanmei/envprint)

Template processor for printing configuration files interpolated with environment variables.

_Equivalent to processing templates using [gliderlabs/sigil](https://github.com/gliderlabs/sigil) with POSIX style variable expansion._

# example

```
$ envprint -help
Usage of envprint:
  -f string
      Use template file instead of STDIN
  -version
      Prints the current version

$ echo 'hello ${WORLD}' | envprint
hello 

$ echo 'hello ${WORLD:-catania}' | envprint
hello catania

$ export WORLD=palermo

$ echo 'hello ${WORLD}' | envprint
hello palermo

$ echo 'hello ${FOOBAR:?this is not a thing}' | envprint
this is not a thing

$ echo $?
1
```
