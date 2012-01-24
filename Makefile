include $(GOROOT)/src/Make.inc

TARG=main
GOFILES=\
	runlist.go\
	activity.go\
	conversion.go\
	oauth.go\
	main.go

include $(GOROOT)/src/Make.pkg