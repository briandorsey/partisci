Implementation Status
=====================

Implemented
-----------

* Initial documentation
* Python update module
* UDP listener
* in-memory ``Version`` storage
* sqlite backed ``Version`` storage (-sqlite)
* REST API
* command line update client (and golang client API)
* optionally only keep recent updates (-trim)

Planned - soon
--------------

* python & golang clients - more similar API?
* document API error results
* add API support for returning a single AppSummary or HostSummary
* rename? verstat? verstate?
* golang post about ok, err combination
* test go get & convert to github import paths
* experiment with cross-compiling windows/linux
* Setup build system

  * create source distribution package with pre-built documentation
  * Windows binaries
  * OS/X binaries
  * linux binaries

* write quickstart documentation
* profile update loop (pass pointer to V instead of V?)
* profile high update load, add buffer to updates chan?
* profile ``Version`` queries
* create partisci_fuzz tool to synthesize many fake updates
* test server with MAXGOPROCS > 1

Planned - later
---------------

* add commands to ``partisci`` client for most API calls
* rename python module pypartisci --> partisci
* PYPI package for python update module.

Possible
--------

* include relative URLs to queries in API results

  * overview --> summary
  * summary --> versions/&with?parameters

* add overview API
* implement other persistence options: redis, etc
* python client: add start_update_thread()
* add ``app`` and ``host`` version summaries with counts of each version
* gzip responses when possible
* Add paging to REST results
* add ``since`` query parameter, which only returns newer ``version`` entries
* store and return app_data with each Version, allowing custom application data
* Store more historical data from updates.

  * Last update time for each app/version/machine. This would give a full version history for each machine.

