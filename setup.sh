#!/bin/bash

rm -f clipchirp

go mod tidy
go build -o clipchirp main.go
chmod +x clipchirp

sudo cp clipchirp /usr/local/bin/clipchirp

echo "ClipChirp setup complete!"