Handler of the custom Web protocol scheme spwebdav.

The spwebdav custom Web protocol is dedicated to open a WebDAV bridge between the client host and Silverpeas. The handler is executed when a file is accessed through an URI starting with the spwebdav scheme. The handler then execute an office suite program to open the document ant the WebDAV endpoint in Silverpeas.

To build the handler, just use one of the following script:
* _buildW32 for Windows 7 32bits
* _buildW64 for Windows 64bits
* _buildL64 for Linux 64bits
* _buildM64 for MacOS X 64bits (Intel processor)


