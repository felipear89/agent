#!/bin/bash

# Generate private key
openssl genrsa -out private.pem 2048

# Generate public key
openssl rsa -in private.pem -outform PEM -pubout -out public.pem

# Output the keys as environment variables
echo "JWT_PRIVATE_KEY='$(cat private.pem)'"
echo "JWT_PUBLIC_KEY='$(cat public.pem)'"

# Clean up
rm private.pem public.pem
