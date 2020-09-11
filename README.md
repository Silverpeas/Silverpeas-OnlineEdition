Handler of the custom Web protocol scheme spwebdav.

The spwebdav custom Web protocol is dedicated to open a WebDAV bridge between the client host and Silverpeas. The handler is executed when a file is accessed through an URI starting with the spwebdav scheme. The handler then execute an office suite program to open the document at the WebDAV endpoint in Silverpeas.

To build the handler, just use one of the following script:
* \_buildW32 for Windows 7 32bits
* \_buildW64 for Windows 64bits
* \_buildL64 for Linux 64bits
* \_buildM64 for MacOS X 64bits (Intel processor)

For Windows, Inno Setup tool is used to generate Setup program from installer_setup32.iss and installer_setup64.iss script files (for respectively Windows 32bits and Windows 64bits).
The setup generation is performed directly by `_buildW32.bat` or `_buildW64.bat` script file (Inno Setup install path **MUST** be referenced by **%INNOSETUP%** environment variable).

For Linux, the \_buildL64.sh compile the Linux version of the handler and then build a distribution archive of the handler with the `spwebdav.desktop` desktop application descriptor.
