## IaaS on Google Cloud Demo

### Prerequisites

1. Make sure you have terraform installed
2. Make sure you have the gcloud client sdk installed
   
### Objectives   

The objective of this demonstration is to highlight how you can create a virtual machine in the Cloud - Google Cloud in this case - and deploy our sample `todo-api` to the virtual machine.  The virtual machine will be accessible over the internet, and open up port `22` which is needed to access the VM over `ssh`, and also open port `1080`, which is the default port needed by our sample `todo-api`.

Although you can do all of these activities using the GCP console, which we will look at, we will follow cloud best practices and automate the process using Terraform.  Terraform is an open source tool, and is likely the most widely used tool for cloud automation. Lets get started:

1. The first step requires you to use the GCP GUI Console.  Before we can use Terraform, we need to create a service account that will provide Terraform the permissions needed to create our infrastructure.  There are a number of steps in this process.  Detailed instructions are here:  [SETUP SERVICE ACCOUNT](setup-service-account.md).

2. Now that you have created your service account, we need to protect the json file we just created.  It has a number of permissions, and you dont want somebody to accidently get access to this file.  A **VERY COMMON** mistake is that they get uploaded to GitHub by accident, and then hackers/crypto miners find them and can use resources in your cloud account that **WILL COST YOU MONEY**.  So before we forget lets create a `.gitignore` file and add the following contents (which also don't copy the terraform temporary and generated files that could contain sensitive information you don't want published to a public github repo):

```
#Minimal .gitignore for Terraform and GCP Keys
.terraform*
terraform.tfvars 
*.json
.terraform/
terraform.*
!terraform.tfvars
```
Notice how the `.gitignore` file blocks all of the terraform generated files from finding their way into your repository.  The last line `!terraform.tfvars` is an override allowing that file to be checked into your repository as it does not have any sensitive information. 

3. Now that we have all of this setup. We can use Terraform to provision our infrastructure.  The 3 Terraform commands you will use are:
    - `terraform init`: This command ensure you have all of the plugins you need installed by terraform and that terraform is properly configured.
    - `terraform plan`: This command will look at the current state of your cloud infrastructure (if it even exists), cross reference your terraform definition, and then figure out a plan for executing deltas to make your cloud infrastructure align to the terraform definition.
    - `terraform apply`: This command will apply your terraform configuration to the cloud provider.  If everything is successful it will automatically provision everything that you need.  If you need to make changes, you can just keep running `terraform apply` and it will figure out the difference and apply them.  Note that this command will ask you to approve before it makes any changes.  You can override this with `terraform apply -auto-approve`.  This will just run the terraform command and apply your changes without asking for permission.
    - `terraform destroy`:  This command will delete all of the resources in the cloud that it created.  Thus running `terraform destroy` followed by `terraform approve` will essentially reinstall everything for you.
    - `terraform apply -refresh-only`:  Terraform keeps track of the expected state of your infrastructure in local files.  If somebody else modified your infrastructure, either through an automation tool or the console, or if you modified your infrastructure via the console, Terraform will be out of sync and will likely throw warnings or errors.  This command examines your terraform definition and probes your cloud resources to rebuild the state of your infrastructure.  It only rebuilds your local state, it does not make any changes. So if you get out of sync, you can run this command first and then do a `terraform apply`.
4. One of the coolest features of terraform is that it not only will build your cloud infrastructure, it will also figure out the proper order to create and/or modify the cloud resources.  This could be a very complex task if done by hand. 

Now that we have the process documented, we will next examine the purpose of each terraform file in this repo (with extension `.tf`) and the `terraform.tfvars` file.

#### Terraform Files

In Terraform you can combine all of your infrastructure automation into a single `.tf` file, but its better to break them down into multiple `.tf` files, each one serving a well defined process. The files we created are as follows:

