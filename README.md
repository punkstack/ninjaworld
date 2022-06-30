# Ninja War

## Overview

The goal of ninja war stimulation is to stimulate the
invasion of otsutsukies on earth.
The ninja world is segregated
as per specific villages, and the otsutsukies are dispersed randomly on the ninja world.
Encounter of more than one otsutsuki in a specific cities leads to destruction of city and
destruction of every otsutsuki in that city.

## Implementation

In Ninja world, each village can have maximum four cardinal directions in which there can be a neighbouring village.
A village and the respective neighbouring villages are connected by bidirectionally. For example,

```
Foo north=Bar
```

where, village Foo has a neighbouring village named Bar in the north direction. Also, vice versa, Bar village has a
neighbouring village named Foo in the south direction.

Otsutsukies are randomly deployed in the villages, and if in a village more than one otsutsukies are deployed, they
destroy not only each other but also the city. And the remaining otsutsukies move randomly to the available neighbouring
villages and repeat the same war protocol in that village.

### Packages

[math/rand](https://golang.org/pkg/math/rand/) package to deploy otsutsukies randomly in the villages and also to pick
neighbouring villages.

[Faker](https://pkg.go.dev/syreclabs.com/go/faker) for generating random otsutsuki names.

[Zap loger](https://pkg.go.dev/go.uber.org/zap) for logging purpose.

[Cobra](https://github.com/spf13/cobra) to handle command line flags and validations.

## Installation

[Golang environment](https://golang.org/doc/install) must be installed to run ninja war application.

Use Make commands to build application.

```bash
make build
```

Use Make test to run tests.

```bash
make test
```

Build will generate Darwin and Linux application files. Use Make clean to delete the build.

```bash
make clean
```

## Usage

```bash
$ ./ninja-world-darwin --help
Usage of ./ninja-world-darwin:
   -i, --input-filename string
        input file name with file path which contains villages
   -o, --output-filename string
        onput file name with file path where result of simulation is stored
   -n, --otsutsuki int
        random number used as entropy seed
```

Run the specific simulation by using the below command

```bash
$ ./ninja-world-darwin -n 10 -i tests/input.txt -o tests/result.txt   
```

## Assumptions

* There is no restriction in the number of otsutsukies deploying in a village.

* If a village has more than one otsutsukies are in a village, then not only the otsutsukies are destroyed but also the
  village.

* A village and the neighbouring village are connected bidirectionally.

* The initial energy of an otsutsuki is 10,000 moves and can use one move in a day, after which it will die.

## Next Steps

* Improvement of I/O validation 
* Need more test cases to improve the test coverage
* Need more map examples
* Custom logger implemented but it is not upto the mark
* CI pipelin for tests
