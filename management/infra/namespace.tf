resource "kubernetes_namespace" "platform" {
  count = var.management_infra_minimal ? 0 : 1

  metadata {
    name = "platform"
  }
}
