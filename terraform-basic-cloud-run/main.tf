provider "google" {
  project = "your-gcp-project-id" 
  region  = "us-central1"
}

resource "google_cloud_run_service" "default" {
  name     = "simple-golang-webapp"
  location = "us-central1"

  template {
    spec {
      containers {
        image = "pierinho13/simple-golang-webapp"  

        env {
          name  = "TEST_ENV_TEST"
          value = "Hello, from Cloud Run!"
        }

        ports {
          container_port = 8080
        }
      }
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }
}

output "cloud_run_url" {
  value = google_cloud_run_service.default.status[0].url
}
