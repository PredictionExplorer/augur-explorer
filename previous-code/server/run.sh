#!/bin/bash
rm /var/tmp/backend-info.log 2>/dev/null; rm /var/tmp/backend-error.log 2>/dev/null; rm /var/tmp/backend-db.log 2>/dev/null; ./server 1>/var/tmp/backend-info.log 2>/var/tmp/backend-error.log

