Implementation Status
=====================

Implemented
-----------

* Initial documentation
* Python update module
* UDP listener
* in-memory ``Version`` storage
* REST API

Planned
-------

* add wait_for_data(url, count) function to tests for timing coordination
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
* profile update loop
* profile ``Version`` queries

Possible
--------

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
