Commands
========


partiscid
-------------------------

``partiscid`` is the server. For a quick test::

    $ partiscid -v
    2012/02/14 08:10:02 Starting.
    2012/02/14 08:10:02 listening on: 0.0.0.0:7777
    ...

This will start the server, listening on port 7777, using memory
to store ``version`` updates.

``partiscid`` also supports the following flags:

.. command-output:: partiscid -help
    :shell:

