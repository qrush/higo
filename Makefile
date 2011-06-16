include vendor/Make.inc

TARG=higo
GOFMT=gofmt

GOFILES=\
				higo.go\

include vendor/Make.cmd

format:
	${GOFMT} -w ${GOFILES}
