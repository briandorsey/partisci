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
    :returncode: 2

The ``-trim`` option will remove references to any instances which have not
sent a version update within the number of seconds specified by ``-trim``.


partisci
-------------------------

``partisci`` is a command line utility to communicate with ``partiscid``.

The main command is ``update``, which sends a single custom update message:

.. command-output:: partisci update "Demo App A" 1.0 host1.example.com 0
    :shell:


``partisci`` also supports the following flags:

.. command-output:: partisci -help
    :shell:
    :returncode: 2

