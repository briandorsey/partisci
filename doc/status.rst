Implementation Status
=====================

Implemented
-----------

 * Initial documentation
 * example Python update module
 * initial UDP listener
 * implement a benchmark test for update parsing

Planned
-------

 * barebones REST API with tests
 * python client: add start_update_thread(), docs
 * in-memory version storage and interface
 * write quickstart documentation
 * fully implement documented REST API
 * Persistent store for the version data
 * Setup build system

   * create source distribution package with pre-built documentation
   * post pre-built documentation online
   * Windows binaries
   * OS/X binaries
   * linux binaries

 * write golang update client
 * Add paging to REST results

Possible
--------

* Store more historical data from updates.

  * Last update time for each app/version/machine. This would give a full version history for each machine.

* Store application specific state information.
