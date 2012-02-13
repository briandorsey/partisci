import json
import os
import pprint
import subprocess
import sys
import time
import urlparse

sys.path.insert(0, os.path.join(os.path.dirname(__file__), "../clients"))
import pypartisci

import requests

server = '127.0.0.1'
port = 7788
endpoint = "http://127.0.0.1:%s/api/v1/" 

class TestPartisci:
    def setup_class(self):
        self.port = port

    def setup_method(self, method):
        self.port += 1
        self.server = subprocess.Popen(["partisci",
                                        "--port=%s" % self.port,
                                        "--listenip=%s" % server,
                                        "--danger"])
        url = urlparse.urljoin(endpoint % self.port, "_partisci/")
        for i in range(100):
            try:
                response = requests.get(url)
                return
            except requests.ConnectionError:
                time.sleep(.01)
                continue
        raise StandardError("partisci never started: %s" % (
                    response.text))

    def teardown_method(self, method):
        if self.server:
            self.server.kill()

    def send_basic_updates(self, prefix):
        apps = ["_zz_%s_app%s" % (prefix, str(i)) for i in range(5)]
        hosts = ["_zz_%s_host%s"  % (prefix, str(i)) for i in range(5)]
        versions = ["1", "2", "3", "2", "1"] * 50
        print "apps:", apps
        print "hosts:", hosts
        def do():
            for app in apps:
                for i, host in enumerate(hosts):
                    ver = versions[i]
                    pypartisci.send_update(server, self.port, app, ver, host)
        do()
        # then again to make sure the server updates, not duplicates
        do()
        return apps, hosts

    def wait_for_data(self, url, count):
        """Wait for enough data, or raise error"""
        print "Waiting for len(data) >= %s at: %s" % (count, url)
        for i in range(50):
            response = requests.get(url)
            info = json.loads(response.content)
            if info and "data" in info:
                if len(info["data"]) >= count:
                    return info
            time.sleep(.1)
        raise StandardError("Never got enough data. info: %s" % (
                    pprint.pformat(info)))

    def test_get_server_info(self):
        url = urlparse.urljoin(endpoint % self.port, "_partisci/")
        print url
        response = requests.get(url)
        print response
        print response.content
        info = json.loads(response.content)
        print info
        assert "version" in info

    def test_summary_app(self):
        url = urlparse.urljoin(endpoint % self.port, "app/")
        print url
        response = requests.get(url)
        print response
        print response.content
        info = json.loads(response.content)
        print info
        # empty result should still be a list.
        assert list() == info["data"]

        apps, hosts = self.send_basic_updates("app")

        info = self.wait_for_data(url, len(apps))

        assert "data" in info
        for v in info["data"]:
            print v
            assert "app" in v
            assert "app_id" in v
            assert "host_count" in v
            assert v["host_count"] > 0
            assert "last_update" in v
            assert "ver" not in v
            assert "host" not in v
            assert "host_ip" not in v
            assert "instance" not in v

        names = set(v["app"] for v in info["data"])
        for app in apps:
            assert app in names

    def test_summary_host(self):
        url = urlparse.urljoin(endpoint % self.port, "host/")
        print url
        response = requests.get(url)
        print response
        print response.content
        info = json.loads(response.content)
        print info
        # empty result should still be a list.
        assert list() == info["data"]

        apps, hosts = self.send_basic_updates("host")

        info = self.wait_for_data(url, len(apps))

        assert "data" in info
        for v in info["data"]:
            print v
            assert "host" in v
            assert "last_update" in v
            assert "app" not in v
            assert "app_id" not in v
            assert "ver" not in v
            assert "host_ip" not in v
            assert "instance" not in v

        names = set(v["host"] for v in info["data"])
        for host in hosts:
            assert host in names

    def test_version(self):
        url = urlparse.urljoin(endpoint % self.port, "version/")
        print url
        response = requests.get(url)
        print response
        print response.content
        info = json.loads(response.content)
        print info
        # empty result should still be a list.
        assert list() == info["data"]

        apps, hosts = self.send_basic_updates("version")

        info = self.wait_for_data(url, len(apps) * len(hosts))

        assert "data" in info
        for v in info["data"]:
            print v
            assert "host" in v
            assert "last_update" in v
            assert "app" in v
            assert "app_id" in v
            assert "ver" in v
            assert "host_ip" in v
            #assert "instance" in v

        app_names = set(v["app"] for v in info["data"])
        host_names = set(v["host"] for v in info["data"])
        print app_names
        print host_names
        for app in apps:
            assert app in app_names
        for host in hosts:
            assert host in host_names

    def test_version_app(self):
        apps, hosts = self.send_basic_updates("version_app")
        url = urlparse.urljoin(endpoint % self.port, "app/")
        info = self.wait_for_data(url, len(apps))

        # pick the first app_id
        app_id = info["data"][0]["app_id"]
        print "Requesting app_id:", app_id

        url = urlparse.urljoin(endpoint % self.port, "version/?app_id=%s" % app_id)
        print url
        response = requests.get(url)
        info = json.loads(response.content)

        for v in info["data"]:
            print v
            assert v["app_id"] == app_id

    def test_version_host(self):
        apps, hosts = self.send_basic_updates("version_host")
        url = urlparse.urljoin(endpoint % self.port, "host/")
        info = self.wait_for_data(url, len(hosts))
        
        # pick the first host
        host = info["data"][0]["host"]
        print "Requesting host:", host

        url = urlparse.urljoin(endpoint % self.port, "version/?host=%s" % host)
        print url
        response = requests.get(url)
        info = json.loads(response.content)

        for v in info["data"]:
            print v
            assert v["host"] == host

    def test_version_app_host(self):
        apps, hosts = self.send_basic_updates("version_app_host")

        # pick the first app_id
        url = urlparse.urljoin(endpoint % self.port, "app/")
        info = self.wait_for_data(url, len(apps))
        app_id = info["data"][0]["app_id"]
        print "Requesting app_id:", app_id

        # pick the first host
        url = urlparse.urljoin(endpoint % self.port, "host/")
        info = self.wait_for_data(url, len(hosts))
        host = info["data"][0]["host"]
        print "Requesting host:", host

        url = urlparse.urljoin(endpoint % self.port,
                               "version/?app_id=%s&host=%s" % (app_id, host))
        print url
        response = requests.get(url)
        info = json.loads(response.content)

        for v in info["data"]:
            print v
            assert v["host"] == host
            assert v["app_id"] == app_id

    def test_version_app_version(self):
        apps, hosts = self.send_basic_updates("version_app_version")

        # pick the first app_id
        url = urlparse.urljoin(endpoint % self.port, "app/")
        info = self.wait_for_data(url, len(apps))
        app_id = info["data"][0]["app_id"]
        print "Requesting app_id:", app_id

        ver = "1"
        url = urlparse.urljoin(endpoint % self.port,
                               "version/?app_id=%s&ver=%s" % (app_id, ver))
        print url
        response = requests.get(url)
        info = json.loads(response.content)

        for v in info["data"]:
            print v
            assert v["ver"] == ver
            assert v["app_id"] == app_id

    def test_update(self):
        app = "http_update"

        url = urlparse.urljoin(endpoint % self.port, "app/")
        response = requests.get(url)
        info = json.loads(response.content)
        assert len(info["data"]) == 0

        code, data = pypartisci.send_update_http(server, self.port, app, "1.0")
        assert code == 200
        
        response = requests.get(url)
        info = json.loads(response.content)
        data = info["data"]
        assert len(data) == 1
        print data
        for v in data:
            assert v["app"] == app
