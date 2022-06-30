# !!!MAKE SURE YOUR GOPATH ENVIRONMENT VARIABLE IS SET FIRST!!!

# Agent file names
W=Windows-x64
L=Linux-x64
B=FreeBSD-x64
A=Linux-arm
M=Linux-mips
D=Darwin-x64

# Merlin version number
VERSION=$(shell cat ./core/core.go |grep "var Version ="|cut -d"\"" -f2)

MAGENT=merlinAgent
PASSWORD=merlin
BUILD=$(shell git rev-parse HEAD)
DIR=bin/v${VERSION}/${BUILD}

# Merlin Agent Variables
XBUILD=-X "main.build=${BUILD}" -X "github.com/testtoto1337/merzhin-agent/agent.build=${BUILD}"
URL ?= https://127.0.0.1:443
XURL=-X "main.url=${URL}"
PSK ?= merlin
XPSK=-X "main.psk=${PSK}"
PROXY ?=
XPROXY =-X "main.proxy=$(PROXY)"
SLEEP ?= 30s
XSLEEP =-X "main.sleep=$(SLEEP)"
HOST ?=
XHOST =-X "main.host=$(HOST)"
PROTO ?= h2
XPROTO =-X "main.protocol=$(PROTO)"
JA3 ?=
XJA3 =-X "main.ja3=$(JA3)"
USERAGENT = Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2214.85 Safari/537.36
XUSERAGENT =-X "main.useragent=$(USERAGENT)"
HEADERS =
XHEADERS =-X "main.headers=$(HEADERS)"
SKEW ?= 3000
XSKEW=-X "main.skew=${SKEW}"
PAD ?= 4096
XPAD=-X "main.padding=${PAD}"
KILLDATE ?= 0
XKILLDATE=-X "main.killdate=${KILLDATE}"
RETRY ?= 7
XRETRY=-X "main.maxretry=${RETRY}"

# Compile Flags
LDFLAGS=-ldflags '-s -w ${XBUILD} ${XPROTO} ${XURL} ${XHOST} ${XPSK} ${XSLEEP} ${XPROXY} $(XUSERAGENT) $(XHEADERS) ${XSKEW} ${XPAD} ${XKILLDATE} ${XRETRY} -buildid='
WINAGENTLDFLAGS=-ldflags '-s -w ${XBUILD} ${XPROTO} ${XURL} ${XHOST} ${XPSK} ${XSLEEP} ${XPROXY} $(XUSERAGENT) $(XHEADERS) ${XSKEW} ${XPAD} ${XKILLDATE} ${XRETRY} -H=windowsgui -buildid='
GCFLAGS=-gcflags=all=-trimpath=$(GOPATH)
ASMFLAGS=-asmflags=all=-trimpath=$(GOPATH)# -asmflags=-trimpath=$(GOPATH)

# Package Command
PACKAGE=7za a -p${PASSWORD} -mhe -mx=9
F=LICENSE

# Misc
# GOGARBLE contains a list of all the packages to obfuscate
GOGARBLE=golang.org,gopkg.in,github.com
# The Merlin server and agent MUST be built with the same seed value
# Set during build with "make linux-garble SEED=<insert seed>
SEED=d0d03a0ae4722535a0e1d5d0c8385ce42015511e68d960fadef4b4eaf5942feb

# Make Directory to store executables
$(shell mkdir -p ${DIR})

# Change default to just make for the host OS and add MAKE ALL to do this
default:
	go build -trimpath ${LDFLAGS} ${GCFLAGS} ${ASMFLAGS} -o ${DIR}/${MAGENT} ./main.go

all: windows linux darwin

# Compile Agent - Windows x64
windows:
	export GOOS=windows GOARCH=amd64;go build -trimpath ${WINAGENTLDFLAGS} ${GCFLAGS} ${ASMFLAGS} -o ${DIR}/${MAGENT}-${W}.exe ./main.go

# Compile Agent - Windows x64 Debug (Can view STDOUT)
windows-debug:
	export GOOS=windows GOARCH=amd64;go build -trimpath ${LDFLAGS} ${GCFLAGS} ${ASMFLAGS} -o ${DIR}/${MAGENT}-Debug-${W}.exe ./main.go

