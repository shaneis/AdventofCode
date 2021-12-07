#! /usr/bin/pwsh

[CmdletBinding()]
param (
    $Path
)

# Helper functions
. $PSScriptRoot/Get-CommonBit.ps1

$DiagnosticReports = Get-Content -Path $Path

Write-Verbose "Diagnostic Reports$([Environment]::NewLine)$($DiagnosticReports | Out-String)"

$DiagnosticReportLength = $DiagnosticReports[0].Length

$OxygenReports = $DiagnosticReports.Clone()
foreach ($Iteration in (0..($DiagnosticReportLength -1))) {
    # Only 1 number left
    if (($OxygenReports).Count -eq 1) {
        Write-Verbose "Only 1 object remaining - cancelling"
        break
    }

    Write-Verbose "Getting most common bit for iteration $Iteration"
    $CommonBit = $OxygenReports | Get-CommonBit -Position $Iteration -Type O2

    Write-Verbose "Removing reports that do not have $($CommonBit.Name) in position $Iteration"
    $OxygenReports = $OxygenReports | Where-Object {
        $_[$Iteration] -eq $CommonBit.Name
    }

    Write-Verbose "New diagnostic reports$([Environment]::NewLine)$($OxygenReports | Out-String)"
}

$CO2Reports = $DiagnosticReports.Clone()
foreach ($Iteration in (0..($DiagnosticReportLength -1))) {
    # Only 1 number left
    if (($CO2Reports).Count -eq 1) {
        Write-Verbose "Only 1 object remaining - cancelling"
        break
    }

    Write-Verbose "Getting least common bit for iteration $Iteration"
    $CommonBit = $CO2Reports | Get-CommonBit -Position $Iteration -Type CO2

    Write-Verbose "Removing reports that do not have $($CommonBit.Name) in position $Iteration"
    $CO2Reports = $CO2Reports | Where-Object {
        $_[$Iteration] -eq $CommonBit.Name
    }

    Write-Verbose "New diagnostic reports$([Environment]::NewLine)$($CO2Reports | Out-String)"
}

$OxygenGeneratorRating = [Convert]::ToInt64($OxygenReports, 2)
$CO2ScrubberRating = [Convert]::ToInt64($CO2Reports, 2)

[PSCustomObject] @{
    OxygenGenerator = $OxygenGeneratorRating
    CO2Scrubber = $CO2ScrubberRating
    LifeSupport = $OxygenGeneratorRating * $CO2ScrubberRating
}
