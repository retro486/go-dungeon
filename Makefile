# dungeon-go makefile (C) 2010 Russell Bernhardt <russell.bernhardt@gmail.com>
#
# targets: director, client-look, client-move
#
GOC = 6g # the go compiler to use; typically 6g on 64-bit systems and 8g on 32-bit systems.
GOL = 6l # the go linker to use; typically 6l on 64-bit systems and 8l on 32-bit systems.
GOCFLAGS = # extra flags to pass to the go compiler
GOLFLAGS = # extra flags to pass to the go linker

COMMONREQS = helpfuls.go maze.go events.go verbs.go verb_actions.go
DIRECTORREQS = director-core.go $(COMMONREQS)
CLIENTLOOKREQS = client-look.go  $(COMMONREQS)
CLIENTMOVEREQS = client-move.go $(COMMONREQS)

all :
	@echo "Valid make targets are: director, client-look, client-move"

director : director-core.6
	$(GOL) $(GOLFLAGS) -o director director-core.6
	
director-core.6 : $(DIRECTORREQS)
	$(GOC) $(GOCFLAGS) $(DIRECTORREQS)

client-look : client-look.6
	$(GOL) $(GOLFLAGS) -o client-look client-look.6
	
client-look.6 : $(CLIENTLOOKREQS)
	$(GOC) $(GOCFLAGS) $(CLIENTLOOKREQS)

client-move : client-move.6
	$(GOL) $(GOLFLAGS) -o client-move client-move.6
	
client-move.6 : $(CLIENTMOVEREQS)
	$(GOC) $(GOCFLAGS) $(CLIENTMOVEREQS)

clean :
	rm -rf *.6
	rm -f director client-look client-move
