import httplib
import json
import time
import random
import socket

__version__ = "1.0"


def serialize(app, ver, host, instance):
    update = dict(
                app=app,
                ver=ver,
                host=host,
                instance=instance)
    data = json.dumps(update)
    return data


def send_udp(server, port, app, ver, host="", instance=0):
    if not host:
        try:
            host = socket.gethostname()
        except StandardError:
            pass

    data = serialize(app, ver, host, instance)
    s = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
    s.connect((server, port))
    s.send(data)
    s.close()
    return


def send_http(server, port, app, ver, host="", instance=0):
    conn = httplib.HTTPConnection(server, port)
    body = serialize(app, ver, host, instance)
    conn.request("POST", "/api/v1/update/", body)
    response = conn.getresponse()
    data = response.read()
    conn.close()
    return response.status, data


if __name__ == '__main__':
    versions = ["1.0", "2.0", "3.0"]
    hosts = ["abc", "def", "ghi"]
    instances = [0, 1, 2, 3]
    while True:
        print "%-14s Sending update" % time.time()
        send_udp('localhost', 7777, 'Python Client demo',
                    random.choice(versions),
                    random.choice(hosts),
                    random.choice(instances))
        time.sleep(2)