# Compile  Agent - Windows x64 with Garble - The SEED must be the exact same that was used when compiling the server
# Garble version 0.5.2 or later must be installed and accessible in the PATH environment variable
windows-garble:
	export GOGARBLE=${GOGARBLE};export GOOS=windows GOARCH=amd64;garble -tiny -literals -seed ${SEED} build -trimpath ${WINAGENTLDFLAGS} ${GCFLAGS} ${ASMFLAGS} -o ${DIR}/${MAGENT}-${W}.exe ./main.go

# Compile Agent - Linux mips
mips:
	export GOOS=linux;export GOARCH=mips;go build -trimpath ${LDFLAGS} ${GCFLAGS} ${ASMFLAGS} -o ${DIR}/${MAGENT}-${M} ./main.go

# Compile Agent - Linux arm
arm:
	export GOOS=linux;export GOARCH=arm;export GOARM=7;go build -trimpath ${LDFLAGS} ${GCFLAGS} ${ASMFLAGS} -o ${DIR}/${MAGENT}-${A} ./main.go

# Compile Agent - Linux x64
linux:
	export GOOS=linux;export GOARCH=amd64;go build -trimpath ${LDFLAGS} ${GCFLAGS} ${ASMFLAGS} -o ${DIR}/${MAGENT}-${L} ./main.go

# Compile  Agent - Linux x64 with Garble - The SEED must be the exact same that was used when compiling the server
# Garble version 0.5.2 or later must be installed and accessible in the PATH environment variable
linux-garble:
	export GOGARBLE=${GOGARBLE};export GOOS=linux GOARCH=amd64;garble -tiny -literals -seed ${SEED} build -trimpath ${LDFLAGS} ${GCFLAGS} ${ASMFLAGS} -o ${DIR}/${MAGENT}-${L} ./main.go

# Compile Agent - FreeBSD x64
freebsd:
	export GOOS=freebsd;export GOARCH=amd64;go build -trimpath ${LDFLAGS} ${GCFLAGS} ${ASMFLAGS} -o ${DIR}/${MAGENT}-${B} ./main.go

# Compile  Agent - FreeBSD x64 with Garble - The SEED must be the exact same that was used when compiling the server
# Garble version 0.5.2 or later must be installed and accessible in the PATH environment variable
freebsd-garble:
	export GOGARBLE=${GOGARBLE};export GOOS=freebsd GOARCH=amd64;garble -tiny -literals -seed ${SEED} build -trimpath ${LDFLAGS} ${GCFLAGS} ${ASMFLAGS} -o ${DIR}/${MAGENT}-${B} ./main.go

# Compile Agent - Darwin x64
darwin:
	export GOOS=darwin;export GOARCH=amd64;go build -trimpath ${LDFLAGS} ${GCFLAGS} ${ASMFLAGS} -o ${DIR}/${MAGENT}-${D} ./main.go

# Compile  Agent - macOS (Darwin) x64 with Garble - The SEED must be the exact same that was used when compiling the server
# Garble version 0.5.2 or later must be installed and accessible in the PATH environment variable
darwin-garble:
	export GOGARBLE=${GOGARBLE};export GOOS=darwin GOARCH=amd64;garble -tiny -literals -seed ${SEED} build -trimpath ${LDFLAGS} ${GCFLAGS} ${ASMFLAGS} -o ${DIR}/${MAGENT}-${D} ./main.go

package-windows:
	${PACKAGE} ${DIR}/${MAGENT}-${W}.7z ${F}
	cd ${DIR};${PACKAGE} ${MAGENT}-${W}.7z ${MAGENT}-${W}.exe

package-linux:
	${PACKAGE} ${DIR}/${MAGENT}-${L}.7z ${F}
	cd ${DIR};${PACKAGE} ${MAGENT}-${L}.7z ${MAGENT}-${L}

package-darwin:
	${PACKAGE} ${DIR}/${MAGENT}-${D}.7z ${F}
	cd ${DIR};${PACKAGE} ${MAGENT}-${D}.7z ${MAGENT}-${D}

package-freebsd:
	${PACKAGE} ${DIR}/${MAGENT}-${B}.7z ${F}
	cd ${DIR};${PACKAGE} ${MAGENT}-${B}.7z ${MAGENT}-${D}

clean:
	rm -rf ${DIR}*

package-all: package-windows package-linux package-darwin

#Build all files for release distribution
distro: clean all package-all
