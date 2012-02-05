.. Partisci documentation master file, created by
   sphinx-quickstart on Sun Jan 29 10:39:26 2012.
   You can adapt this file completely to your liking, but it should at least
   contain the root `toctree` directive.

Welcome to Partisci's documentation!
====================================

.. WARNING::
   This project is using readme driven development and has not been fully implemented yet.

Where is your software installed?  Is version X still in use anywhere?

Partisci answers these questions by collecting updates from your programs and providing a REST API to access the data.

Partisci can answer these questions:
 * What hosts is application X installed on?
 * Which versions of application X are active?
 * Is version Y of application X still active anywhere?
 * Which hosts are runinng version Y of application X?
 * When did application X last update? (from host Z?)
 * What applications are installed on host Z?

However, Partisci *only* knows about applications which have been modified to send it updates.


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


Contents
--------

.. toctree::
   :maxdepth: 2

   api
   testing
   status

Indices and tables
==================

* :ref:`genindex`
* :ref:`search`

