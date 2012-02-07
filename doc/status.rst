Implementation Status
=====================

Implemented
-----------

* Initial documentation
* example Python update module
* initial UDP listener
* implement a benchmark test for update parsing
* barebones REST API with tests
* in-memory version storage and interface

Planned
-------

* fully implement the documented REST API
* add ``count`` field to summary results
* support both / terminated and not urls
* add ``app`` and ``host`` version summaries with counts of each version
* include relative URLs to queries in API results

  * overview --> summary
  * summary --> versions/&with?parameters

* implement version timeout and config (only active versions kept)
* implement and test multiple instance support
* write golang update client
* python client: add start_update_thread(), docs
* create partisci_fuzz tool to synthesize many fake updates
* create a persistent store for the version data
* Setup build system

  * create source distribution package with pre-built documentation
  * post pre-built documentation online
  * Windows binaries
  * OS/X binaries
  * linux binaries

* write quickstart documentation

Possible
--------

* gzip responses when possible
* Add paging to REST results
* add ``since`` query parameter, which only returns newer ``version`` entries
* Store more historical data from updates.

  * Last update time for each app/version/machine. This would give a full version history for each machine.

* Store application specific state information.
