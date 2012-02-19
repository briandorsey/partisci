
Partisci's documentation
========================

.. WARNING::
   This project is using readme driven development and has not been fully implemented yet.

Where is your software installed?  Is version X still in use anywhere?

Partisci answers these questions by collecting updates from your programs and providing a REST API to access the data.

Partisci can answer these questions:

* What hosts is application A installed on?
* Which versions of application A are active?
* Is version V of application A still active anywhere?
* Which hosts are runinng version V of application A?
* When did application A last update? (from host H?)
* What applications are installed on host H?

However, Partisci *only* knows about applications which send updates. Each application needs a small change to send these updates.


Details
=======

Some limitations:

* Applications must be modified to send updates.
* If applications are not running, no information is updated.
* Partisci only logs the most recent data fro each host/app/instance combination. In the future, Partisci may keep more historical data.
* It is intended to be used within an organization to track custom software. It has no features to support publicly released software, or anything your organization can't modify.

Partisci accepts update updates via UDP and HTTP. Sending UDP updates is recommended. When UDP isn't easy or possible, use HTTP (for example: network restrictions or browser based apps). Partisci requres each application to be modified to send these updates.

Updates contain:

* application name
* version
* hostname
* instance number

Partisci also logs:

* client address (as seen by Partisci)
* time the update is recieved

Where does the name come from?

* It is a truncated portmanteau from Partially Omniscient. Within its limited domain, Partisci strives to be omniscient.



.. toctree::
   :maxdepth: 2

   commands
   api
   testing
   status

Indices and tables
==================

* :ref:`genindex`
* :ref:`search`

