**WARNING**
> > This project uses README driven development and is not yet fully implemented. See docs for current status.

# Partisci - partially omnicient version monitoring service

* Where is your software installed?
* Is version 1.2.3 still in use anywhere?

Partisci answers these questions by collecting updates from your programs and providing a REST API to access the data.

Partisci can answer these questions:

* What hosts is application A installed on?
* Which versions of application A are active?
* Is version V of application A still active anywhere?
* Which hosts are runinng version V of application A?
* When did application A last update? (from host H?)
* What applications are installed on host H?

However, Partisci *only* knows about applications which have been modified to send it updates.

Full documentation here: http://briandorsey.github.com/Partisci/