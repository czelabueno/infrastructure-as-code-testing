output "subscriptionId" {
  value = data.azurerm_subscription.current.subscription_id
}
output "resourceName" {
  value = azurerm_storage_account.iacexample.name
}
output "resourceGroup" {
  value = azurerm_storage_account.iacexample.resource_group_name
}
output "resourceId" {
  value = azurerm_storage_account.iacexample.id
}
