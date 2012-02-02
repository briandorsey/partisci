API
===

Partisci provides access to the data it collects via a REST API.

The REST API also inclues an update resource for situtaions where applications cannot update via UDP.

Partisci can answer the following questions:

 * What hosts is application X installed on?
 * Which versions of application X are active?
 * Is version Y of application X still active anywhere?
 * Which hosts are runinng version Y of application X?
 * When did application X last update? (from host Z?)
 * What applications are installed on host Z?


======  ==========================  ====
verb    path                        description
======  ==========================  ====
GET     /api/v1/_partisci           information about this partisci instance
---     ---                         --- Below items not implemented yet ---
GET     /api/v1/                    overview
GET     /api/v1/app/                distinct list of applications
GET     /api/v1/app/X               current 'version' for every Z running X
GET     /api/v1/app/X/version       distinct active versions
GET     /api/v1/app/X/version/Y     only 'version's running version Y
GET     /api/v1/app/X/host          distinct active hosts running X
GET     /api/v1/app/X/host/Z        current app 'version' for Z
GET     /api/v1/host                distinct active hosts
GET     /api/v1/host/Z              all active 'version's for all X on host Z
POST    /api/v1/version             endpoint for appliction updates
======  ==========================  ====

Version JSON
------------

Version updates have the following JSON structure::

    {
      "name" : "Application Name",
      "version" : "1.2.3dev",
      "host" : "hostname",
      "instance" : 0,
    }

name, version & host are limited to 50 unicode characters & instance is an integer <= 65535 (uint16).

TODO: name format: underscores, no spaces, etc? Or accept anything, covert it, and use the simplified form on the urls, as an id? Yes, this.

When returned from Partisci, the following additional fields will be added::

    "app_id" : "application_name"
    "host_ip" : "10.0.0.1"
    "last_update" : 1327940599

Where host_ip is the IP address of the sending machine as seen by Partisci and last_update is a unix epoch time stamp, rounded to the nearest second. app_id is a simplified form of "name" for use in referring to this application in the REST API.

Update clients
--------------

Clients should send update packets via UDP. Update packets are raw UTF8 encoded bytes containing the version JSON.

For clients which cannot use UDP, they can post the version JSON to the /version URL.

Example client in Python
------------------------

<TODO>

