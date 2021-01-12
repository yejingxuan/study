## win10实用工具

- 文件搜索工具：everything


- 看图工具：极速看图


- 批量粘贴板：ditto


- 屏幕录制成gif：screentogif


- 免费图床工具：PicGos
  - token:58aa809552fbff8b2827f17dc24673ac
  - https://molunerfinn.com/PicGo/



- 在此处开启cmd——OpenCmdHere.reg

  ```yaml
  Windows Registry Editor Version 5.00
  [HKEY_CLASSES_ROOT\Directory\shell\OpenCmdHere]
  @="open cmd on here"
  "Icon"="cmd.exe"
  [HKEY_CLASSES_ROOT\Directory\shell\OpenCmdHere\command]
  @="PowerShell -windowstyle hidden -Command \"Start-Process cmd.exe -ArgumentList '/s,/k, pushd,%V' -Verb RunAs\""
  [HKEY_CLASSES_ROOT\Directory\Background\shell\OpenCmdHere]
  @="open cmd on here"
  "Icon"="cmd.exe"
  [HKEY_CLASSES_ROOT\Directory\Background\shell\OpenCmdHere\command]
  @="PowerShell -windowstyle hidden -Command \"Start-Process cmd.exe -ArgumentList '/s,/k, pushd,%V' -Verb RunAs\""
  [HKEY_CLASSES_ROOT\Drive\shell\OpenCmdHere]
  @="open cmd on here"
  "Icon"="cmd.exe"
  [HKEY_CLASSES_ROOT\Drive\shell\OpenCmdHere\command]
  @="PowerShell -windowstyle hidden -Command \"Start-Process cmd.exe -ArgumentList '/s,/k, pushd,%V' -Verb RunAs\""
  [HKEY_CLASSES_ROOT\LibraryFolder\background\shell\OpenCmdHere]
  @="open cmd on here"
  "Icon"="cmd.exe"
  [HKEY_CLASSES_ROOT\LibraryFolder\background\shell\OpenCmdHere\command]
  @="PowerShell -windowstyle hidden -Command \"Start-Process cmd.exe -ArgumentList '/s,/k, pushd,%V' -Verb RunAs\""
  ```

  

