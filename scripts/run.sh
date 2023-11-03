#!/bin/bash

# Check if the binary exists
if [ -f "$BINARY" ]; then
  chmod +x "./$BINARY"
  "./$BINARY"
else
  echo "Error: $BINARY does not exist or was not created."
fi
