.. Partisci documentation master file, created by
   sphinx-quickstart on Sun Jan 29 10:39:26 2012.
   You can adapt this file completely to your liking, but it should at least
   contain the root `toctree` directive.

Welcome to Partisci's documentation!
====================================

.. include:: ../README.rst

Details:
--------

Some limitations:
 * Applications must be modified to send updates.
 * If applications are not running, no information is updated.
 * Partisci only logs the most recent data fro each host/app/instance combination. In the future, Partisci may keep more historical data.
 * It is intended to be used within an organization to track custom software. It has no features to support publicly released software, or anything your organization can't modify.

Partisci accepts update updates via UDP and HTTP. Sending UDP updates is recommended. When UDP isn't easy or possible, use HTTP (for example: network restrictions or browser based apps). Partisci requres each application to be modified to send these updates.

Updates contain:
 * appname
 * appversion
 * hostname
 * instance number

Partisci also logs:
 * client address (as seen by Partisci)
 * time the update is recieved

Where does the name come from?
 * It is a truncated portmanteau from Partially Omniscient.

----

.. toctree::
   :maxdepth: 2


Indices and tables
==================

* :ref:`genindex`
* :ref:`search`

