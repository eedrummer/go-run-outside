include $(GOROOT)/src/Make.inc

TARG=main
GOFILES=\
	runlist.go\
	activity.go\
	conversion.go\
	oauth.go\
	nikeplus.go\
	runkeeper.go\
	main.go

include $(GOROOT)/src/Make.cmd