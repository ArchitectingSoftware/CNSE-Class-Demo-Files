

resource "google_service_account" "vm_service_account" {
  account_id   = "cst680-vm-service-account"
  display_name = "CS-T680 Service Account"
}

resource "google_project_iam_member" "vm_service_account_roles" {
  project = var.project_id
  role    = "roles/compute.instanceAdmin"
  member  = "serviceAccount:${google_service_account.vm_service_account.email}"
}

resource "google_project_iam_member" "vm_service_account_ssh" {
  project = var.project_id
  role    = "roles/compute.osAdminLogin"
  member  = "serviceAccount:${google_service_account.vm_service_account.email}"
}

resource "google_project_iam_member" "vm_service_account_storage_viewer" {
  project = var.project_id
  role    = "roles/storage.objectViewer"
  member  = "serviceAccount:${google_service_account.vm_service_account.email}"
}

# Note this gets bound in terraform.tfvars
variable "project_id" {
  description = "Google Cloud project ID"
  default     = "not-set"
}

# Note this gets bound in terraform.tfvars
variable "svc_acct_email" {
  description = "Google Cloud Service Acct Email for Terraform"
  default     = "not-set"
}

