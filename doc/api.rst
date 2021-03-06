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


All of the following urls are rooted at ``/api/v1/``. Ex: ``app/`` is at
``/api/v1/app/``.

======  ===========================  ====
verb    path                         description
======  ===========================  ====
GET     app/                         distinct active applications
GET     host/                        distinct active hosts
GET     version/                     every A & H with their most recent ``version``
GET     version/?app_id=A            ``version`` for every H running A
GET     version/?host=H              ``version`` for every A on host H
GET     version/?app_id=A&host=H     ``version`` for app A on H
GET     version/?app_id=A&ver=V      ``version`` for app A, version V
POST    update/                      accepts a ``version`` update body
GET     debug/vars                   process statistics (memory use, etc)
---     ---                          --- only when running in -danger mode
POST    _danger/clear/               clear the entire version database
---     ---                          --- items below not implemented yet
GET     /                            overview
======  ===========================  ====

Version JSON
------------

A ``version`` update has the following JSON structure::

    {
      "app" : "Application Name",
      "ver" : "1.2.3dev",
      "host" : "hostname",
      "instance" : 0,
    }

``app`` and ``ver`` are required. ``instance`` defaults to 0 and host will be inferred from the IP connection if not specified.

``app``, ``ver`` & ``host`` are limited to 50 unicode characters &
``instance`` is an integer 0-65535 (uint16).

When returned from Partisci, the following additional fields will be added::

    "app_id" : "application_name"
    "host_ip" : "10.0.0.1"
    "last_update" : 1327940599

Where ``host_ip`` is the IP address of the sending machine as seen by Partisci
and ``last_update`` is a unix epoch time stamp, rounded to the nearest second.
``app_id`` is a simplified form of ``app`` for use in referring to the application in the REST API.

All ``version/`` urls return a full version JSON object.

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

GET app/
------------------------

The response contains a distinct list of all known application names, ``app_id``,  and ``last_update`` for any version of the app from any host.

.. command-output:: curl 'http://localhost:7777/api/v1/app/' | python -m json.tool
    :shell:
    :nostderr:


GET host/
-------------------------

The response contains a distinct list of all known hosts and ``last_update`` for any version and any application.

.. command-output:: curl 'http://localhost:7777/api/v1/host/' | python -m json.tool
    :shell:
    :nostderr:


GET version/
-------------------------

The response contains every ``app_id``, ``host``, ``ver`` combination known. Only the most recent ``version`` is saved for every ``app_id``, ``host`` pair.

.. command-output:: curl 'http://localhost:7777/api/v1/version/' | python -m json.tool
    :shell:
    :nostderr:


GET version/?app=A
-------------------------

``app_id`` can be used as a parameter to filter the results.

.. command-output:: curl 'http://localhost:7777/api/v1/version/?app_id=demo_app_a' | python -m json.tool
    :shell:
    :nostderr:

GET version/?app=A&ver=V
-------------------------

``ver`` can be added to see a specific ``app`` / ``ver`` combination. Useful to see which hosts are running a version which needs updating.

.. command-output:: curl 'http://localhost:7777/api/v1/version/?app_id=demo_app_a&ver=1.0' | python -m json.tool
    :shell:
    :nostderr:

GET version/?host=H
-------------------------

``host`` can be used as a parameter to filter the results. Either alone inventory all applications:

.. command-output:: curl 'http://localhost:7777/api/v1/version/?host=host1.example.com' | python -m json.tool
    :shell:
    :nostderr:

or for a specific application:

.. command-output:: curl 'http://localhost:7777/api/v1/version/?app_id=demo_app_a&host=host1.example.com' | python -m json.tool
    :shell:
    :nostderr:


POST update/
-------------------------

Clients can POST a ``version`` update body to this url.

.. command-output:: curl 'http://localhost:7777/api/v1/update/' --data '{"instance": 0, "host": "terminal.example.com", "ver": "1.0", "app": "updatenator"}'
    :shell:
    :nostderr:

.. command-output:: curl 'http://localhost:7777/api/v1/version/?app_id=updatenator' | python -m json.tool
    :shell:
    :nostderr:


Error results
-------------------------

All error results will be returned with an appropriate HTTP status code and a JSON   document body in the following format::

    {
        "error" : "<ERROR MESSAGE>"
    }


For example, an update missing keys:
(curl -i also includes response headers in the output)

.. command-output:: curl -i 'http://localhost:7777/api/v1/update/' --data '{}'
    :shell:
    :nostderr:

Or attempting a GET on a POST only resource:

.. command-output:: curl -i 'http://localhost:7777/api/v1/update/'
    :shell:
    :nostderr:

