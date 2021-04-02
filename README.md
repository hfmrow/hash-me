# Hash Me

*Calculate the checksum using many methods for one or more files at a time. Save the results in a '.SUM' text file for each given file.*

#### Last update 2021-04-02

Take a look [here, H.F.M repositories](https://github.com/hfmrow/) for other useful linux softwares.

- If you just want to use it, simply download the '*.deb' compiled version under the "release" tab.

- If you want to play inside code, see below "How to compile" section.

## How it's made

- Programmed with go language: [golang](https://golang.org/doc/) 
- GUI provided by [Gotk3 (gtk3 v3.22)](https://github.com/gotk3/gotk3), GUI library for Go (minimum required v3.16).
- I use home-made software: "Gotk3ObjHandler" to embed images/icons, UI-information and manage/generate gtk3 objects code from [glade ui designer](https://glade.gnome.org/). and "Gotk3ObjTranslate" to generate the language files and the assignment of a tooltip on the gtk3 objects (both are not published at the moment, in fact, they need documentations and, for the moment, I have not had the time to do them).

## Functionalities

- Create hash checksum for given files, allow to create .SUM files corresponding to each file.
- Includes Md4, Md5, Sha1, Sha256, Sha384, Sha512, Sha3_256, Sha3_384, Sha3_512, Blake2b256, Blake2b384, Blake2b512 methods.
- Several previews available.
- Multi-file management.
- OS reminder for Mac, Window, Linux (how to control checksum).
- Command line version available.

## Some pictures

**Main window.** 
![loading-ag-640](assets/readme/main.jpg)

###### Cli version![](assets/readme/cli.jpg)

## How to compile

- 

- Open terminal window and at command prompt, type: `go get github.com/hfmrow/hash-me`

- See [Gotk3 Installation instructions](https://github.com/gotk3/gotk3/wiki#installation) for gui installation instruction.

- To change gtk3 interface you need to use a home made software, (not published actually). So don't change gtk3 interface (glade file) ...

- To change language file you need to use another home made software, (not published actually). So don't change language file ...

- To Produce a stand-alone executable, you must change inside "main.go" file:
  
        func main() {
            devMode = true
        ...
  
  into
  
        func main() {
            devMode = false
        ...

This operation indicate that externals data (Image/Icons) must be embedded into the executable file.

### Os informations (build with)

| Name                                                       | Version / Info / Name                          |
| ---------------------------------------------------------- | ---------------------------------------------- |
| GOLANG                                                     | V1.16.2 -> GO111MODULE="off", GOPROXY="direct" |
| DISTRIB                                                    | LinuxMint Xfce                                 |
| VERSION                                                    | 20                                             |
| CODENAME                                                   | ulyana                                         |
| RELEASE                                                    | #46-Ubuntu SMP Fri Jul 10 00:24:02 UTC 2020    |
| UBUNTU_CODENAME                                            | focal                                          |
| KERNEL                                                     | 5.8.0-48-generic                               |
| HDWPLATFORM                                                | x86_64                                         |
| GTK+ 3                                                     | 3.24.20                                        |
| GLIB 2                                                     | 2.64.3                                         |
| CAIRO                                                      | 1.16.0                                         |
| [GtkSourceView](https://github.com/hfmrow/gotk3_gtksource) | 4.6.0                                          |
| [LiteIDE](https://github.com/visualfc/liteide)             | 37.4 qt5.x                                     |
| Qt5                                                        | 5.12.8 in /usr/lib/x86_64-linux-gnu            |

- The compilation have not been tested under Windows or Mac OS, but all file access functions, line-end manipulations or charset implementation are made with OS portability in mind.

## You got an issue ?

- Give information (as above), about used platform and OS version.
- Provide a method to reproduce the problem.
