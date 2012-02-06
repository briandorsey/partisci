API
===

Partisci provides access to the data it collects via a REST API.

The REST API also inclues an update resource for situtaions where applications cannot update via UDP.

Partisci can answer the following questions:

 * Which hosts is application A installed on?
 * Which versions of application A are active?
 * Is version V of application A still active anywhere?
 * Which hosts are runinng version V of application A?
 * When did application A last update? (from host H?)
 * What applications are installed on host H?


All of the following urls are rooted at ``/api/v1/``. Ex: ``summary/app/`` is at
``/api/v1/summary/app/``.

======  ==========================  ====
verb    path                        description
======  ==========================  ====
GET     _partisci/                  information about this partisci instance
GET     summary/app/                distinct active applications
GET     summary/hosts/              distinct active hosts
GET     version/                    every A & H with their most recent ``version``
---     ---                         --- only when running in -danger mode
POST    _danger/clear/              clear the entire version database
---     ---                         --- items below not implemented yet
GET     /                           overview
GET     version/?app=A              ``version`` for every H running A
GET     version/?app=A&host=H       ``version`` for app A on H
GET     version/?app=A&version=V    ``version`` for app A, version V
GET     version/?host=H             ``version`` for all A on host H
POST    update/                     accepts a ``version`` update body
======  ==========================  ====

Version JSON
------------

A ``version`` update has the following JSON structure::

    {
      "app" : "Application Name",
      "version" : "1.2.3dev",
      "host" : "hostname",
      "instance" : 0,
    }

``app``, ``version`` & ``host`` are limited to 50 unicode characters &
``instance`` is an integer 0-65535 (uint16).

When returned from Partisci, the following additional fields will be added::

    "app_id" : "application_name"
    "host_ip" : "10.0.0.1"
    "last_update" : 1327940599

Where ``host_ip`` is the IP address of the sending machine as seen by Partisci
and ``last_update`` is a unix epoch time stamp, rounded to the nearest second.
``app_id`` is a simplified form of ``app`` for use in referring to the application in the REST API.

All ``version/`` urls return a full version JSON object. The ``summary/`` urls return a subset of the fields.

Update clients
--------------

Clients should send update packets via UDP. Update packets are raw UTF8 encoded bytes containing the version JSON.

For clients which cannot use UDP, they can post the version JSON to the
``update/`` URL.

Update timing recommendations
-----------------------------

Clients should send a version update at every start up. Long running processes should send an update at least once every 24 hours, but not more than every hour.

Example client in Python
------------------------

<TODO>

REST API details
----------------

All response bodies are JSON objects.

The examples below assume Partisci is running on localhost, port 7777 (default).

GET _partisci/
----------------------

This call returns basic information about the Partisci instance. Currently, very limited.

.. command-output:: curl http://localhost:7777/api/v1/_partisci/ | python -m json.tool
    :shell:
    :nostderr:


GET summary/app/
------------------------

The response contains a distinct list of all known application names, ``app_id``,  and ``last_update`` for any version of the app from any host.

.. command-output:: curl http://localhost:7777/api/v1/summary/app/ | python -m json.tool
    :shell:
    :nostderr:


GET summary/host/
-------------------------

The response contains a distinct list of all known hosts and ``last_update`` for any version and any application.

.. command-output:: curl http://localhost:7777/api/v1/summary/host/ | python -m json.tool
    :shell:
    :nostderr:


GET version/
-------------------------

The response contains every ``app_id``, ``host``, ``version`` combination known. Only the most recent ``version`` is saved for every ``app_id``, ``host`` pair.

.. command-output:: curl http://localhost:7777/api/v1/version/ | python -m json.tool
    :shell:
    :nostderr:

