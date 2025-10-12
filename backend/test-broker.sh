#!/bin/bash
# test-brokerx.sh
# Quick script to test if BrokerX backend is running properly

echo "🔍 BrokerX Connection Test"
echo "=========================="
echo ""

# Test 1: Health Check
echo "1️⃣  Testing Health Endpoint..."
HEALTH=$(curl -s http://localhost:8080/health 2>&1)
if [ $? -eq 0 ]; then
  echo "✅ Health endpoint responding: $HEALTH"
else
  echo "❌ Health endpoint failed - Backend might not be running!"
  echo "   Start backend with: go run cmd/server/main.go"
  exit 1
fi
echo ""

# Test 2: Metrics Endpoint
echo "2️⃣  Testing Metrics Endpoint..."
METRICS=$(curl -s http://localhost:8080/metrics 2>&1)
if [ $? -eq 0 ]; then
  echo "✅ Metrics endpoint responding"
else
  echo "❌ Metrics endpoint failed"
  exit 1
fi
echo ""

# Test 3: Topics Endpoint
echo "3️⃣  Testing Topics Endpoint..."
TOPICS=$(curl -s http://localhost:8080/topics 2>&1)
if [ $? -eq 0 ]; then
  echo "✅ Topics endpoint responding: $TOPICS"
else
  echo "❌ Topics endpoint failed"
  exit 1
fi
echo ""

# Test 4: Publish Test
echo "4️⃣  Testing Publish Endpoint..."
PUBLISH=$(curl -s -X POST http://localhost:8080/publish \
  -H "Content-Type: application/json" \
  -d '{"topic":"test","sender":"test-script","payload":{"message":"hello"}}' 2>&1)
if [ $? -eq 0 ]; then
  echo "✅ Publish endpoint responding: $PUBLISH"
else
  echo "❌ Publish endpoint failed"
  exit 1
fi
echo ""

# Test 5: WebSocket Test (requires wscat)
echo "5️⃣  Testing WebSocket Connection..."
if command -v wscat &> /dev/null; then
  echo "Testing with wscat (will timeout after 3 seconds)..."
  timeout 3 wscat -c "ws://localhost:8080/subscribe?topic=test" 2>&1 | head -n 5
  if [ $? -eq 124 ]; then
    echo "✅ WebSocket connection established (timed out as expected)"
  else
    echo "⚠️  WebSocket test completed"
  fi
else
  echo "⚠️  wscat not installed. Install with: npm install -g wscat"
  echo "   Skipping WebSocket test..."
fi
echo ""

echo "=========================="
echo "✅ Backend Tests Complete!"
echo ""
echo "If all tests passed, your backend is running correctly."
echo "If WebSocket still fails in the browser, check:"
echo "  1. Browser console for CORS errors"
echo "  2. Make sure you're using http://localhost:3000 (not https)"
echo "  3. Clear browser cache and reload"