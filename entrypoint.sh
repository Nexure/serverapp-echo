#!/bin/bash
glide install -v
while /bin/true; do fresh -c /fresh.conf ; sleep 5; done
