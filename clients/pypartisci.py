import json
import time
import socket

__version__ = "1.0"

def send_update(server, port, app, ver, host="", instance=0):
    if not host:
        try:
            host = socket.gethostname()
        except StandardError:
            pass

    update = dict(
                app=app,
                ver=ver,
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

