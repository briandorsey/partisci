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
GET     /                           overview
GET     /application/               distinct list of applications
GET     /application/X              most recent 'version' for every Z running X
GET     /application/X/version      distinct active versions
GET     /application/X/version/Y    only 'version's running version Y
GET     /host                       distinct active hosts
GET     /host/Z                     all active 'version's for all X on host Z
POST    /version                    endpoint for appliction updates
======  ==========================  ====

Version structure
-----------------

Version updates have the following JSON structure::

    {
      "appname" : "application_name",
      "appversion" : "1.2.3dev",
      "host" : "hostname",
      "instance" : 0,
    }

When returned from Partisci, the following additional fields will be added::

    "host_ip" : "10.0.0.1"
    "last_update" : 1327940599



