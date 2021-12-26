#! /usr/bin/pwsh

[CmdletBinding()]
param (
    $Path
)

[int[]] $CrabPos = (Get-Content -Path $Path) -split ','

$High = (
    $CrabPos |
    Group-Object |
    Sort-Object -Property Name -Descending
)[0].Name

$LowestFuel = [int]::MaxValue
foreach ($Pos in 0..$High) {
    Write-Verbose ('Checking position: {0}' -f $Pos)

    $TotalFuel = 0

    foreach ($Crab in $CrabPos) {
        $TotalFuel += [Math]::Abs(($Crab - $Pos))
    }

    $LowestFuel = $TotalFuel -lt $LowestFuel ? $TotalFuel : $LowestFuel
}
'Total fuel: {0}' -f $LowestFuel
