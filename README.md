# Specifications for Container Storage Interfaces

The purpose of this project is to define the various interfaces between cloud native schedulers and persistent data services.

There are 4 specific interfaces this project defines that schedulers and data service providers need to implement:

## Bootstrap Deployment of Data Service Resources
This section of the API describes how data service providers are deployed by orchestration software.  For example, these providers can be packaged as Linux Containers and they would need to be depoyed on the physical infrastructure by the orchestration software.

## Discovery of Data Services
Applications that rely on data services should be able to dynamically discover where the provisioned resources are available.  The data service API should also be able to influence where and when these services should be scheduled based on the underlying constraints.

## Provisioning and Instantiation of Data Services
The allocation, use (read and write) and destruction (what used to be known as CRUD) needs to be orchestrated through this interface.

## Lifecycle Operations on Data Services
Data state and its lifecycle, such as retention levels, version levels, access controls should be separated from the actual application that uses them.  It should also be controlled by the scheduling software and it is the goal of this API to define how that is goverened.

## Licensing
`CNCF-CSI` is licensed under the Apache License, Version 2.0. See LICENSE for the full license text.

## Contributing
Want to collaborate and add? Here are instructions to [get started contributing code](contributing.md)
