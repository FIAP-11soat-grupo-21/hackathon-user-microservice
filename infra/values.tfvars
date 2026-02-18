application_name = "user-api"
image_name       = "GHCR_IMAGE_TAG"
image_port       = 8083
app_path_pattern = ["/users*"]

# =======================================================
# Configurações do ECS Service
# =======================================================
container_environment_variables = {
  GO_ENV : "production"
  API_PORT : "8083"
  API_HOST : "0.0.0.0"

  AWS_REGION : "us-east-2"

  DB_RUN_MIGRATIONS : "true"
  DB_NAME : "payment_db"
  DB_PORT : "5432"
}

container_secrets = {}
health_check_path = "/health"
task_role_policy_arns = [
  "arn:aws:iam::aws:policy/AmazonCognitoPowerUser",
  "arn:aws:iam::aws:policy/AmazonRDSFullAccess",
]
alb_is_internal = true

# =======================================================
# Configurações do API Gateaway
# =======================================================
# API Gateway
apigw_integration_type       = "HTTP_PROXY"
apigw_integration_method     = "ANY"
apigw_payload_format_version = "1.0"
apigw_connection_type        = "VPC_LINK"

authorization_name = "CognitoAuthorizer"