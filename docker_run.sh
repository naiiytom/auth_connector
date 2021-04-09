#!/bin/bash

docker run -itd --env-file="env-file" -p 5000:5000 auth-middleware