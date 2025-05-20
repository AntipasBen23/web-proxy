SERVICE=$1

if [ -z "$SERVICE" ]; then
  echo "❌ Please provide a service name (e.g., proxy-service, auth-service)"
  exit 1
fi

SERVICE_PATH="./services/$SERVICE"

if [ ! -d "$SERVICE_PATH" ]; then
  echo "❌ Service directory not found: $SERVICE_PATH"
  exit 1
fi

ENV_FILE="$SERVICE_PATH/config/.env"

export ENV=development
export LOG_LEVEL=debug
export CONFIG_SOURCE=env

echo "✅ Loaded base environment"

if [ -f "$ENV_FILE" ]; then
  echo "🔄 Loading environment from $ENV_FILE"
  set -a
  source "$ENV_FILE"
  set +a
else
  echo "⚠️  No .env file found at $ENV_FILE — using only base vars"
fi

echo "✅ Environment ready for: $SERVICE"
