## IaaS on Google Cloud Demo

### Prerequisites

1. Make sure you have terraform installed
2. Make sure you have the gcloud client sdk installed
   
### Objectives   

The objective of this demonstration is to highlight how you can create a virtual machine in the Cloud - Google Cloud in this case - and deploy our sample `todo-api` to the virtual machine.  The virtual machine will be accessible over the internet, and open up port `22` which is needed to access the VM over `ssh`, and also open port `1080`, which is the default port needed by our sample `todo-api`.

Although you can do all of these activities using the GCP console, which we will look at, we will follow cloud best practices and automate the process using Terraform.  Terraform is an open source tool, and is likely the most widely used tool for cloud automation. Lets get started:

1. The first step requires you to use the GCP GUI Console.  Before we can use Terraform, we need to create a service account that will provide Terraform the permissions needed to create our infrastructure.  There are a number of steps in this process.  Detailed instructions are here:  [SETUP SERVICE ACCOUNT](setup-service-account.md).

2. Now that you have created your service account, we need to protect the json file we just created.  It has a number of permissions, and you dont want somebody to accidently get access to this file.  A **VERY COMMON** mistake is that they get uploaded to GitHub by accident, and then hackers/crypto miners find them and can use resources in your cloud account that **WILL COST YOU MONEY**.  So before we forget lets create a `.gitignore` file and add the following contents (which also don't copy the terraform temp files):

```
#Minimal .gitignore for Terraform and GCP Keys
.terraform*
*.json
.terraform/
```
3. Now that we have all of this setup. We can use Terraform to provision our infrastructure.  The 3 Terraform commands you will use are:
    - `terraform init`: This command ensure you have all of the plugins you need installed by terraform and that terraform is properly configured.
    - `terraform plan`: This command will look at the current state of your cloud infrastructure (if it even exists), cross reference your terraform definition, and then figure out a plan for executing deltas to make your cloud infrastructure align to the terraform definition.
    - `terraform apply`: This command will apply your terraform configuration to the cloud provider.  If everything is successful it will automatically provision everything that you need.  If you need to make changes, you can just keep running `terraform apply` and it will figure out the difference and apply them.  Note that this command will ask you to approve before it makes any changes.  You can override this with `terraform apply -auto-approve`.  This will just run the terraform command and apply your changes without asking for permission.
    - `terraform destroy`:  This command will delete all of the resources in the cloud that it created.  Thus running `terraform destroy` followed by `terraform approve` will essentially reinstall everything for you.
4. One of the coolest features of terraform is that it not only will build your cloud infrastructure, it will also figure out the proper order to create and/or modify the cloud resources.  This could be a very complex task if done by hand. 

Now that we have the process documented, we will next examine the purpose of each terraform file in this repo (with extension `.tf`) and the `terraform.tfvars` file.

#### Terraform Files

In Terraform you can combine all of your infrastructure automation into a single `.tf` file, but its better to break them down into multiple `.tf` files, each one serving a well defined process. The files we created are as follows:

