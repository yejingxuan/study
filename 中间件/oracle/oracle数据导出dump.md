
```cmd
//导出整个库的数据
exp shanghai_proc/iflytek@192.168.65.32/orcl file=/opt/export/expdat.dmp log=/opt/export/expdat.log full=y


//导入整个库的数据
imp shanghai_proc/iflytek@192.168.75.136/orcl file = /opt/export/expdat.dmp  full=y
```