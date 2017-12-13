
### Introduction

Security is primary concern for all devices connected within network including Internet of Things (IoT) or Smart devices. IoT devices are more vulnerable to attacks because unauthorized access potentially control your home and workspaces. Hence it requires constant monitoring and analysis of IoT device state changes.

### Background

In this project, we used **Belkin Wemo Insight switch** which is connected with an electrical device e.g. Printer, Computer, Lamp e.t.c. Wemo switch acts as an IoT device, it can be controlled from Wifi networks. __Though Belkin Wemo has an official mobile app, it is very basic one, does not store/share fine grained information of device state changes across time__. The switch could be controlled by more than one mobile app user and it is required to monitor the device control request from multiple users. Further, the data collected from device state changes are visualized in highly interactive dashboard for real time monitoring and alerts (in near future).

### Technologies used

- Golang
- AWS RDS
- MySQL
- AWS EC2
- Apache Superset
- Raspberry Pi

### Features

- Monitor IoT device state changes through event subscription
- Store the device state changes in highly scalable and secure relation database (MySQL) on AWS cloud.
- Fine grained control of access control through Security groups configuration
- Visualize the device state trends in highly interactive dashboard for real time monitoring.

### How it works

This module runs in Raspberry Pi Model 3B, a mini computer that has limited computing and memory capacity, approximately Dual core CPU, 1 GB RAM, 32 GB memory, Bluetooth and Wifi support. Raspberry Pi is good enough for basic computing needs, costs only $35 and afforable solution for anyone to learn computing. We chose to use Raspberry Pi because it can run this module without the need for an expensive server/hardware.

#### Why Golang

Golang is a modern programming language from Google. It is simple, edgy, good as any object oriented programming language, built-in concurrency, cross platform support and yet does not demand more device resources such as Java.

#### Wemo (IoT) programming interface

Wemo device published SOAP webervice definition for third party app. developers. SOAP webservices can be consumed through HTTP request/response or HTTP Producer/subscriber pattern.

    http://{local_ip}:{port}/setup.xml
    
c.	Setup and execution instructions - Prerequisites
•	Install Golang from https://golang.org/doc/install
•	Setup RDS instance (MySQL) in AWS free tier account
•	Create EC2 instance and install superset using AWS free tier account
•	Clone source code from GitHub- git clone https://github.com/gowsalyathiru/device-manager.git
•	Navigate to the project directory

               cd devmonitor

•	Download dependency libraries

  go get go.evanpurkhiser.com/netgear
  go get github.com/go-sql-driver/mysql

•	Compile the source files

go build *.go

•	Set program variables as environment variables to avoid hardcoding sensitive data

 export MYSQL_DB_INSTANCE=XXXXX-ABCDEFGH-UV12XY.us-east-1.rds.amazonaws.com
 export MYSQL_DB_USER=XXXXXXX
 export MYSQL_DB_PASSWORD=XXXXXXXXXXXXXXX
export RT_USERNAME=XXXXXXX
export RT_PASSWORD=ZZZZZZZZZZZZ

•	Run the program

               go run *.go
