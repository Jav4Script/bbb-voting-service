#!/bin/sh

sleep 10

rabbitmqadmin declare queue name="$VOTE_QUEUE" durable=true