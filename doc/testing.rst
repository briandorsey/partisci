Testing
=======

Requirements & running the tests
--------------------------------

Partisci itself is written in Go (golang.org). The (whitebox) test suite is written in Python. To run the tests, you'll need Python 2.7x and the following modules:

 * py.test (pytest.org)
 * requests (python-requests.org)

Partisci should be running on localhost. Then open a shell to the test directory and run::

  $ py.test


