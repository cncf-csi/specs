# Specifications for Container Storage Interfaces

The purpose of this project is to define the various (vendor agnostic) interfaces between cloud native schedulers and persistent data services.

This spec covers two aspects of orchestrating the deployment of data services via a scheduler:

1. The bootstrap deployment of the data service container itself.
2. The runtime communication between a scheduler agent and the data service container.

## Bootstrap Deployment of Data Service Resources
This section of the spec describes how data service providers are deployed by orchestration software.  For example, these providers can be packaged as Linux Containers and they would need to be depoyed on the physical infrastructure by the orchestration software.  This is specified in [api/bootstrap.go](api/bootstrap.go).

## Runtime communication between the scheduler and the data service
Once the data service has been deployed, there are 4 specific interfaces that schedulers and data service providers need to implement.  This is specified in [api/provider.go](api/provider.go).  The scheduler and the provider could communicate via a runtime `UNIX sock` file on the agent host machine (TBD).

### 1. Discovery of Data Services
Applications that rely on data services should be able to dynamically discover where the provisioned resources are available.  The data service API should also be able to influence where and when these services should be scheduled based on the underlying constraints.

### 2. Provisioning and Instantiation of Data Services
The allocation, use (read and write) and destruction (what used to be known as CRUD) needs to be orchestrated through this interface.

### 3. Lifecycle Operations on Data Services (TBD)
Data state and its lifecycle, such as retention levels, version levels, access controls should be separated from the actual application that uses them.  It should also be controlled by the scheduling software and it is the goal of this API to define how that is goverened.

### 4. Security (TBD)
This defines a set of constraints around how a container can authenticate itself in order to operate on a storage service.  This would prevent a container launched by a user from accessing a volume they do not have access to.  


## Licensing
`CNCF-CSI` is licensed under the Apache License, Version 2.0. See LICENSE for the full license text.

## Contributing
Want to collaborate and add? Here are instructions to [get started contributing code](contributing.md)
