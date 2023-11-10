# GCP Instance modules with tests

## [GCP Instance Module](gcp-instance-module)

It contains terraform definition files to be used as module, this module creates one GCP instance and one firewall rule to expose hello world app deployed in instance via startup script

## [Test for GCP Instance creation](gcp-instance-spec)

It contains go script to run test using terratest, explore the code to understand how it works

### What it test?

- Creation of gcp instance
- GCP instance name

To perform test following next steps:

```sh
# Move to directory that contains go script
cd gcp-instance-spec 

# Update go version and modules
go mod tidy

# Run tests
go test -v

```

## [Test for Hello world app inside GCP instance](hello-world-app-spec)

It contains go script to run test using terratest, explore the code to understand how it works

### What it test?

- Creation of gcp instance
- That the deployed application is alive

To perform test following next steps:

```sh
# Move to directory that contains go script
cd hello-world-app-spec 

# Update go version and modules
go mod tidy

# Run tests
go test -v

```
