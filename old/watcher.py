#!/usr/bin/python

try:
    from AppKit import NSWorkspace
except ImportError:
    print "Can't import AppKit -- maybe you're running python from brew?"
    print "Try running with Apple's /usr/bin/python instead."
    exit(1)

from datetime import datetime
from time import sleep
import json
import sys


def chchchchanges():
    last_active_name = None
    while True:
        active_app = NSWorkspace.sharedWorkspace().activeApplication()
        if active_app['NSApplicationName'] != last_active_name:
            last_active_name = active_app['NSApplicationName']
            print json.dumps({
                'Time': datetime.now().strftime('%Y-%m-%d %H:%M:%S'),
                'Name': active_app['NSApplicationName'],
                'Path': active_app['NSApplicationPath'],
                'BundleID': active_app['NSApplicationBundleIdentifier']
            })
            sys.stdout.flush()
        sleep(.5)

def what_doin():
    active_app = NSWorkspace.sharedWorkspace().activeApplication()
    return  {
        'Name': active_app['NSApplicationName'],
        'Path': active_app['NSApplicationPath'],
        'BundleID': active_app['NSApplicationBundleIdentifier']}


# print json.dumps(what_doin())

chchchchanges()
