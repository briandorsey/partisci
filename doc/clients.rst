Clients
=======

Partisci ships with clients for several languages.

python
-------------------------

The ``pypartisci`` module implements an simple client which supports updates
via either UDP or HTTP. It does not support the full REST API. 

The most applications should send updates via UDP:

.. literalinclude:: ../clients/python/example_client_udp.py
    :language: python

If an application needs to be sure a version update is recieved, 
HTTP is also supported. If the server is unreachable for any reason,
this function will raise an exception:

.. literalinclude:: ../clients/python/example_client_http.py
    :language: python



Both functions also accept host and instance parameters:

.. code-block:: python

    send_update(SERVER, PORT, APP, VER, host, instance)

The ``PORT`` and ``instance`` variables should be integers, all others are
strings.

go
--

The golang client supports sending updates via UDP:

.. literalinclude:: ../clients/go/example_client/example_client.go
    :language: go
    :lines: 2-

