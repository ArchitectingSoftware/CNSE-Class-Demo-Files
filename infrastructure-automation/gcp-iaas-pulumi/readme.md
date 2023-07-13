## Pulumi Automation

### Objectives
While Terraform is the industry leader in cloud automation.  Terraform is based on HCL, a proprietary markup language.  In this demo we are going to be looking at Pulumi, that provisions infrastructure using standard programming languages such as Javascript, Typescript, Python, Go, .Net, etc.

Note that the directions below assume you have some familiarity with configuring some things directly in the GCP console.  The Terraform demonstration has a lot of documentation that assumes you do not know anything.  I highly recommend that you start with that demonstration first. 

### Directions

1. Install Pulumi
2. Create a GCP project for this infrastructure using the console
3. Setup Service Account for Pulumi in GCP. Look at the terraform documentation for how to do this.
4. Create new project: `pulumi new gcp-typescript`.  Use defaults, we are going to setup a stack for development called "dev", which is the default
5. Setup .gitignore
6. Add credentials to config: `pulumi config set gcp:credentials pulumi-credentials.json`
7. Setup variables for GCP region and zone:
    - `pulumi config set region us-central1`
    - `pulumi config set zone  us-central1-a`
8. To provision:  `pulumi up`
9. To destroy:  `pulumi destroy`