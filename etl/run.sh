#!/bin/bash
rm /var/tmp/etl-info.log 2>/dev/null; rm /var/tmp/etl-error.log 2>/dev/null; rm /var/tmp/db.log 2>/dev/null; ./etl 1>/var/tmp/etl-info.log 2>/var/tmp/etl-error.log
