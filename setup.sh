#!/bin/bash

rm -f clipchirp

go mod tidy
go build -o clipchirp .
chmod +x clipchirp

sudo cp clipchirp /usr/local/bin/clipchirp

echo "ClipChirp setup complete!"