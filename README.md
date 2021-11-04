# Virtual Desktop Environment Using GO and Fyne
<img src="https://user-images.githubusercontent.com/89251393/139284579-441f4be3-1069-4981-b879-b96bf5fd6ad6.png" width="150"/> <img src="https://user-images.githubusercontent.com/89251393/139284585-4beb4fba-52a6-4a16-abc0-a0ca439558ab.png" width="60"/> <img src="https://user-images.githubusercontent.com/89251393/139284590-a15400c8-6c0c-4703-aff1-876cc5c67e10.png" width="70"/> 

Hello Reader. If you have clicked this you must be excited to know about how to make a Virtual Environment or how to use one for yourself. In this project I have created a basic and simple Virtual Desktop using **GO Language** and **FYNE Module**.


## Contents 
1. [Introduction](#introduction)
2. [Installation](#installation)
3. [Building Virtual Desktop](#building-virtual-desktop)
4. [Working](#working)
5. [Screenshots](#screenshots)
6. [Contact Information](#contact-information)

## Introduction

**GO** is a statically typed, compiled programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson. Go is syntactically similar to C, but with memory safety, garbage collection, structural typing, and CSP-style concurrency. The language is often referred to as Golang because of its domain name, golang.org, but the proper name is Go. *[(Wikipedia)](https://en.wikipedia.org/wiki/Go_(programming_language))*

**Fyne** is a free and open-source cross-platform widget toolkit for creating graphical user interfaces (GUIs) across desktop and mobile platforms. Fyne uses OpenGL to provide cross-platform graphics. It is inspired by the principles of Material Design to create applications that look and behave consistently across all platforms. It is licensed under the terms of the 3-clause BSD License, supporting the creation of free and proprietary applications. In December 2019 Fyne became the most popular GUI toolkit for Go, by GitHub star count and in early February 2020 it was trending as #1 project in GitHub trending ranks. [*(Wikipedia)*](https://en.wikipedia.org/wiki/Fyne_(software))

## Installation

For using this desktop environment in your PC you need to install the following softwares.
1. Download [GoLang](https://golang.org/dl/) for your Operating System.
2. Download [FYNE](https://developer.fyne.io/started/).

Installing GO and Fyne can be tricky so ensure that you are following the steps as mentioned in the official site.

## Building Virtual Desktop
Now that GO and Fyne are installed in your OS you are ready to build the virtual desktop.
> You must be connected to a stable internet connection to build and run the application.


1. Download the [code](https://github.com/madhur3u/GOVirtualDesktop) and place it in a folder.
2. Now open terminal in this folder and execute the following command.
```go
go build main.go calc.go notes.go gallery.go password.go weather.go music.go numbergame.go getTime.go
```
3. After successful execution of the code you will have a executable ***main*** file in your folder.
4. Open the main executable file and your virtual desktop environment will open.

## Working
The applications are made in different files and then their functions are called in [main.go](https://github.com/madhur3u/GOVirtualDesktop/blob/main/main.go) file to integrate them into the virtual desktop. The icons used in the main screen are loaded from net and all other images used for app background are included in code folder. 

Codes for applications -
1. [Main File](https://github.com/madhur3u/GOVirtualDesktop/blob/main/main.go)
2. [Date Time Module](https://github.com/madhur3u/GOVirtualDesktop/blob/main/getTime.go)
3. [Calculator](https://github.com/madhur3u/GOVirtualDesktop/blob/main/calc.go)
4. [Notepad](https://github.com/madhur3u/GOVirtualDesktop/blob/main/notes.go)
5. [Gallery](https://github.com/madhur3u/GOVirtualDesktop/blob/main/gallery.go)
6. [Weather](https://github.com/madhur3u/GOVirtualDesktop/blob/main/weather.go)
7. [Password Generator](https://github.com/madhur3u/GOVirtualDesktop/blob/main/password.go)
8. [Music Player](https://github.com/madhur3u/GOVirtualDesktop/blob/main/music.go)
9. [Number Game](https://github.com/madhur3u/GOVirtualDesktop/blob/main/numbergame.go)

## Screenshots

![Screenshot from 2021-11-04 10-41-20](https://user-images.githubusercontent.com/89251393/140261587-10c8bd70-f8d3-49f8-93a5-31249758de1e.png)
![Screenshot from 2021-11-04 10-42-57](https://user-images.githubusercontent.com/89251393/140261596-5a91210e-2562-40aa-af28-941cce00e0ab.png)
![Screenshot from 2021-11-04 10-43-04](https://user-images.githubusercontent.com/89251393/140261606-e1f6e7ec-44c6-499b-9883-91b79df7f125.png)
![Screenshot from 2021-11-04 10-43-10](https://user-images.githubusercontent.com/89251393/140261611-2b26db7f-26ac-41d5-bb9d-ebff837681b1.png)
![Screenshot from 2021-11-04 10-43-16](https://user-images.githubusercontent.com/89251393/140261620-ace44920-fcbc-4dca-bd53-5841af8ad773.png)
![Screenshot from 2021-11-04 10-43-22](https://user-images.githubusercontent.com/89251393/140261627-489e9fe3-5e1e-40b3-a5c0-0cb092533084.png)
![Screenshot from 2021-11-04 10-43-29](https://user-images.githubusercontent.com/89251393/140261649-4bec4a93-3796-4866-bc2d-00ab5685ef0d.png)
![Screenshot from 2021-11-04 10-43-44](https://user-images.githubusercontent.com/89251393/140261660-3b5d880e-aa4e-4b13-ba12-686c0ce18136.png)

## Contact Information

I hope you like this. If you want you can modify or add features in the applications according to your preferences.
If you have any query, suggestion, doubts etc., you can contact me on the links given below.

<img src="https://user-images.githubusercontent.com/89251393/138821704-5538f667-ca94-4d9f-ad49-b3c48e1cdb0c.png" width="15"/> [LinkedIn](https://www.linkedin.com/in/madhur3u/)<br> <img src="https://user-images.githubusercontent.com/89251393/138821710-7b7585e0-4766-49ba-8543-c116d4da82c4.png" width="15"/> [Instagram](https://www.instagram.com/madhur3u/) <br><img src="https://user-images.githubusercontent.com/89251393/138821715-eab2496c-e895-4113-a26b-96c087a83d9b.png" width="15"/> m3333u@gmail.com <br><img src="https://user-images.githubusercontent.com/89251393/138822281-9aaf6bdc-2fe0-469a-bd31-ed43bc96dcc2.png" width="15"/> [@madhur3u](https://web.telegram.org/)

<!-- <img src="" width="123"/>  -->
