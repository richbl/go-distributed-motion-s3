# Distributed Motion Surveillance Security System (DMS<sup>3</sup>)

![GitHub release (latest SemVer including pre-releases)](https://img.shields.io/github/v/release/richbl/go-distributed-motion-s3?include_prereleases) [![Go Report Card](https://goreportcard.com/badge/github.com/richbl/go-distributed-motion-s3)](https://goreportcard.com/report/github.com/richbl/go-distributed-motion-s3) [![Codacy Badge](https://app.codacy.com/project/badge/Grade/d81b7869ac134229b78105544e783667)](https://app.codacy.com/gh/richbl/go-distributed-motion-s3/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade) [![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=richbl_go-distributed-motion-s3&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=richbl_go-distributed-motion-s3)

## What Is **DMS<sup>3</sup>**?

<p align="center">
  <img src="https://user-images.githubusercontent.com/10182110/150719391-a562ac4a-154e-4dad-b4bc-6c88f4d2b425.png" alt="DMS3Mail Event">
</p>

**Distributed Motion Surveillance Security System (DMS<sup>3</sup>)** is a [Go-based](https://golang.org/ "Go") application that integrates third-party open-source motion detection applications (*e.g.*, the [Motion](https://motion-project.github.io/ "Motion") motion detection software package, or [OpenCV](http://opencv.org/ "OpenCV"), the Open Source Computer Vision Library) into an automated distributed motion surveillance system that:

- Using a local network, wirelessly senses when someone is "at home" and when someone is "not at home" and automatically enables or disables the surveillance system
- Through the **DMS<sup>3</sup>Server**, the system coordinates video stream processing, reporting, and user notification to participating device clients (*e.g.*, a Raspberry Pi or similar) running the **DMS<sup>3</sup>Client** component which:
    - Greatly minimizes network congestion, particularly during high-bandwidth surveillance events of interest
    - Better utilizes device client CPU/GPU processing power: keeping stream processing on-board and distributed around the network
- Optionally, **DMS<sup>3</sup>Clients** can generate email reports for events of interest containing images or video using the available **DMS<sup>3</sup>Mail** component
- Optionally, the **DMS<sup>3</sup>Server** can display the current state of all reporting **DMS<sup>3</sup>Clients** visually through the use of the **DMS<sup>3</sup>Dashboard** component
- Works cooperatively with "less smart" device clients such as IP cameras (wired or WiFi), webcams, and other USB camera devices

## Want to Know More?

For more information about **DMS<sup>3</sup>**, check out the [DMS<sup>3</sup> project wiki](https://github.com/richbl/go-distributed-motion-s3/wiki). The wiki includes the following sections:

- Project overview
    - Use cases
    - Features
    - Components
    - Architecture
    - How **DMS<sup>3</sup>** works
    - Requirements
- **DMS<sup>3</sup>** Release Notes
- Application installation
    - Downloading, building, and installing the application
    - Running the application
- Project roadmap
- Project license