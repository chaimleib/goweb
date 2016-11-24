#!/bin/bash

fswatch -0 src |
    xargs -0 -n1 'make qtest all'

