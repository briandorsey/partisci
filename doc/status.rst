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

 * write quickstart documentation
 * fully implement documented REST API
 * python client: add start_update_thread(), docs
 * implement version timeout and config (only active versions kept)

   * and/or - make this a query parameter?

 * implement and test multiple instance support
 * add relative URLs queries in API results
 * write golang update client
 * create partisci_fuzz tool to synthesize many fake updates
 * Persistent store for the version data
 * Setup build system

   * create source distribution package with pre-built documentation
   * post pre-built documentation online
   * Windows binaries
   * OS/X binaries
   * linux binaries

 * Add paging to REST results

Possible
--------

* Store more historical data from updates.

  * Last update time for each app/version/machine. This would give a full version history for each machine.

* Store application specific state information.
