resource "google_compute_network" "vpc_network" {
  name                    = "ts680-network"
  auto_create_subnetworks = false
  mtu                     = 1460
}

resource "google_compute_subnetwork" "default" {
  name          = "my-custom-subnet"
  ip_cidr_range = "10.0.1.0/24"
  region        = "us-central1"
  network       = google_compute_network.vpc_network.id
}

# Create a single Compute Engine instance
resource "google_compute_instance" "default" {
  name         = "api-vm"
  # machine_type = "f1-micro"
  machine_type = "n1-standard-1"
  zone         = "us-central1-a"
  tags         = ["ssh","api-server"]

 service_account {
  email  = "${google_service_account.vm_service_account.email}"
  scopes = ["https://www.googleapis.com/auth/cloud-platform"]
}


  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }

  # Install Python and Go
  metadata_startup_script = <<SCRIPT
    sudo apt-get update
    sudo apt-get install -yq build-essential python3-pip rsync git snapd
    sudo snap install --classic --channel=1.20/stable go
    gsutil cp gs://${google_storage_bucket.my_bucket.name}/${google_storage_bucket_object.arm_api_object.name} /usr/local/bin/${google_storage_bucket_object.arm_api_object.name} > /tmp/gcp_arm_copy.log 2>&1
    chmod +x /usr/local/bin/${google_storage_bucket_object.arm_api_object.name} 
    gsutil cp gs://${google_storage_bucket.my_bucket.name}/${google_storage_bucket_object.amd_api_object.name} /usr/local/bin/${google_storage_bucket_object.amd_api_object.name} > /tmp/gcp_amd_copy.log 2>&1
    chmod +x /usr/local/bin/${google_storage_bucket_object.amd_api_object.name} 
  SCRIPT
  network_interface {
    subnetwork = google_compute_subnetwork.default.id

    access_config {
      # Include this section to give the VM an external IP address
    }
  }

 
}

resource "google_compute_firewall" "ssh" {
  name = "allow-ssh"
  allow {
    ports    = ["22"]
    protocol = "tcp"
  }
  direction     = "INGRESS"
  network       = google_compute_network.vpc_network.id
  priority      = 1000
  source_ranges = ["0.0.0.0/0"]
  target_tags   = ["ssh"]
}

resource "google_compute_firewall" "my-api" {
  name    = "api-firewall"
  network = google_compute_network.vpc_network.id

  allow {
    protocol = "tcp"
    ports    = ["1080","80","22"]
  }
  source_ranges = ["0.0.0.0/0"]
}

resource "google_compute_project_metadata_item" "ssh-keys" {
  key   = "ssh-keys"
  value = "${file("/Users/bsm23/.ssh/id_rsa_gcp.pub")}"
}





output "ip" {
  value = "${google_compute_instance.default.network_interface.0.access_config.0.nat_ip}"
}