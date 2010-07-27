# dungeon-go makefile (C) 2010 Russell Bernhardt <russell.bernhardt@gmail.com>
#
# targets: director, client-look, client-move
#
include $(GOROOT)/src/Make.$(GOARCH)

GOCFLAGS = # extra flags to pass to the go compiler
GOLFLAGS = # extra flags to pass to the go linker

COMMONREQS = helpfuls.go maze.go events.go verbs.go verb_actions.go
DIRECTORREQS = director-core.go $(COMMONREQS)
CLIENTLOOKREQS = client-look.go  $(COMMONREQS)
CLIENTMOVEREQS = client-move.go $(COMMONREQS)

all :
	@echo "Valid make targets are: director, client-look, client-move"

director : director-core.$(O)
	$(LD) $(GOLFLAGS) -o director director-core.$(O)
	
director-core.$(O) : $(DIRECTORREQS)
	$(GC) $(GOCFLAGS) $(DIRECTORREQS)

client-look : client-look.$(O)
	$(LD) $(GOLFLAGS) -o client-look client-look.$(O)
	
client-look.$(O) : $(CLIENTLOOKREQS)
	$(GC) $(GOCFLAGS) $(CLIENTLOOKREQS)

client-move : client-move.$(O)
	$(LD) $(GOLFLAGS) -o client-move client-move.$(O)
	
client-move.$(O) : $(CLIENTMOVEREQS)
	$(GC) $(GOCFLAGS) $(CLIENTMOVEREQS)

clean :
	rm -rf *.$(O)
	rm -f director client-look client-move
