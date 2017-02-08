# Specifications for Container Storage Interfaces

The purpose of this project is to define the various (vendor agnostic) interfaces between cloud native schedulers and persistent data services.

This spec covers two aspects of orchestrating the deployment of data services via a scheduler:

1. The bootstrap deployment of the data service container itself.
2. The runtime communication between a scheduler agent and the data service container.

## Bootstrap Deployment of Data Service Resources
This section of the spec describes how data service providers are deployed by orchestration software.  For example, these providers can be packaged as Linux Containers and they would need to be depoyed on the physical infrastructure by the orchestration software.  This is specified in [api/bootstrap.go](api/bootstrap.go).

## Runtime communication between the scheduler and the data service
Once the data service has been deployed, there are 4 specific interfaces that schedulers and data service providers need to implement.  This is done via REST over a UNIX socket file placed at `/var/run/csi/csi.sock`.

### 1. Discovery of Data Services
Applications that rely on data services should be able to dynamically discover where the provisioned resources are available.  The data service API should also be able to influence where and when these services should be scheduled based on the underlying constraints.

### 2. Provisioning and Instantiation of Data Services
The allocation, use (read and write) and destruction (what used to be known as CRUD) needs to be orchestrated through this interface.

### 3. Lifecycle Operations on Data Services
Data state and its lifecycle, such as retention levels, version levels, access controls should be separated from the actual application that uses them.  It should also be controlled by the scheduling software and it is the goal of this API to define how that is goverened.

### 4. Security
This defines a set of constraints around how a container can authenticate itself in order to operate on a storage service.  This would prevent a container launched by a user from accessing a volume they do not have access to.

### 5. Performance SLAs
This defines how a scheduler can request a specific level of performance from the data services layer for a given application container and gets scheduled in the node that can deliver the performance and meets the application performance requirements

### 6. Multi-tenancy
This defines how a scheduler can ensure that the deployment of application containers with different security, performance and availability requirements in the same scale-out cluster without compromising each other's SLA metrics

### 7. Availability
This definies how a scheduler can describe the availability levels of a given data service, that ensures the availability is provided across rack failures, data center failures, public cloud zone and region failures. This is how a cloud native application can request different levels of availability for the requested data service and the underlying data services layer depending on its capability and available resources at its disposal can either allocate the resources or decline. 

### 8. Statistics
This defines the set of standard performance counters, utilization counters and error counters for the scheduler to query about a given data service from the data services layer. This could help application users perform troubleshooting, capacity planning and general monitoring of application performance and insight

### 9. Billing and Chargeback
This defines how an underlying data services layer can provide cloud native application consumers a view of the costs incurred in consuming the data service they have requested and handle credits of newly acquired service as well as unused servies





## Licensing
`CNCF-CSI` is licensed under the Apache License, Version 2.0. See LICENSE for the full license text.

## Contributing
Want to collaborate and add? Here are instructions to [get started contributing code](contributing.md)
