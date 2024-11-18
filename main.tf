terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "4.51.0"
    }
  }
}

provider "google" {
  project = var.project_id
  region  = var.region
}

resource "random_id" "default" {
  byte_length = 8
}

// 日志过滤器
resource "google_logging_project_sink" "my-sink" {
  name                       = "my-pubsub-instance-sink"
  destination                = "pubsub.googleapis.com/${google_pubsub_topic.example.id}"
  filter                     = <<EOF
protoPayload.@type="type.googleapis.com/google.cloud.audit.AuditLog" AND
(
    (
        protoPayload.serviceName="compute.googleapis.com" AND
        (
            protoPayload.methodName=~"^(.*instances.(insert|setMachineType|create)|.*regionInstances.bulkInsert|.*disks.insert|.*addresses.insert|.*globalAddresses.insert)$"
        )
    ) OR
    (
        protoPayload.serviceName="cloudsql.googleapis.com" AND
        protoPayload.methodName=~"^(.*instances.(create|update))$"
    ) OR
    (
        protoPayload.serviceName="redis.googleapis.com" AND
        protoPayload.methodName=~"^(CreateInstance|UpdateInstance)$"
    ) OR
    (
        protoPayload.serviceName="bigquery.googleapis.com" AND
        protoPayload.methodName=~"^(InsertDataset|insert)$"
    ) OR
    (
        protoPayload.serviceName="dataproc.googleapis.com" AND
        protoPayload.methodName=~"^CreateCluster$"
    ) OR
    (
        protoPayload.serviceName="storage.googleapis.com" AND
        protoPayload.methodName=~"^buckets.create$"
    ) OR
    (
        protoPayload.serviceName="file.googleapis.com" AND
        protoPayload.methodName=~"^(CreateInstance|CreateBackup)$"
    ) OR
    (
        protoPayload.serviceName="container.googleapis.com" AND
        protoPayload.methodName=~"^CreateCluster$"
    ) OR
    (
        protoPayload.serviceName="artifactregistry.googleapis.com" AND
        protoPayload.methodName=~"^CreateRepository$"
    ) OR
    (
        protoPayload.serviceName="apigateway.googleapis.com" AND
        protoPayload.methodName=~"^CreateApi$"
    ) OR
    (
        protoPayload.serviceName="clouddeploy.googleapis.com" AND
        protoPayload.methodName=~"^(CreateRollout|CreateTarget)$"
    )
)
EOF
 unique_writer_identity = true
}

// pubsub 消息队列

resource "google_pubsub_topic" "example" {
  name     = "example-topic"
  labels = {
    foo = "bar"
  }
  message_retention_duration = "86600s"
}



# 创建 服务账号
resource "google_service_account" "default" {
  account_id   = "gcf-sa-${random_id.default.hex}"
  display_name = "Cloud Functions Service Account"
}


# 给服务账号 挂角色
resource "google_project_iam_binding" "autolabeling" {
	project = var.project_id
	role = "roles/editor"
	members = [
	"serviceAccount:${google_service_account.default.email}"
	]
}


resource "google_project_iam_binding" "invoker" {
	project = var.project_id
	role = "roles/run.invoker"
	members = [
	"serviceAccount:${google_service_account.default.email}"
	]
}


# 将本地文件夹压缩为压缩包
data "archive_file" "default" {
	type	="zip"
	output_path = "/tmp/function-source.zip"
	source_dir = "function-source/"
}
# 创建存储桶
resource "google_storage_bucket" "autolabel_bucket" {
	name	= "${random_id.default.hex}-gcf-source-bucket"
	location = "US"
	uniform_bucket_level_access = true
}

# 将 压缩包 上传至存储桶
resource "google_storage_bucket_object" "default" {
	name	= "function-source.zip"
	bucket = google_storage_bucket.autolabel_bucket.name
	source = data.archive_file.default.output_path
}

# 创建 cloud function 
resource "google_cloudfunctions2_function" "autolabel" {
	name = "autolabel"
	location = var.region
	description = " label instances resource"
	
	build_config {
		runtime = "go122"
		entry_point = "labelResource"
		
		source {
			storage_source {
				bucket = google_storage_bucket.autolabel_bucket.name
				object = google_storage_bucket_object.default.name
			}
		}
	}
	
	service_config {
		max_instance_count = 3
		min_instance_count = 0
		available_memory = "256M"
		timeout_seconds = 60
		ingress_settings = "ALLOW_INTERNAL_ONLY"
		all_traffic_on_latest_revision = true
		service_account_email = google_service_account.default.email
	}
	
	event_trigger {
		trigger_region = var.region
		event_type = "google.cloud.pubsub.topic.v1.messagePublished"
		pubsub_topic = google_pubsub_topic.example.id
		retry_policy = "RETRY_POLICY_RETRY"
		service_account_email = google_service_account.default.email
	}
}
	




