module "user_api" {
  source     = "git::ssh://git@github.com/FIAP-11soat-grupo-21/infra-core.git//modules/ECS-Service?ref=main"
  depends_on = [aws_lb_listener.listener]

  cluster_id            = data.terraform_remote_state.ecs.outputs.cluster_id
  ecs_security_group_id = data.terraform_remote_state.ecs.outputs.ecs_security_group_id

  cloudwatch_log_group     = data.terraform_remote_state.ecs.outputs.cloudwatch_log_group
  ecs_container_image      = var.image_name
  ecs_container_name       = var.application_name
  ecs_container_port       = var.image_port
  ecs_service_name         = var.application_name
  ecs_desired_count        = var.desired_count
  registry_credentials_arn = data.terraform_remote_state.ghcr_secret.outputs.secret_arn

  ecs_container_environment_variables = merge(var.container_environment_variables,
    {
      DB_HOST : data.terraform_remote_state.rds.outputs.db_connection,
      AWS_COGNITO_USER_POOL_ID : data.terraform_remote_state.cognito.outputs.user_pool_id
      AWS_COGNITO_USER_POOL_CLIENT_ID : data.terraform_remote_state.cognito.outputs.user_pool_client_id
      USER_PASSWORD_AUTH : data.terraform_remote_state.cognito.outputs.user_pool_client_secret
  })

  ecs_container_secrets = merge(
    var.container_secrets
    , {
      DB_PASSWORD : data.terraform_remote_state.rds.outputs.db_secret_password_arn
    }
  )

  private_subnet_ids      = data.terraform_remote_state.network_vpc.outputs.private_subnets
  task_execution_role_arn = data.terraform_remote_state.ecs.outputs.task_execution_role_arn
  task_role_policy_arns   = var.task_role_policy_arns
  alb_target_group_arn    = aws_alb_target_group.target_group.arn
  alb_security_group_id   = data.terraform_remote_state.alb.outputs.alb_security_group_id
}

module "GetUserAPIRoute" {
  source     = "git::ssh://git@github.com/FIAP-11soat-grupo-21/infra-core.git//modules/API-Gateway-Routes?ref=main"
  depends_on = [module.user_api]

  api_id       = data.terraform_remote_state.api_gateway.outputs.api_id
  alb_proxy_id = aws_apigatewayv2_integration.alb_proxy.id

  endpoints = {
    get_user = {
      route_key           = "GET /users/{id}"
      restricted          = false
      auth_integration_id = data.terraform_remote_state.auth.outputs.auth_id
    },
    get_all_users = {
      route_key           = "GET /users"
      restricted          = true
      auth_integration_id = data.terraform_remote_state.auth.outputs.auth_id
    },
    create_user = {
      route_key           = "POST /users/register"
      restricted          = false
    },
    update_user = {
      route_key           = "PUT /users/{id}"
      restricted          = true
      auth_integration_id = data.terraform_remote_state.auth.outputs.auth_id
    },
    update_user_password = {
      route_key           = "PATCH /users/{id}/password"
      restricted          = true
      auth_integration_id = data.terraform_remote_state.auth.outputs.auth_id
    },
    delete_user = {
      route_key           = "DELETE /users/{id}"
      restricted          = true
      auth_integration_id = data.terraform_remote_state.auth.outputs.auth_id
    },
    restore_user = {
      route_key           = "PATCH /users/{id}/restore"
      restricted          = true
      auth_integration_id = data.terraform_remote_state.auth.outputs.auth_id
    },
  }
}