import json
import os
import subprocess
import sys
import urlparse

sys.path.insert(0, os.path.join(os.path.dirname(__file__), "../clients"))
import pypartisci

import requests

server, port = "localhost", 7788
endpoint = "http://%s:%s/api/v1/" % (server, port)

class TestPartisci:
    def setup_class(self):
        self.server = subprocess.Popen(["partisci",
                                        "--port=%s" % port,
                                        "--danger"])

    def teardown_class(self):
        self.server.kill()

    def setup_method(self, method):
        clear_url = urlparse.urljoin(endpoint, "_danger/clear/")
        response = requests.post(clear_url)
        print response.content
        assert response.ok

    def send_basic_updates(self):
        apps = ["_zz_app" + str(i) for i in range(5)]
        hosts = ["_zz_host" + str(i) for i in range(5)]
        print "apps:", apps
        print "hosts:", hosts
        for app in apps:
            for host in hosts:
                pypartisci.send_update(server, port, app, "ver", host)
        return apps, hosts

    def test_get_server_info(self):
        url = urlparse.urljoin(endpoint, "_partisci/")
        print url
        response = requests.get(url)
        print response
        print response.content
        info = json.loads(response.content)
        print info
        assert "version" in info

    def test_summary_app(self):
        url = urlparse.urljoin(endpoint, "summary/app/")
        print url
        response = requests.get(url)
        print response
        print response.content
        info = json.loads(response.content)
        print info
        # empty result should still be a list.
        assert list() == info["data"]

        apps, hosts = self.send_basic_updates()

        response = requests.get(url)
        print response
        info = json.loads(response.content)

        assert "data" in info
        for v in info["data"]:
            print v
            assert "app" in v
            assert "app_id" in v
            assert "last_update" in v
            assert "version" not in v
            assert "host" not in v
            assert "host_ip" not in v
            assert "instance" not in v

        names = set(v["app"] for v in info["data"])
        for app in apps:
            assert app in names

    def test_summary_host(self):
        url = urlparse.urljoin(endpoint, "summary/host/")
        print url
        response = requests.get(url)
        print response
        print response.content
        info = json.loads(response.content)
        print info
        # empty result should still be a list.
        assert list() == info["data"]

        apps, hosts = self.send_basic_updates()

        response = requests.get(url)
        print response
        info = json.loads(response.content)

        assert "data" in info
        for v in info["data"]:
            print v
            assert "host" in v
            assert "last_update" in v
            assert "app" not in v
            assert "app_id" not in v
            assert "version" not in v
            assert "host_ip" not in v
            assert "instance" not in v

        names = set(v["host"] for v in info["data"])
        for host in hosts:
            assert host in names

    def test_version(self):
        url = urlparse.urljoin(endpoint, "version/")
        print url
        response = requests.get(url)
        print response
        print response.content
        info = json.loads(response.content)
        print info
        # empty result should still be a list.
        assert list() == info["data"]

        apps, hosts = self.send_basic_updates()

        response = requests.get(url)
        print response
        info = json.loads(response.content)

        assert "data" in info
        for v in info["data"]:
            print v
            assert "host" in v
            assert "last_update" in v
            assert "app" in v
            assert "app_id" in v
            assert "version" in v
            assert "host_ip" in v
            #assert "instance" in v

        app_names = set(v["app"] for v in info["data"])
        host_names = set(v["host"] for v in info["data"])
        for app in apps:
            assert app in app_names
        for host in hosts:
            assert host in host_names

    def test_version_app(self):
        apps, hosts = self.send_basic_updates()
        url = urlparse.urljoin(endpoint, "summary/app/")

        response = requests.get(url)
        info = json.loads(response.content)

        # pick the first app_id
        app_id = info["data"][0]["app_id"]
        print "Requesting app_id:", app_id

        url = urlparse.urljoin(endpoint, "version/?app=%s" % app_id)
        print url
        response = requests.get(url)
        info = json.loads(response.content)

        for v in info["data"]:
            print v
            assert v["app_id"] == app_id

    def test_version_host(self):
        apps, hosts = self.send_basic_updates()
        url = urlparse.urljoin(endpoint, "summary/host/")

        response = requests.get(url)
        info = json.loads(response.content)

        # pick the first host
        host = info["data"][0]["host"]
        print "Requesting host:", host

        url = urlparse.urljoin(endpoint, "version/?host=%s" % host)
        print url
        response = requests.get(url)
        info = json.loads(response.content)

        for v in info["data"]:
            print v
            assert v["host"] == host

    def test_version_app_host(self):
        apps, hosts = self.send_basic_updates()

        # pick the first app_id
        url = urlparse.urljoin(endpoint, "summary/app/")
        response = requests.get(url)
        info = json.loads(response.content)
        app_id = info["data"][0]["app_id"]
        print "Requesting app_id:", app_id

        # pick the first host
        url = urlparse.urljoin(endpoint, "summary/host/")
        response = requests.get(url)
        info = json.loads(response.content)
        host = info["data"][0]["host"]
        print "Requesting host:", host

        url = urlparse.urljoin(endpoint,
                               "version/?app=%s&host=%s" % (app_id, host))
        print url
        response = requests.get(url)
        info = json.loads(response.content)

        for v in info["data"]:
            print v
            assert v["host"] == host
            assert v["app_id"] == app_id
