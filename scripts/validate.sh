#!/bin/bash
echo "🔍 Validating app..."
curl -f http://localhost:5000 || exit 1
