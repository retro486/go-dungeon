# dungeon-go makefile (C) 2010 Russell Bernhardt <russell.bernhardt@gmail.com>
#
# targets: director, client-look, client-move, client-show-message,
#          director-command
#
include $(GOROOT)/src/Make.$(GOARCH)

GOCFLAGS = # extra flags to pass to the go compiler
GOLFLAGS = # extra flags to pass to the go linker

COMMONREQS = helpfuls.go maze.go events.go verbs.go verb_actions.go structures.go
DIRECTORREQS = director-core.go $(COMMONREQS)
DIRECTORCOMMREQS = send-command.go $(COMMONREQS)
CLIENTLOOKREQS = client-look.go  $(COMMONREQS)
CLIENTMOVEREQS = client-move.go $(COMMONREQS)
CLIENTSHOWMESSAGEREQS = show-message.go helpfuls.go

all :
	@echo "Valid make targets are: director, client-look, client-move,"
	@echo "    client-show-message, clean, director-command"

director : director-core.$(O)
	$(LD) $(GOLFLAGS) -o director director-core.$(O)
	
director-core.$(O) : $(DIRECTORREQS)
	$(GC) $(GOCFLAGS) $(DIRECTORREQS)

client-look : client-look.$(O)
	$(LD) $(GOLFLAGS) -o look client-look.$(O)
	
client-look.$(O) : $(CLIENTLOOKREQS)
	$(GC) $(GOCFLAGS) $(CLIENTLOOKREQS)

client-move : client-move.$(O)
	$(LD) $(GOLFLAGS) -o move client-move.$(O)
	
client-move.$(O) : $(CLIENTMOVEREQS)
	$(GC) $(GOCFLAGS) $(CLIENTMOVEREQS)

client-show-message : show-message.$(O)
	$(LD) $(GOLFLAGS) -o show-message show-message.$(O)

show-message.$(O) : $(CLIENTSHOWMESSAGEREQS)
	$(GC) $(GOCFLAGS) $(CLIENTSHOWMESSAGEREQS)

director-command : send-command.$(O)
	$(LD) $(GOLFLAGS) -o director-command send-command.$(O)
	
send-command.$(O) : $(DIRECTORCOMMREQS)
	$(GC) $(GOCFLAGS) $(DIRECTORCOMMREQS)

clean :
	rm -rf *.$(O)
	rm -f director look move show-message director-command
