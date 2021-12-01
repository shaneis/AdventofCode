#! /usr/bin/pwsh

[CmdletBinding()]
param (
	[string] $Path
)

$Measurements = Get-Content -Path $Path

foreach ($Measurement in $Measurements) {
    if ($null -eq $IncreaseCounter) {
        $IncreaseCounter = 0
    } elseif ([int] $Measurement -gt $PrevMeasurement) {
        $IncreaseCounter++
    }

    [PSCustomObject] @{
        Measurement = $Measurement
        Increases = $IncreaseCounter
    }

    [int] $PrevMeasurement = $Measurement
}

