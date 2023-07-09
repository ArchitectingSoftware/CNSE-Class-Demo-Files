resource "google_storage_bucket" "my_bucket" {
  name = "csts680-class-bucket"
  location = "us-central1"
}

resource "google_storage_bucket_object" "amd_api_object" {
  name   = "todo_linux_amd64"
  bucket = google_storage_bucket.my_bucket.name
  source = "../todo-api/todo-linux-amd64"
}

resource "google_storage_bucket_object" "arm_api_object" {
  name   = "todo_linux_arm64"
  bucket = google_storage_bucket.my_bucket.name
  source = "../todo-api/todo-linux-arm64"
}