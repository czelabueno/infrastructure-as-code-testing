terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "~> 1.44.0"
    }
  }
}

resource "random_integer" "iacexample" {
  min = 1
  max = 10
}

resource "azurerm_resource_group" "iacexample" {
  name     = "iacunittesting"
  location = "eastus2"
  tags = local.tags
}

resource "azurerm_storage_account" "iacexample" {
  name                     = "iacstorage${random_integer.iacexample.result}"
  resource_group_name      = azurerm_resource_group.iacexample.name
  location                 = azurerm_resource_group.iacexample.location
  account_tier             = var.account_tier
  account_replication_type = var.account_replication_type
  tags  = local.tags
}

resource "azurerm_storage_container" "iacexample" {
  name                  = "contentexample"
  storage_account_name  = azurerm_storage_account.iacexample.name
  container_access_type = "private"
}