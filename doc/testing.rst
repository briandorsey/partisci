Testing
=======

Requirements & running the tests
--------------------------------

Partisci is written in Go (golang.org).

Unit and benchmark tests are written using the Go `testing` package and can be run from the root of the Partisci repository::

  go test ./...

The REST API test suite is written in Python (python.org). Requirements:

* partiscid built and available on your path
* Python 2.7
* py.test (pytest.org)
* requests (python-requests.org)

Open a shell to the test directory and run::

  $ py.test


