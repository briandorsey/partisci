import json
import time
import socket

__version__ = "0.1demo"

def update_template():
    return dict(
        name="",
        version="",
        host="",
        instance=0)


def send_update(server, port, name, version, host="", instance=0):
    if not host:
        try:
            host = socket.gethostname()
        except StandardError:
            pass

    update = dict(
                name=name,
                version=version,
                host=host,
                instance=instance)
    data = json.dumps(update)
    s = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
    s.connect((server, port))
    s.send(data)
    s.close()
    return

if __name__ == '__main__':
    while True:
        print "%-14s Sending update" % time.time()
        send_update('localhost', 7777, 'Python Client demo', __version__)
        time.sleep(3)