1. `provider.tf`:  Since terraform can be used for a variety of different clouds we need to tell it that we are working with Googles cloud.  This file specifies that for google cloud we need to provide our credentials.json file (securely), identify our project, and the region were we will be operating.  Take special note of the line `credentials = "${file("service-acct.json")}"`.  If you recall we did some pre-work to create a service account via the cloud console and then downloaded a file called `service-account.json`.  This file enables Terraform to interact with the cloud and to have all of the permissions that were previously setup in this service account.  
2. `network.tf`: One common operation you need to do in the cloud is to create and manage your own private network.  We will see later where this enables resilience and scale.   For this simple demo, we are creating a simple virtual private cloud (VPC) network and naming it `vpc-network`.  We then create a subnet to deploy our cloud resources and base it in the `us-central1` region.  Ill cover why i picked this region in class, basically its because they have access to a particular cheap machine type - more on this soon.  By default, cloud resources do not get very much access, so `network.tf` defines two firewall rules - one to enable `ssh` ingress over port 22 which is helpful for management, and another firewall rule for our test API that opens a few ports, including the default port for the api example, which is 1080.
3. `storage.tf`:  One thing we want to ultimately be able to do is to run our demo api system from our VM.  In order to do this, we want to push our code to the cloud, so that we can then copy it down once we create our virtual machine.  This terraform file first creates a storage bucket named `csts680-class-bucket` and then it copies both the AMD and ARM based binaries from our local file system up to this bucket.  When we create our virtual machine we will specify the machine type.  Most cloud providers support both Intel(AMD) and ARM based processors.  To be as flexible as possible we upload binaries compiled into both formats for later use. 
4. `service-account.tf`:  If you recall we setup a service account manually to give Terraform the permissions that it needs.  Our goal is to deploy and run a virtual machine in the cloud.  When that machine "boots", we want it to also run under a service account.  Note the one that we created for terraform has a lot of permissions that are not needed for the VM.  By default, VMs running in the cloud get little to no permission to interact with the cloud environment.  This terraform module creates a service account named `cst680-vm-service-account` and assigns it a few specific roles.  To learn more about these roles you can read about them here: https://cloud.google.com/compute/docs/access/iam.  At a high level we want to our VM to be able to SSH to other cloud resources, enable us to login as a local administrator, and finally, the ability to view and copy object from storage buckets. We need the last permission to copy down the binaries that we uploaded in the `storage.tf` file.  The last thing of note here, is at the bottom we define some terraform variables, `project_id` and `vc_acct_email`.  We did not have to do this, but I wanted to demonstrate how we can use variables.  Note these values reference keys in the `terraform.tfvars` file.
5. `main.tf`:  This is were we create the VM itself.  This is a key file to understand.  The main resource we are creating here is a `google_compute_instance`.  At the top we provide some top level parameters such as:
   - The VM Name, `api-vm`
   - The Machine Type, this tells google what type of compute, CPU, memory, etc you want.  For this we picked a minimal machine, `n1-standard-1`
   - What availability zone we want the machine to run in, `us-central1-a`
   - The `service_account` section assigns the service account that the VM will run under (see `service-account.tf`)
   - The `boot_disk` section tells google we want to run a `debian-11` version of Unix.  Note the main cloud providers provide a lot of different options.
   - The `metadata_startup_script` specifies the startup script to run once the machine boots.  This script updates the OS itself (aka patches it), installs a few packages that we might hack with, for example `go`, `python`, `snap`, etc. It then uses the `gsutil` command that is packaged with the `debian-11` OS to copy down the executables for our sample app from the storage bucket we created earlier.
   - The `netowork_interface` section does two things.  First is associates the network subnet that the VM will run on.  Its the one we created in the `network.tf` file. It also specifies an empty `access_config` section that instructs google cloud to create and assign a public IP address to this instance. 
   - Finally, at the bottom you will see one last command, `output "ip"`.  This simply prints the IP address that was assigned to this instance to the terminal.  This is helpful so that we can start interacting with our VM without having to go into the console to look up the VM's public IP address.

