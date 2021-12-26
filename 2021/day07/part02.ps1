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
    Write-Verbose ('Checking position: {0}/{1}' -f $Pos, $High)

    $TotalFuel = 0

    foreach ($Crab in $CrabPos) {
        $FuelCost = 0

        $Steps = [Math]::Abs($Crab - $Pos)
        <# https://en.wikipedia.org/wiki/1_%2B_2_%2B_3_%2B_4_%2B_%E2%8B%AF #>
        $FuelCost = ($Steps * ($Steps + 1)) / 2

        $TotalFuel += $FuelCost
    }

    $LowestFuel = $TotalFuel -lt $LowestFuel ? $TotalFuel : $LowestFuel
}
'Total fuel: {0}' -f $LowestFuel
