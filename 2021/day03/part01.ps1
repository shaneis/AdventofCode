#! /usr/bin/pwsh

[CmdletBinding()]
param (
    $Path
)

$DiagnosticReport = Get-Content -Path $Path

$DiagnosisLength = $DiagnosticReport[0].Length

Write-Verbose ($DiagnosticReport | Out-String)

$SummaryReport = foreach ($Iteration in 0..($DiagnosisLength - 1)) {
    $Ones = ($DiagnosticReport | Where-Object {$_[$Iteration] -eq '1'}).Count
    $Zeros = $DiagnosticReport.Count - $Ones

    Write-Verbose "Iteration [$Iteration] has [$Ones] ones and [$Zeros] zeros"

    [PSCustomObject] @{
        Iteration = $Iteration
        Ones = $Ones
        Zeros = $Zeros
        Gamma = (0, 1)[$Ones -gt $Zeros]
        Epsilon = (0, 1)[$Zeros -gt $Ones]
    }
    $Ones = $Zeros = 0
}

$TotalReport = $SummaryReport |
    Select-Object -Property *, @{
        Name = 'GammaDecimal'; Expression = {[Convert]::ToInt64(($SummaryReport.Gamma -join ''), 2)}
    }, @{
        Name = 'EpsilonDecimal'; Expression = {[Convert]::ToInt64(($SummaryReport.Epsilon -join ''), 2)}
    } 

$TotalReport | Format-Table

$FirstGamma = $TotalReport[0].GammaDecimal
$FirstEpsilon = $TotalReport[0].EpsilonDecimal
"Submarine power consumption: Gamma $FirstGamma * Epsilon $FirstEpsilon = $($FirstGamma * $FirstEpsilon)"
