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

* rename? verstat? verstate? verinuse? VerInUse? verusage?
  * vstat, vstate, vuse, vinuse, vusage
* golang post about ok, err combination
* Manually create and test Windows binary
* Manually create and test linux binary
* experiment with cross-compiling windows/linux
* decide on tagging/release branch plan
* Setup job to automatically build Windows binaries on new release
* Setup job to automatically build OS/X binaries on new release
* Setup job to automatically build Linux binaries on new release
* Setup job to automatically build source distribution package on new release
* write Windows install docs
* write OS/X install docs
* write Linux install docs
* write quickstart documentation
* profile update loop (pass pointer to V instead of V?)
* profile high update load, add buffer to updates chan?
* profile ``Version`` queries
* create partisci_fuzz tool to synthesize many fake updates
* test server with MAXGOPROCS > 1

Planned - later
---------------

* add API support for returning a single AppSummary or HostSummary
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

