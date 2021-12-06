#! /usr/bin/pwsh

[CmdletBinding()]
param (
    $Path
)

# Helper functions
. $PSScriptRoot/Get-MostCommonBit.ps1

$DiagnosticReports = Get-Content -Path $Path

$DiagnosticReports

"Most common bit position 0"
$DiagnosticReports | Get-MostCommonBit -Position 0

$DiagnosticReportLength = $DiagnosticReports[0].Length
foreach ($Iteration in (0..($DiagnosticReportLength -1))) {
    # Only 1 number left
    if (($DiagnosticReports).Count -eq 1) {
        Write-Verbose "Only 1 object remaining - cancelling"
        break
    }

    Write-Verbose "Getting most common bit for iteration $Iteration"
    $MostCommonBit = $DiagnosticReports | Get-MostCommonBit -Position $Iteration

    Write-Verbose "Removing reports that do not have $($MostCommonBit.Name) in position $Iteration"
    $DiagnosticReports = $DiagnosticReports | Where-Object {
        $_[$Iteration] -eq $MostCommonBit.Name
    }

    Write-Verbose "New diagnostic reports"
    Write-Verbose "$([Environment]::NewLine)$($DiagnosticReports | Out-String)"
}

$OxygenGeneratorRating = [Convert]::ToInt64($DiagnosticReports, 2)
$OxygenGeneratorRating
