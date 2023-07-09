# Specify the provider (GCP, AWS, Azure)
provider "google" {
    credentials = "${file("service-acct.json")}"
    project = "cs-t680-demos"
    region = "us-central1"
}