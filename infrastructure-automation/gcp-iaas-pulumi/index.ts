import * as pulumi from "@pulumi/pulumi";
import * as gcp from "@pulumi/gcp";

//Pull in some configuration parameters
const config = new pulumi.Config();
const GcpRegion = config.require("region");
const GcpZone = config.require("zone");

// STORAGE BUCKET - Create a GCP Storage Bucket and Copy our executable to it

// Create a GCP resource (Storage Bucket)
const bucket = new gcp.storage.Bucket("cst680-class-bucket", {
    name: "cst680-class-bucket",
    location: GcpRegion,
});

// Copy the file to the newly created bucket
const object = new gcp.storage.BucketObject("to-do-api", {
    bucket: bucket.name,
    name: "todo-linux-amd64",  // Destination file name in the bucket
    source: new pulumi.asset.FileAsset("../../todo-api/todo-linux-amd64"), // Local file path
});

//NETWORK - SETUP THE NETWORK for our VM
const network = new gcp.compute.Network("network");
const computeFirewall = new gcp.compute.Firewall("firewall", {
    network: network.id,
    allows: [{
        protocol: "tcp",
        ports: [ "22", "80", "1080" ],
    }],
    sourceRanges: ["0.0.0.0/0"]
});


//SERVICE ACCOUNT - Create a service account for our VM.  This will be the
//account that the VM runs under giving it permissions to our bucket and to 
//access the internet
const serviceAccount = new gcp.serviceaccount.Account("vm-service-acct", {
    accountId: "vm-account-id",
    displayName: "Service Account for VM",
});

const bucketAccess = new gcp.storage.BucketIAMBinding("my-bucket-access", {
    bucket: bucket.name,
    role: "roles/storage.objectViewer",
    members: [ pulumi.interpolate`serviceAccount:${serviceAccount.email}`],
}); 


//VM

const startupScript = pulumi.interpolate `
#!/bin/bash
sudo apt-get update
gsutil cp gs://${bucket.name}/${object.name} /usr/local/bin/${object.name} > /tmp/gcp_amd_copy.log 2>&1
chmod +x /usr/local/bin/${object.name}
`;


const computeInstance = new gcp.compute.Instance("api-vm", {
    name: "api-vm",
    machineType: "n1-standard-1",
    zone: "us-west1-a",
    metadataStartupScript: startupScript,
    bootDisk: { initializeParams: { image: "debian-cloud/debian-11" } },
    networkInterfaces: [{
        network: network.id,
        // accessConfigs must include a single empty config to request an ephemeral IP
        accessConfigs: [{}],
    }],
    serviceAccount: {
        email: serviceAccount.email,
        scopes: ["https://www.googleapis.com/auth/cloud-platform"],
    },
});


//EXPORTS - these will include important information that are helpful for
//connecting to the VM and bucket.  You can always get this information from
//the GCP console, but its nice to have it here for reference.

// Export the DNS name of the bucket
export const bucketName = bucket.url;
// Export the name and IP address of the Instance
export const instanceName = computeInstance.name;
export const publicIp = computeInstance.networkInterfaces.apply(nic => nic[0]?.accessConfigs?.[0]?.natIp);
