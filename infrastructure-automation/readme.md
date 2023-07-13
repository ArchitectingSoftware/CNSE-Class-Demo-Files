## Infrastructure Automation

### Introduction
This repository will illustrate how we can use infrastructure automation to deploy our sample [todo-api](../todo-api/) to the cloud.  For this demonstration we will be using GCP.  Two different approaches are shown. 

### Google Cloud Platform

This demonstration uses the Google Cloud Platform (GCP).  If you want to play with GCP, you need a google account.  You can also see me and I can get you some credits to play with the platform.  Every major cloud platform does things a little bit differently.  For this Demo we will be creating and using two different _GCP Projects_, one to demonstrate deployment using the Terraform tooling, and one to do the same thing using Pulumi tooling.  GCP uses the concept of _"Projects"_ to organize a collection of cloud resources.  

### Demonstrations

1. [Terraform Demo](./gcp-iaas-terraform/).  This directory contains all of the infrastructure automation to deploy our sample API to GCP using Terraform.
2. [Pulumi Demo](./gcp-iaas-pulumi/).  This directory contains all of the infrastructure automation to deploy our sample API to GCP using Pulumi.

### Warnings - PLEASE READ
When working with cloud resources in conjunction with a service like GitHub, you need to be **VERY CAREFUL** that certain information is not accidentally pushed to GitHub, GitLab, or any publicly accessible location on the Web.  When you do infrastructure automation you must give the tools you use permissions to create, destroy and configure resources in your cloud account (that is common for all cloud providers, not just GCP).

While all cloud providers take a somewhat different approach, they all generally involve configuring a _"Service Account"_ that will be granted permissions to create and manage your infrastructure. During the creation process, certain keys will be generated and then you will be given the opportunity to download these keys in a certain file format. For example, in this demonstration we will be creating a service account in GCP, and then downloading the credentials associated with this _Service Account_ in the form of a _JSON_ file.  We will be configuring both Terraform and Pulumi to use these credential files so that these tools have the proper permissions to manage your infrastructure. 

There have been **MANY** documented instances where people accidentally checked in credential files into a repository service like GitHub, and then bad actors acquired these credentials to hijack your account. Leaking these credentials can result in a bad actor messing with your infrastructure, which is not a problem for this class, but can lead to **VERY BAD** things in a corporate use case.  Some of the bad things that can happen include:

1. Messing with, reconfiguring, deleting components of your infrastructure
2. Gaining access to sensitive information, databases, etc
3. Provisioning their own resources to do things like crypto-mining that can result in massive cloud bills that you may be responsible for paying.

Many corporate and enterprise users of the cloud spend a lot of time and effort implementing protective controls to avoid this situation, but even in the best organizations there is no way to 100% avoid accidental leakage of sensitive information, so you need to be extra careful.

As a student learning the cloud there are a couple of things that you can do and be aware of to avoid these issues:

1. If you are using a cloud provider for education, you will often be granted a number of credits.  After those credits are exhausted, the cloud provider will disable all of your resources and block you from creating any more until you provide either more credits, or a credit-card for payment.  
2. You can setup spending limits to protect yourself.  For example, on my personal AWS account I setup a $25/month spending limit.  Thus, I cut my losses at $25 if something bad happens, and of course i want to do more where $25 is not enough, I can adjust the spending limit. 
3. The best defense is to be mindful and careful following best practices.  In this demonstration I will be using `.gitignore` files to make sure things that should not be pushed to my repository are kept local on my machine.  Please follow these practices.  If you somehow messup, and then realize, you not only have to adjust your `.gitignore` file to prevent pushing sensitive information, you **must also carefully prune your git history**.  Git keeps track of all of your history, so bad actors, who are quite capable, can still find these things in your history. 

**YOU HAVE BEEN WARNED - DON'T LET THIS PROBLEM HAPPEN TO YOU**

### Notes

_**Note 1: Both Terraform and Pulumi offer free open source and commercial support.  For this demo the free version will work fine for both tools.  These tutorials do not cover how to install these tools on your machine**_

_**Note 2: If you are new to infrastructure automation, please start with the Terraform demo first.  There is a lot of documentation.  The Pulumi demo assumes you are familiar with some of the concepts in the Terraform demo and has lighter documentation on how to use it**_
