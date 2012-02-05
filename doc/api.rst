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


======  ==================================  ====
verb    path                                description
======  ==================================  ====
GET     /api/v1/_partisci/                  information about this partisci instance
GET     /api/v1/summary/apps/               distinct active applications
---     ---                                 --- items below not implemented yet ---
GET     /api/v1/                            overview
GET     /api/v1/version/?app=A              'version's for every H running A
GET     /api/v1/version/?app=A&host=H       'version's for app A on H
GET     /api/v1/version/?app=A&version=V    'version's for app A, version V
GET     /api/v1/version/?host=H             'version's for all A on host H
GET     /api/v1/summary/versions/           distinct active versions
GET     /api/v1/summary/versions/?app=A     distinct active versions running A
GET     /api/v1/summary/hosts/              distinct active hosts
GET     /api/v1/summary/hosts/?app=A        distinct active hosts running A
POST    /api/v1/update/                     endpoint for appliction updates
======  ==================================  ====

Version JSON
------------

Version updates have the following JSON structure::

    {
      "app" : "Application Name",
      "version" : "1.2.3dev",
      "host" : "hostname",
      "instance" : 0,
    }

app, version & host are limited to 50 unicode characters & instance is an
integer 0-65535 (uint16).

When returned from Partisci, the following additional fields will be added::

    "app_id" : "application_name"
    "host_ip" : "10.0.0.1"
    "last_update" : 1327940599

Where host_ip is the IP address of the sending machine as seen by Partisci and last_update is a unix epoch time stamp, rounded to the nearest second. app_id is a simplified form of "app" for use in referring to this application in the REST API.

Update clients
--------------

Clients should send update packets via UDP. Update packets are raw UTF8 encoded bytes containing the version JSON.

For clients which cannot use UDP, they can post the version JSON to the /version URL.

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

GET /api/v1/_partisci/
----------------------

This call returns basic information about the Partisci instance. Currently, very limited.

.. command-output:: curl -s http://localhost:7777/api/v1/_partisci/ | python -m json.tool
    :shell:


GET /api/v1/summary/app/
------------------------

The response contains a distinct list of all known application names, app_ids,
last_update for any version of the app.

.. command-output:: curl -s http://localhost:7777/api/v1/summary/app/ | python -m json.tool
    :shell:

