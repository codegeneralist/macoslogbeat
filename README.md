# MacosLogbeat

Welcome to MacosLogbeat.

Ensure that this folder is at the following location:
`${GOPATH}/src/github.com/jaakkoo/macoslogbeat`

## Getting Started with MacosLogbeat

### Requirements

* [Golang](https://golang.org/dl/) 1.14
* [mage](https://github.com/magefile/mage) 1.8.0

For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).

Make sure you have ${GOPATH} set and ${GOPATH}/bin in PATH.

### Build

To build the binary for MacosLogbeat run the command below. This will generate a binary
in the same directory with the name macoslogbeat.

```
make
```


### Run

To run MacosLogbeat with debugging output enabled, run:

```
./macoslogbeat -c macoslogbeat.yml -e -d "*"
```


### Test

To test MacosLogbeat, run the following command:
(As of writing this there are no tests, but maybe this changes over time)

```
make testsuite
```

alternatively:
```
make unit-tests
make system-tests
make integration-tests
make coverage-report
```

The test coverage is reported in the folder `./build/coverage/`

### Update

Each beat has a template for the mapping in elasticsearch and a documentation for the fields
which is automatically generated based on `fields.yml` by running the following command.

```
make update
```


### Cleanup

To clean  MacosLogbeat source code, run the following command:

```
make fmt
```

To clean up the build directory and generated artifacts, run:

```
make clean
```


### Clone

To clone MacosLogbeat from the git repository, run the following commands:

```
mkdir -p ${GOPATH}/src/github.com/jaakkoo/macoslogbeat
git clone https://github.com/jaakkoo/macoslogbeat ${GOPATH}/src/github.com/jaakkoo/macoslogbeat
```


For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).


## Packaging

The beat frameworks provides tools to crosscompile and package your beat for different platforms. This requires [docker](https://www.docker.com/) and vendoring as described above. To build packages of your beat, run the following command:

```
make pkg
```

This will fetch and create all images required for the build process. The whole process to finish can take several minutes.

It will also generate Macos installer (.pkg) and it is made available in the same directory (build/) than everything
else.

## Install

Once macoslogbeat-<version>.pkg package is installed the files will be available in `/opt/macoslogbeat`

The easiest way to verify the application works is to run it manually from commandline:
`sudo /opt/macoslogbeat/macoslogbeat -c /opt/macoslogbeat/macoslogbeat.yml -e`. 

Get the help: `sudo /opt/macoslogbeat/macoslogbeat -h`

### Configuration
All configuration is located in `/opt/macoslogbeat/macoslogbeat.yml`

### Run automatically

The launchd configuration is available in `/opt/macoslogbeat/install/com.reaktor.macoslogbeat.plist`.

Install it: `sudo launchctl load /opt/macoslogbeat/install/com.reaktor.macoslogbeat.plist`

Start it: `sudo launchctl start com.reaktor.macoslogbeat`.

Launcd will configure `macoslogbeat` to start automatically when OS boots. If you wish to change the behaviour edit the plist.

## Uninstall

First unload the launchd config
`sudo launchctl unload /opt/macoslogbeat/install/com.reaktor.macoslogbeat.plist`

Remove the package from database
`sudo pkgutil --forget com.reaktor.macoslogbeat`

Delete all files
`rm -rf /opt/macoslogbeat`
