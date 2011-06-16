include $(GOROOT)/src/Make.inc

TARG=higo
GOFMT=gofmt

GOFILES=\
				main.go\

include $(GOROOT)/src/Make.cmd

format:
	${GOFMT} -w ${GOFILES}
