Implementation Status
=====================

Implemented
-----------

* Initial documentation
* Python update module
* UDP listener
* in-memory ``Version`` storage
* REST API
* command line update client (and golang client API)

Planned
-------

* test updates with missing app, ver, host values (empty strings)
* create partisci_fuzz tool to synthesize many fake updates
* test server with MAXGOPROCS > 1
* python client: add docs
* create a persistent store for the version data

  * goleveldb
  * redis

* write quickstart documentation
* implement version timeout and config (only active versions kept)
* test go get & convert to github import paths
* Setup build system

  * create source distribution package with pre-built documentation
  * Windows binaries
  * OS/X binaries
  * linux binaries

* tests - switch to REST update? allows single server process?
* profile update loop (pass pointer to V instead of V?)
* profile high update load, add buffer to updates chan?
* profile ``Version`` queries

Possible
--------

* python client: add start_update_thread()
* include relative URLs to queries in API results

  * overview --> summary
  * summary --> versions/&with?parameters

* add overview API
* add ``count`` field to host summary results?
* add ``app`` and ``host`` version summaries with counts of each version
* gzip responses when possible
* Add paging to REST results
* add ``since`` query parameter, which only returns newer ``version`` entries
* store and return app_data with each Version, allowing custom data
* Store more historical data from updates.

  * Last update time for each app/version/machine. This would give a full version history for each machine.

* Store application specific state information.
