$saveToLocation = Split-Path $script:MyInvocation.MyCommand.Path
$devices = Get-WmiObject Win32_PnPEntity | ?{$_.PNPClass -eq 'AudioEndpoint'} | ?{$_.Present -eq 'True'} | Select Name | ft -HideTableHeaders | Out-String
$export = $devices -replace "(?s)\s*$"
$export